{{- define "go-helm-app.fullname" -}}
{{ .Release.Name }}
{{- end }}

{{- define "go-helm-app.name" -}}
{{ .Chart.Name }}
{{- end }}
