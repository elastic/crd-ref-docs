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
package renderer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/elastic/crd-ref-docs/config"
	"github.com/elastic/crd-ref-docs/types"
)

const mainTemplate = "gvList"

type Renderer interface {
	Render(gvd []types.GroupVersionDetails) error
}

func New(conf *config.Config) (Renderer, error) {
	switch conf.Renderer {
	case "asciidoctor":
		return NewAsciidoctorRenderer(conf)
	case "markdown":
		return NewMarkdownRenderer(conf)
	default:
		return nil, fmt.Errorf("unknown renderer: %s", conf.Renderer)
	}
}

func loadTemplate(templatesFS fs.FS, funcs template.FuncMap) (*template.Template, error) {
	return template.New("").Funcs(funcs).ParseFS(templatesFS, "*.tpl")
}

type funcMap struct {
	prefix string
	funcs  template.FuncMap
}

func combinedFuncMap(funcs ...funcMap) template.FuncMap {
	m := make(template.FuncMap)
	for _, f := range funcs {
		for k, v := range f.funcs {
			m[f.prefix+k] = v
		}
	}

	return m
}

func createOutFile(outputPath string, defaultFileName string) (*os.File, error) {
	finfo, err := os.Stat(outputPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if finfo != nil && finfo.IsDir() {
		outputPath = filepath.Join(outputPath, defaultFileName)
	}

	return os.Create(outputPath)
}
