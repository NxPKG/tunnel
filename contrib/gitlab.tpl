{{- /* Template based on https://docs.gitlab.com/ee/user/application_security/container_scanning/#reports-json-format */ -}}
{
  "version": "14.0.6",
  "vulnerabilities": [
  {{- $t_first := true }}
  {{- range . }}
  {{- $target := .Target }}
    {{- $image := $target | regexFind "[^\\s]+" }}
    {{- range .Vulnerabilities -}}
    {{- if $t_first -}}
      {{- $t_first = false -}}
    {{ else -}}
      ,
    {{- end }}
    {
      "id": "{{ .VulnerabilityID }}",
      "category": "container_scanning",
      "message": {{ .Title | printf "%q" }},
      "description": {{ .Description | printf "%q" }},
      {{- /* cve is a deprecated key, use id instead */}}
      "cve": "{{ .VulnerabilityID }}",
      "severity": {{ if eq .Severity "UNKNOWN" -}}
                    "Unknown"
                  {{- else if eq .Severity "LOW" -}}
                    "Low"
                  {{- else if eq .Severity "MEDIUM" -}}
                    "Medium"
                  {{- else if eq .Severity "HIGH" -}}
                    "High"
                  {{- else if eq .Severity "CRITICAL" -}}
                    "Critical"
                  {{-  else -}}
                    "{{ .Severity }}"
                  {{- end }},
      "solution": {{ if .FixedVersion -}}
                    "Upgrade {{ .PkgName }} to {{ .FixedVersion }}"
                  {{- else -}}
                    "No solution provided"
                  {{- end }},
      "scanner": {
        "id": "tunnel",
        "name": "tunnel"
      },
      "location": {
        "dependency": {
          "package": {
            "name": "{{ .PkgName }}"
          },
          "version": "{{ .InstalledVersion }}"
        },
        {{- /* TODO: No mapping available - https://github.com/khulnasoft/tunnel/issues/332 */}}
        "operating_system": "Unknown",
        "image": "{{ $image }}"
      },
      "identifiers": [
        {
	  {{- /* TODO: Type not extractable - https://github.com/khulnasoft-lab/vul-db/pull/24 */}}
          "type": "cve",
          "name": "{{ .VulnerabilityID }}",
          "value": "{{ .VulnerabilityID }}",
          "url": "{{ .PrimaryURL }}"
        }
      ],
      "links": [
        {{- $l_first := true -}}
        {{- range .References -}}
        {{- if $l_first -}}
          {{- $l_first = false }}
        {{- else -}}
          ,
        {{- end -}}
        {
          "url": "{{ regexFind "[^ ]+" . }}"
        }
        {{- end }}
      ]
    }
    {{- end -}}
  {{- end }}
  ],
  "remediations": []
}
