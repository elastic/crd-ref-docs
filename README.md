CRD Reference Documentation Generator
======================================

Generates API reference documentation by scanning a source tree for exported CRD types.

This is a fresh implementation inspired by the https://github.com/ahmetb/gen-crd-api-reference-docs project that addresses some of its shortcomings such as the lack of Go module support, slow scan times and strong coupling of rendering logic.


```
go run main.go \
    --source-path=$GOPATH/src/github.com/elastic/cloud-on-k8s/pkg/apis \
    --config=config.yaml \
    --renderer=asciidoctor \
    --templates-dir=templates/asciidoctor
```
