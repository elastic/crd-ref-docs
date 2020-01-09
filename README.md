CRD Reference Documentation Generator
======================================

Generates API reference documentation by scanning a source tree for exported CRD types.

This is a fresh implementation inspired by the https://github.com/ahmetb/gen-crd-api-reference-docs project. While trying to adopt the `gen-crd-api-refernce-docs` to generate documentation for [Elastic Cloud on Kubernetes](https://github.com/elastic/cloud-on-k8s), we encountered a few shortcomings such as the lack of support for Go modules, slow scan times, and rendering logic that was hard to adapt to Asciidoc (our preferred documentation markup language). This project attempts to address those issues by re-implementing the type discovery logic and decoupling the rendering logic so that different markup formats can be supported.


Usage
-----

Pre-built Linux binaries can be downloaded from the Github Releases tab. Alternatively, you can download and build the source with Go tooling:

```
go get -u github.com/elastic/crd-ref-docs
```

The tool can be invoked as follows to generate documentation in Asciidoc format:

```
crd-ref-docs \
    --source-path=$GOPATH/src/github.com/elastic/cloud-on-k8s/pkg/apis \
    --config=config.yaml \
    --renderer=asciidoctor \
    --templates-dir=templates/asciidoctor
```
