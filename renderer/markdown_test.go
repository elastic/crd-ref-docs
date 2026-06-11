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
	"testing"

	"github.com/elastic/crd-ref-docs/config"
	"github.com/stretchr/testify/assert"
)

func newTestConfig(t *testing.T) *config.Config {
	t.Helper()
	conf := &config.Config{}
	if err := conf.TemplateKeyValues.Set("foo=bar"); err != nil {
		t.Fatal(err)
	}
	if err := conf.TemplateKeyValues.Set("hello=world"); err != nil {
		t.Fatal(err)
	}
	return conf
}

func TestMarkdownRenderer_RewriteLinks(t *testing.T) {
	conf := &config.Config{
		Render: config.RenderConfig{
			LinkMappings: []*config.LinkMapping{
				{
					URL:  "https://example.com/old",
					Link: "docs-content://new/page.md",
					Text: "New page",
				},
				{
					URL:  "https://example.com/other",
					Link: "kibana://reference/other.md",
					Text: "Other",
				},
			},
		},
	}
	r := &MarkdownRenderer{conf: conf}

	tests := []struct {
		name     string
		renderer *MarkdownRenderer
		text     string
		want     string
	}{
		{
			name:     "single substitution",
			renderer: r,
			text:     "See https://example.com/old for details.",
			want:     "See [New page](docs-content://new/page.md) for details.",
		},
		{
			name:     "multiple substitutions",
			renderer: r,
			text:     "See https://example.com/old and https://example.com/other.",
			want:     "See [New page](docs-content://new/page.md) and [Other](kibana://reference/other.md).",
		},
		{
			name:     "no match leaves text unchanged",
			renderer: r,
			text:     "See https://example.com/unmapped for details.",
			want:     "See https://example.com/unmapped for details.",
		},
		{
			name:     "no mappings configured",
			renderer: &MarkdownRenderer{conf: &config.Config{}},
			text:     "See https://example.com/old for details.",
			want:     "See https://example.com/old for details.",
		},
		{
			name:     "nil config",
			renderer: &MarkdownRenderer{conf: nil},
			text:     "See https://example.com/old for details.",
			want:     "See https://example.com/old for details.",
		},
		{
			name:     "nil renderer",
			renderer: nil,
			text:     "See https://example.com/old for details.",
			want:     "See https://example.com/old for details.",
		},
		{
			name:     "multiple occurrences of the same URL are all rewritten",
			renderer: r,
			text:     "See https://example.com/old and again https://example.com/old.",
			want:     "See [New page](docs-content://new/page.md) and again [New page](docs-content://new/page.md).",
		},
		{
			// The shorter, prefix URL is declared first; the longer/more specific
			// URL must still win. RewriteLinks sorts by descending URL length, so
			// the result is independent of config order.
			name: "longer URL wins over shorter prefix regardless of order",
			renderer: &MarkdownRenderer{conf: &config.Config{
				Render: config.RenderConfig{
					LinkMappings: []*config.LinkMapping{
						{URL: "https://example.com/old", Link: "docs-content://other.md", Text: "Other"},
						{URL: "https://example.com/old-page", Link: "docs-content://new/page.md", Text: "New page"},
					},
				},
			}},
			text: "See https://example.com/old-page for details.",
			want: "See [New page](docs-content://new/page.md) for details.",
		},
		{
			// KNOWN LIMITATION: mappings target bare URLs. A URL that already
			// appears inside a Markdown link gets rewritten too, producing nested
			// (invalid) Markdown. This is acceptable because the feature is meant
			// for plain URLs in godoc; do not map a URL that is already linked.
			name:     "KNOWN LIMITATION: URL inside an existing Markdown link produces nested output",
			renderer: r,
			text:     "See [the page](https://example.com/old) for details.",
			want:     "See [the page]([New page](docs-content://new/page.md)) for details.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.renderer.RewriteLinks(tt.text))
		})
	}
}

func TestMarkdownRenderer_RenderFieldDoc_appliesLinkMappings(t *testing.T) {
	conf := &config.Config{
		Render: config.RenderConfig{
			LinkMappings: []*config.LinkMapping{
				{
					URL:  "https://example.com/old",
					Link: "docs-content://new/page.md",
					Text: "New page",
				},
			},
		},
	}
	r := &MarkdownRenderer{conf: conf}

	got := r.RenderFieldDoc("See https://example.com/old for details.")
	assert.Equal(t, "See [New page](docs-content://new/page.md) for details.", got)
}

func TestMarkdownRenderer_TemplateValue(t *testing.T) {
	tests := []struct {
		name     string
		renderer *MarkdownRenderer
		key      string
		want     string
	}{
		{
			name:     "existing key",
			renderer: &MarkdownRenderer{conf: newTestConfig(t)},
			key:      "foo",
			want:     "bar",
		},
		{
			name:     "another existing key",
			renderer: &MarkdownRenderer{conf: newTestConfig(t)},
			key:      "hello",
			want:     "world",
		},
		{
			name:     "missing key",
			renderer: &MarkdownRenderer{conf: newTestConfig(t)},
			key:      "missing",
			want:     "",
		},
		{
			name:     "nil config",
			renderer: &MarkdownRenderer{conf: nil},
			key:      "foo",
			want:     "",
		},
		{
			name:     "nil renderer",
			renderer: nil,
			key:      "foo",
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.renderer.TemplateValue(tt.key))
		})
	}
}
