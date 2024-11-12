{{- define "gvDetails" -}}
{{- $gv := . -}}

## <a id="{{ markdownGroupVersionID $gv | markdownSafeID }}">{{ $gv.GroupVersionString }}</a>

{{ $gv.Doc }}

{{- if $gv.Kinds  }}
### Resource Types
{{- range $gv.SortedKinds }}
- {{ $gv.TypeForKind . | markdownRenderTypeLink }}
{{- end }}
{{ end }}

{{ range $gv.SortedTypes }}
{{ template "type" . }}
{{ end }}

{{- end -}}
