{{- define "chart.fullname" -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Namespace $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "chart.volume" -}}
{{- if .Values.configmap.enabled -}}
volumes:
- name: {{ include "chart.fullname" . }}-volume
  configMap:
    name: {{ include "chart.fullname" . }}-configmap
{{- end -}}
{{- if .Values.sealedsecret.enabled }}
- name: {{ include "chart.fullname" . }}-sealedsecret-volume
  secret:
    secretName: {{ include "chart.fullname" . }}-sealedsecret
{{- end }}
{{- end -}}

{{- define "chart.volume.mount" -}}
{{- if .Values.configmap.enabled -}}
volumeMounts:
- mountPath: {{ .Values.configmap.configPath }}/{{ .Values.configmap.file_name_app }}
  subPath: {{ .Values.configmap.file_name_app }}
  name: {{ include "chart.fullname" . }}-volume
{{- end -}}
{{- if .Values.sealedsecret.enabled }}
- mountPath: {{ .Values.sealedsecret.configPath }}/{{ .Values.sealedsecret.filename }}
  subPath: {{ .Values.sealedsecret.filename }}
  name: {{ include "chart.fullname" . }}-sealedsecret-volume
  readOnly: true
{{- end }}
{{- end -}}
