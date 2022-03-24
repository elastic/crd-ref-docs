package processor

import (
	"testing"

	"github.com/elastic/crd-ref-docs/config"
	"github.com/stretchr/testify/require"
)

func TestCompiledConfig(t *testing.T) {
	conf := &config.Config{
		Processor: config.ProcessorConfig{
			IgnoreTypes:         []string{"typex$"},
			IgnoreFields:        []string{`mytype\.Fieldy$`},
			IgnoreGroupVersions: []string{"groupz/v1$"},
		},
	}

	cc, err := compileConfig(conf)
	require.NoError(t, err)

	t.Run("ignoreType", func(t *testing.T) {
		require.True(t, cc.shouldIgnoreType("mytypex"))
		require.False(t, cc.shouldIgnoreType("typexyz"))
	})

	t.Run("ignoreField", func(t *testing.T) {
		require.True(t, cc.shouldIgnoreField("mytype", "Fieldy"))
		require.False(t, cc.shouldIgnoreField("mytype", "Fieldyz"))
	})

	t.Run("ignoreGroupVersion", func(t *testing.T) {
		require.True(t, cc.shouldIgnoreGroupVersion("groupz/v1"))
		require.False(t, cc.shouldIgnoreGroupVersion("groupz/v1beta1"))
	})
}
