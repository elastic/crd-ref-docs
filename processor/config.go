package processor

import (
	"fmt"
	"regexp"

	"github.com/elastic/crd-ref-docs/config"
)

func compileConfig(conf *config.Config) (cc *compiledConfig, err error) {
	if conf == nil {
		return nil, nil
	}

	cc = &compiledConfig{
		ignoreTypes:  make([]*regexp.Regexp, len(conf.Processor.IgnoreTypes)),
		ignoreFields: make([]*regexp.Regexp, len(conf.Processor.IgnoreFields)),
	}

	for i, t := range conf.Processor.IgnoreTypes {
		if cc.ignoreTypes[i], err = regexp.Compile(t); err != nil {
			return nil, fmt.Errorf("failed to compile type regex '%s': %w", t, err)
		}
	}

	for i, f := range conf.Processor.IgnoreFields {
		if cc.ignoreFields[i], err = regexp.Compile(f); err != nil {
			return nil, fmt.Errorf("failed to compile field regex '%s': %w", f, err)
		}
	}

	return
}

type compiledConfig struct {
	ignoreTypes  []*regexp.Regexp
	ignoreFields []*regexp.Regexp
}

func (cc *compiledConfig) shouldIgnoreType(fqn string) bool {
	if cc == nil {
		return false
	}

	for _, re := range cc.ignoreTypes {
		if re.MatchString(fqn) {
			return true
		}
	}

	return false
}

func (cc *compiledConfig) shouldIgnoreField(typeName, fieldName string) bool {
	if cc == nil {
		return false
	}

	if fieldName == "-" {
		return true
	}

	fqn := typeName + "." + fieldName
	for _, re := range cc.ignoreFields {
		if re.MatchString(fqn) {
			return true
		}
	}

	return false
}
