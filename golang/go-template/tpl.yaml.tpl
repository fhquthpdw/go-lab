---
apiVersion: monitoring.coreos.com/v1
kind: PrometheusRule
metadata:
  annotations:
    prometheus-operator-validated: "true"
  labels:
    role: prometheus-rules
    repo: k8sresources
spec:
  groups:
      rules:
        - alert: HighHttpLatency
          annotations:
description: "HTTP requests average latency for uri {{"{{"}} $labels.uri {{"}}"}} too High in 5min\n  VALUE = {{"{{"}} $value {{"}}"}}."
dashboard: "https://devops.cdt.thelegogroup.cn/grafana/d/apimetrics/api-metrics?orgId=1&refresh=30s&var-datasource=prometheus&var-namespace={{"{{"}} $labels.namespace {{"}}"}}&var-application={{"{{"}} $labels.application {{"}}"}}&var-pod_name=All"
value: {{"\"{{"}} $value {{"}}\""}}
description: "HTTP requests average latency for uri {{"{{"}} $labels.uri {{"}}"}} too High in 5min\n  VALUE = {{"{{"}} $value {{"}}"}}."
dashboard: "https://devops.ppd.cdt.thelegogroup.cn/grafana/d/apimetrics/api-metrics?orgId=1&refresh=30s&var-datasource=prometheus&var-namespace={{"{{"}} $labels.namespace {{"}}"}}&var-application={{"{{"}} $labels.application {{"}}"}}&var-pod_name=All"
value: {{"\"{{"}} $value {{"}}\""}}
for: 2m
labels:
  severity: critical
  sloth_severity: teams
  threshold: ">3s / 2m"
