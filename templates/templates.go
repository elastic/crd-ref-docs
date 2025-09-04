package templates

import "embed"

//go:embed asciidoctor
//go:embed markdown
//go:embed markdown-x
var Root embed.FS
