package renderer

import (
	"fmt"
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
	default:
		return nil, fmt.Errorf("unknown renderer: %s", conf.Renderer)
	}
}

func loadTemplate(templatesDir string, funcs template.FuncMap) (*template.Template, error) {
	return template.New("").Funcs(funcs).ParseGlob(filepath.Join(templatesDir, "*.tpl"))
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
