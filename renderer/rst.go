package renderer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/elastic/crd-ref-docs/config"
	"github.com/elastic/crd-ref-docs/types"
)

type ReStructuredTextRenderer struct {
	conf *config.Config
	*Functions
}

func NewReStructuredTextRenderer(conf *config.Config) (*ReStructuredTextRenderer, error) {
	baseFuncs, err := NewFunctions(conf)
	if err != nil {
		return nil, err
	}
	return &ReStructuredTextRenderer{conf: conf, Functions: baseFuncs}, nil
}

func (rst *ReStructuredTextRenderer) Render(gvd []types.GroupVersionDetails) error {
	funcMap := combinedFuncMap(funcMap{prefix: "rst", funcs: rst.ToFuncMap()}, funcMap{funcs: sprig.TxtFuncMap()})
	tmpl, err := loadTemplate(rst.conf.TemplatesDir, funcMap)
	if err != nil {
		return err
	}

	outputFile := rst.conf.OutputPath
	finfo, err := os.Stat(outputFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if finfo != nil && finfo.IsDir() {
		outputFile = filepath.Join(outputFile, "out.rst")
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.ExecuteTemplate(f, mainTemplate, gvd)
}

func (rst *ReStructuredTextRenderer) ToFuncMap() template.FuncMap {
	return template.FuncMap{
		"GroupVersionID":     rst.GroupVersionID,
		"RenderExternalLink": rst.RenderExternalLink,
		"RenderGVLink":       rst.RenderGVLink,
		"RenderLocalLink":    rst.RenderLocalLink,
		"RenderType":         rst.RenderType,
		"RenderTypeLink":     rst.RenderTypeLink,
		"SafeID":             rst.SafeID,
		"ShouldRenderType":   rst.ShouldRenderType,
		"TypeID":             rst.TypeID,
		"RenderUnderline":    rst.RenderUnderline,
		"BlankLine":          rst.BlankLine,
		"Spaces":             rst.Spaces,
	}
}

func (rst *ReStructuredTextRenderer) ShouldRenderType(t *types.Type) bool {
	return t != nil && (t.GVK != nil || len(t.References) > 0)
}

func (rst *ReStructuredTextRenderer) RenderType(t *types.Type) string {
	var sb strings.Builder
	switch t.Kind {
	case types.MapKind:
		sb.WriteString("object (")
		sb.WriteString("keys:")
		sb.WriteString(rst.RenderTypeLink(t.KeyType))
		sb.WriteString(", values:")
		sb.WriteString(rst.RenderTypeLink(t.ValueType))
		sb.WriteString(")")
	case types.ArrayKind, types.SliceKind:
		sb.WriteString(rst.RenderTypeLink(t.UnderlyingType))
		sb.WriteString(" array")
	default:
		sb.WriteString(rst.RenderTypeLink(t))
	}

	return sb.String()
}

func (rst *ReStructuredTextRenderer) RenderTypeLink(t *types.Type) string {
	text := rst.SimplifiedTypeName(t)

	link, local := rst.LinkForType(t)
	if link == "" {
		return text
	}

	if local {
		return rst.RenderLocalLink(text)
	} else {
		return rst.RenderExternalLink(link, text)
	}
}

func (rst *ReStructuredTextRenderer) RenderLocalLink(text string) string {
	return fmt.Sprintf(":ref:`%s`", text)
}

func (rst *ReStructuredTextRenderer) RenderExternalLink(link, text string) string {
	return fmt.Sprintf("`%s <%s>`_", text, link)
}

func (rst *ReStructuredTextRenderer) RenderGVLink(gv types.GroupVersionDetails) string {
	return rst.RenderLocalLink(gv.GroupVersionString())
}

func (rst *ReStructuredTextRenderer) RenderUnderline(title string, k string) string {
	var underline = ""
	for range title {
		underline += k
	}
	return underline
}

func (rst *ReStructuredTextRenderer) BlankLine() string {
	return ""
}

func (rst *ReStructuredTextRenderer) Spaces(spaces int) string {
	var a = ""
	for i := 0; i < spaces; i++ {
		a += " "
	}
	return a
}
