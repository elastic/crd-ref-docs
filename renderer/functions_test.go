package renderer

import (
	"testing"

	"github.com/elastic/crd-ref-docs/config"
	"github.com/elastic/crd-ref-docs/types"
	"github.com/stretchr/testify/require"
)

func TestKubernetesHelper(t *testing.T) {
	conf := config.Config{
		Render: config.RenderConfig{
			KubernetesVersion: "1.15",
		},
	}

	kh, err := newKubernetesHelper(&conf)
	require.NoError(t, err)

	link := kh.LinkForKubeType(&types.Type{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ObjectMeta"})
	require.Equal(t, "https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.15/#objectmeta-v1-meta", link)
}
