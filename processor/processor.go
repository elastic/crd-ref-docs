// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package processor

import (
	"fmt"
	gotypes "go/types"
	"regexp"
	"sort"
	"strings"

	"github.com/elastic/crd-ref-docs/config"
	"github.com/elastic/crd-ref-docs/types"
	"go.uber.org/zap"
	"golang.org/x/tools/go/packages"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-tools/pkg/crd"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

const (
	groupNameMarker   = "groupName"
	objectRootMarker  = "kubebuilder:object:root"
	versionNameMarker = "versionName"
)

var ignoredCommentRegex = regexp.MustCompile(`\s*^(?i:\+|copyright)`)

type groupVersionInfo struct {
	schema.GroupVersion
	*loader.Package
	doc   string
	kinds map[string]struct{}
	types types.TypeMap
}

func Process(config *config.Config) ([]types.GroupVersionDetails, error) {
	compiledConfig, err := compileConfig(config)
	if err != nil {
		return nil, err
	}

	p := newProcessor(compiledConfig, config.Flags.MaxDepth)
	// locate the packages annotated with group names
	if err := p.findAPITypes(config.SourcePath); err != nil {
		return nil, fmt.Errorf("failed to find API types in directory %s:%w", config.SourcePath, err)
	}

	p.types.InlineTypes(p.propagateReference)

	// collect references between types
	for typeName, refs := range p.references {
		typeDef, ok := p.types[typeName]
		if !ok {
			return nil, fmt.Errorf("type not loaded: %s", typeName)
		}

		for ref, _ := range refs {
			if rd, ok := p.types[ref]; ok {
				typeDef.References = append(typeDef.References, rd)
			}
		}
	}

	// build the return array
	var gvDetails []types.GroupVersionDetails
	for _, gvi := range p.groupVersions {
		details := types.GroupVersionDetails{GroupVersion: gvi.GroupVersion, Doc: gvi.doc}
		for k, _ := range gvi.kinds {
			details.Kinds = append(details.Kinds, k)
		}

		details.Types = make(types.TypeMap)
		for name, t := range gvi.types {
			key := types.Key(t)

			if p.shouldIgnoreType(key) {
				zap.S().Debugw("Skipping excluded type", "type", name)
				continue
			}
			if typeDef, ok := p.types[key]; ok && typeDef != nil {
				details.Types[name] = typeDef
			} else {
				zap.S().Fatalw("Type not loaded", "type", key)
			}
		}

		gvDetails = append(gvDetails, details)
	}

	// sort the array by GV
	sort.SliceStable(gvDetails, func(i, j int) bool {
		if gvDetails[i].Group < gvDetails[j].Group {
			return true
		}

		if gvDetails[i].Group == gvDetails[j].Group {
			return gvDetails[i].Version < gvDetails[j].Version
		}

		return false
	})

	return gvDetails, nil
}

func newProcessor(compiledConfig *compiledConfig, maxDepth int) *processor {
	p := &processor{
		compiledConfig: compiledConfig,
		maxDepth:       maxDepth,
		parser: &crd.Parser{
			Collector: &markers.Collector{Registry: &markers.Registry{}},
			Checker:   &loader.TypeChecker{},
		},
		groupVersions: make(map[schema.GroupVersion]*groupVersionInfo),
		types:         make(types.TypeMap),
		references:    make(map[string]map[string]struct{}),
	}

	crd.AddKnownTypes(p.parser)
	return p
}

type processor struct {
	*compiledConfig
	maxDepth      int
	parser        *crd.Parser
	groupVersions map[schema.GroupVersion]*groupVersionInfo
	types         types.TypeMap
	references    map[string]map[string]struct{}
}

func (p *processor) findAPITypes(directory string) error {
	cfg := &packages.Config{Dir: directory}
	pkgs, err := loader.LoadRootsWithConfig(cfg, "./...")
	if err != nil {
		return err
	}

	collector := &markers.Collector{Registry: mkRegistry()}
	for _, pkg := range pkgs {
		gvInfo := p.extractGroupVersionIfExists(collector, pkg)
		if gvInfo == nil {
			continue
		}

		if p.shouldIgnoreGroupVersion(gvInfo.GroupVersion.String()) {
			continue
		}

		// let the parser know that we need this package
		p.parser.AddPackage(pkg)

		// if we have encountered this GV before, use that instead
		if gv, ok := p.groupVersions[gvInfo.GroupVersion]; ok {
			gvInfo = gv
		} else {
			p.groupVersions[gvInfo.GroupVersion] = gvInfo
		}

		if gvInfo.types == nil {
			gvInfo.types = make(types.TypeMap)
		}

		// locate the kinds
		markers.EachType(collector, pkg, func(info *markers.TypeInfo) {
			// ignore types explicitly listed by the user
			if p.shouldIgnoreType(fmt.Sprintf("%s.%s", pkg.PkgPath, info.Name)) {
				return
			}

			// ignore unexported types
			if info.RawSpec.Name == nil || !info.RawSpec.Name.IsExported() {
				return
			}

			// load the type
			key := fmt.Sprintf("%s.%s", pkg.PkgPath, info.Name)
			typeDef, ok := p.types[key]
			if !ok {
				typeDef = p.processType(pkg, info, 0)
			}

			p.types[key] = typeDef
			if typeDef != nil && typeDef.Kind != types.BasicKind {
				gvInfo.types[info.Name] = typeDef
			}

			// is this a root object?
			if root := info.Markers.Get(objectRootMarker); root != nil {
				if gvInfo.kinds == nil {
					gvInfo.kinds = make(map[string]struct{})
				}
				gvInfo.kinds[info.Name] = struct{}{}
				typeDef.GVK = &schema.GroupVersionKind{Group: gvInfo.Group, Version: gvInfo.Version, Kind: info.Name}
			}

		})
	}

	return nil
}

func (p *processor) extractGroupVersionIfExists(collector *markers.Collector, pkg *loader.Package) *groupVersionInfo {
	markerValues, err := markers.PackageMarkers(collector, pkg)
	if err != nil {
		pkg.AddError(err)
		return nil
	}

	groupName := markerValues.Get(groupNameMarker)
	if groupName == nil {
		return nil
	}

	version := pkg.Name
	if v := markerValues.Get(versionNameMarker); v != nil {
		version = v.(string)
	}

	gvInfo := &groupVersionInfo{
		GroupVersion: schema.GroupVersion{
			Group:   groupName.(string),
			Version: version,
		},
		doc:     p.extractPkgDocumentation(pkg),
		Package: pkg,
	}

	return gvInfo
}

func (p *processor) extractPkgDocumentation(pkg *loader.Package) string {
	var pkgComments []string

	pkg.NeedSyntax()
	for _, n := range pkg.Syntax {
		if n.Doc == nil {
			continue
		}
		comment := n.Doc.Text()
		commentLines := strings.Split(comment, "\n")
		for _, line := range commentLines {
			if !ignoredCommentRegex.MatchString(line) {
				pkgComments = append(pkgComments, line)
			}
		}
	}

	return strings.Join(pkgComments, "\n")
}

func (p *processor) processType(pkg *loader.Package, info *markers.TypeInfo, depth int) *types.Type {
	typeDef := &types.Type{
		Name:    info.Name,
		Package: pkg.PkgPath,
		Doc:     info.Doc,
	}

	if p.useRawDocstring && info.RawDecl != nil {
		// use raw docstring to support multi-line and indent preservation
		typeDef.Doc = strings.TrimSuffix(info.RawDecl.Doc.Text(), "\n")
	}

	// if the field list is non-empty, this is a struct
	if len(info.Fields) > 0 {
		typeDef.Kind = types.StructKind
		return p.processStructFields(typeDef, pkg, info, depth)
	}

	t := pkg.TypesInfo.TypeOf(info.RawSpec.Type)
	if t == nil {
		zap.S().Warnw("Failed to determine AST type", "package", pkg.PkgPath, "type", info.Name)
		typeDef.Kind = types.UnknownKind
		return typeDef
	}

	tmpType := p.loadType(pkg, t, depth)
	if tmpType == nil {
		typeDef.Kind = types.UnknownKind
		return typeDef
	}

	tmpType.Name = typeDef.Name
	tmpType.Package = typeDef.Package
	tmpType.Doc = typeDef.Doc
	return tmpType
}

func (p *processor) processStructFields(parentType *types.Type, pkg *loader.Package, info *markers.TypeInfo, depth int) *types.Type {
	logger := zap.S().With("package", pkg.PkgPath, "type", parentType.String())
	logger.Debugw("Processing struct fields")
	parentTypeKey := types.Key(parentType)

	for _, f := range info.Fields {
		t := pkg.TypesInfo.TypeOf(f.RawField.Type)
		if t == nil {
			zap.S().Debugw("Failed to determine type of field", "field", f.Name)
			continue
		}

		fieldDef := &types.Field{
			Name:     f.Name,
			Doc:      f.Doc,
			Embedded: f.Name == "",
		}

		if tagVal, ok := f.Tag.Lookup("json"); ok {
			args := strings.Split(tagVal, ",")
			if len(args) > 0 && args[0] != "" {
				fieldDef.Name = args[0]
			}
		}

		logger.Debugw("Loading field type", "field", fieldDef.Name)
		if fieldDef.Type = p.loadType(pkg, t, depth); fieldDef.Type == nil {
			logger.Debugw("Failed to load type for field", "field", f.Name, "type", t.String())
			continue
		}

		if fieldDef.Embedded {
			fieldDef.Inlined = fieldDef.Name == ""
			if fieldDef.Name == "" {
				fieldDef.Name = fieldDef.Type.Name
			}
		}

		if p.shouldIgnoreField(parentTypeKey, fieldDef.Name) {
			zap.S().Debugw("Skipping excluded field", "type", parentType.String(), "field", fieldDef.Name)
			continue
		}

		parentType.Fields = append(parentType.Fields, fieldDef)

		// add to references map
		p.addReference(parentType, fieldDef.Type)
	}

	return parentType
}

func (p *processor) loadType(pkg *loader.Package, t gotypes.Type, depth int) *types.Type {
	if depth > p.maxDepth {
		zap.S().Debugw("Not loading type due to reaching max recursion depth", "type", t.String())
		return nil
	}

	typeDef := mkType(pkg, t)

	zap.S().Debugw("Load", "package", typeDef.Package, "name", typeDef.Name)

	switch x := t.(type) {
	case *gotypes.Pointer:
		if y, ok := p.types[types.Key(typeDef)]; ok {
			typeDef = y.Copy()
			typeDef.UnderlyingType = y
		} else {
			typeDef.UnderlyingType = p.loadType(pkg, x.Elem(), depth+1)
		}
		typeDef.Kind = types.PointerKind
		if typeDef.UnderlyingType != nil && typeDef.UnderlyingType.Kind == types.BasicKind {
			typeDef.Package = ""
		}
		return typeDef

	case *gotypes.Slice:
		if y, ok := p.types[types.Key(typeDef)]; ok {
			typeDef = y.Copy()
			typeDef.UnderlyingType = y
		} else {
			typeDef.UnderlyingType = p.loadType(pkg, x.Elem(), depth+1)
		}
		typeDef.Kind = types.SliceKind
		if typeDef.UnderlyingType != nil && typeDef.UnderlyingType.Kind == types.BasicKind {
			typeDef.Package = ""
		}
		return typeDef

	case *gotypes.Array:
		if y, ok := p.types[types.Key(typeDef)]; ok {
			typeDef = y.Copy()
			typeDef.UnderlyingType = y
		} else {
			typeDef.UnderlyingType = p.loadType(pkg, x.Elem(), depth+1)
		}
		typeDef.Kind = types.ArrayKind
		if typeDef.UnderlyingType != nil && typeDef.UnderlyingType.Kind == types.BasicKind {
			typeDef.Package = ""
		}
		return typeDef
	}

	if x, ok := p.types[types.Key(typeDef)]; ok {
		return x
	}

	switch x := t.(type) {
	case *gotypes.Basic:
		typeDef.Kind = types.BasicKind
		typeDef.Package = ""

	case *gotypes.Map:
		typeDef.Kind = types.MapKind
		typeDef.KeyType = p.loadType(pkg, x.Key(), depth+1)
		typeDef.ValueType = p.loadType(pkg, x.Elem(), depth+1)

	case *gotypes.Named:
		typeDef.Kind = types.AliasKind
		typeDef = p.loadAliasType(typeDef, pkg, x.Underlying(), depth)

	case *gotypes.Interface:
		if x.Empty() {
			typeDef.Kind = types.InterfaceKind
		} else {
			return nil
		}

	default:
		return nil
	}

	p.types[types.Key(typeDef)] = typeDef

	return typeDef
}

func mkType(pkg *loader.Package, t gotypes.Type) *types.Type {
	qualifier := gotypes.RelativeTo(pkg.Types)
	cleanTypeName := strings.TrimLeft(gotypes.TypeString(t, qualifier), "*[]")

	typeDef := &types.Type{
		Name:    cleanTypeName,
		Package: pkg.PkgPath,
	}

	// is this is an imported type?
	if dotPos := strings.LastIndexByte(cleanTypeName, '.'); dotPos >= 0 {
		typeDef.Name = cleanTypeName[dotPos+1:]
		typeDef.Package = cleanTypeName[:dotPos]
		typeDef.Imported = true
	}

	return typeDef
}

func (p *processor) loadAliasType(typeDef *types.Type, pkg *loader.Package, underlying gotypes.Type, depth int) *types.Type {
	tPkg := pkg

	// check whether this type is imported
	if typeDef.Package != pkg.PkgPath {
		imports := pkg.Imports()
		importPkg, ok := imports[typeDef.Package]
		if !ok {
			zap.S().Warnw("Imported type cannot be found", "name", typeDef.Name, "package", typeDef.Package)
			return typeDef
		}

		p.parser.NeedPackage(importPkg)
		tPkg = importPkg
	}

	// find the type from the parser
	tInfo := p.parser.LookupType(tPkg, typeDef.Name)
	if tInfo == nil {
		zap.S().Warnw("Failed to find type", "name", typeDef.Name, "package", typeDef.Package)
		return typeDef
	}

	if bt, ok := underlying.(*gotypes.Basic); ok {
		typeDef.UnderlyingType = &types.Type{Name: bt.String(), Kind: types.BasicKind}
		typeDef.Doc = tInfo.Doc
		return typeDef
	}

	return p.processType(tPkg, tInfo, depth+1)
}

// Every child thas has a reference to 'originalRef', will also get a reference
// to 'additionalRef'.
func (p *processor) propagateReference(originalRef *types.Type, additionalRef *types.Type) {
	originalTypeKey := types.Key(originalRef)
	additionalTypeKey := types.Key(additionalRef)

	for _, parents := range p.references {
		if _, ok := parents[originalTypeKey]; ok {
			parents[additionalTypeKey] = struct{}{}
		}
	}
}

func (p *processor) addReference(parent *types.Type, child *types.Type) {
	if child == nil || child.Kind == types.BasicKind {
		return
	}

	addRef := func(t *types.Type) {
		if t == nil || t.Kind == types.BasicKind {
			return
		}

		parentTypeKey := types.Key(parent)
		childTypeKey := types.Key(t)
		if p.references[childTypeKey] == nil {
			p.references[childTypeKey] = make(map[string]struct{})
		}
		p.references[childTypeKey][parentTypeKey] = struct{}{}
	}

	switch child.Kind {
	case types.AliasKind, types.StructKind:
		addRef(child)
	case types.ArrayKind, types.SliceKind, types.PointerKind:
		addRef(child.UnderlyingType)
	case types.MapKind:
		addRef(child.KeyType)
		addRef(child.ValueType)
	}
}

func mkRegistry() *markers.Registry {
	registry := &markers.Registry{}
	registry.Define(groupNameMarker, markers.DescribesPackage, "")
	registry.Define(objectRootMarker, markers.DescribesType, true)
	registry.Define(versionNameMarker, markers.DescribesPackage, "")
	return registry
}
