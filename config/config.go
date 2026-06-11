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

package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Processor ProcessorConfig `json:"processor"`
	Render    RenderConfig    `json:"render"`
	Flags     `json:"-"`
}

// NamingConvention identifies a field-naming style to derive as an alias
// when a struct field carries the json `case:ignore` tag option.
type NamingConvention string

const (
	// NamingConventionCamelCase derives the lowerCamelCase form (e.g. "groupWait" for "group_wait").
	NamingConventionCamelCase NamingConvention = "camelCase"
	// NamingConventionSnakeCase derives the snake_case form (e.g. "group_wait" for "groupWait").
	NamingConventionSnakeCase NamingConvention = "snake_case"
)

type ProcessorConfig struct {
	MaxDepth            int      `json:"maxDepth"`
	IgnoreTypes         []string `json:"ignoreTypes"`
	IgnoreFields        []string `json:"ignoreFields"`
	IgnoreGroupVersions []string `json:"ignoreGroupVersions"`
	UseRawDocstring     bool     `json:"useRawDocstring"`
	CustomMarkers       []Marker `json:"customMarkers"`
	// CaseIgnoreAliases lists the naming conventions to derive and display as
	// alternative field names whenever a struct field carries the json `case:ignore`
	// tag option. When empty, no aliases are shown for such fields.
	CaseIgnoreAliases []NamingConvention `json:"caseIgnoreAliases"`
}

type Marker struct {
	Name   string
	Target TargetType
}

type TargetType string

const (
	TargetTypePackage TargetType = "package"
	TargetTypeType    TargetType = "type"
	TargetTypeField   TargetType = "field"
)

type RenderConfig struct {
	KnownTypes        []*KnownType   `json:"knownTypes"`
	KubernetesVersion string         `json:"kubernetesVersion"`
	LinkMappings      []*LinkMapping `json:"linkMappings"`
}

type KnownType struct {
	Name    string `json:"name"`
	Package string `json:"package"`
	Link    string `json:"link"`
}

type LinkMapping struct {
	URL  string `json:"url"`
	Link string `json:"link"`
	Text string `json:"text"`
}

const (
	OutputModeSingle = "single"
	OutputModeGroup  = "group"
)

type Flags struct {
	Config            string
	LogLevel          string
	OutputPath        string
	Renderer          string
	SourcePath        string
	TemplatesDir      string
	OutputMode        string
	MaxDepth          int
	TemplateKeyValues KeyValueFlags
}

func Load(flags Flags) (*Config, error) {
	f, err := os.Open(flags.Config)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	var conf Config
	if err := decoder.Decode(&conf); err != nil {
		return nil, err
	}

	conf.Flags = flags

	for i, lm := range conf.Render.LinkMappings {
		if lm.URL == "" || lm.Link == "" || lm.Text == "" {
			return nil, fmt.Errorf("render.linkMappings[%d]: url, link, and text are all required", i)
		}
	}

	return &conf, nil
}
