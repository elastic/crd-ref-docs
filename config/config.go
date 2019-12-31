package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Processor ProcessorConfig `json:"processor"`
	Render    RenderConfig    `json:"render"`
	Flags     `json:"-"`
}

type ProcessorConfig struct {
	MaxDepth     int      `json:"maxDepth"`
	IgnoreTypes  []string `json:"ignoreTypes"`
	IgnoreFields []string `json:"ignoreFields"`
}

type RenderConfig struct {
	KubernetesVersion string `json:"kubernetesVersion"`
}

type Flags struct {
	Config       string
	LogLevel     string
	OutputPath   string
	Renderer     string
	SourcePath   string
	TemplatesDir string
	MaxDepth     int
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
	return &conf, nil
}
