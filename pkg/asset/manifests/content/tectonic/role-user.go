package tectonic

const (
	// RoleUser  is the variable/constant representing the contents of the respective file
	RoleUser = `
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: user
rules:
  - apiGroups: [""]
    resources: [
      "bindings", "configmaps", "events", "pods", "replicationcontrollers",
      "secrets", "services", "serviceaccounts",
      "pods/attach",
      "pods/binding",
      "pods/exec",
      "pods/log",
      "pods/portforward",
      "pods/proxy",
      "pods/status",
      "replicationcontrollers/scale",
      "replicationcontrollers/status",
      "services/proxy",
      "services/status"
    ]
    verbs: ["*"]
    nonResourceURLs: []

  - apiGroups: [""]
    resources: [
      "componentstatuses", "endpoints", "limitranges", "nodes", "nodes/proxy", "nodes/status",
      "namespaces", "namespaces/status", "namespaces/finalize",
      "persistentvolumeclaims", "persistentvolumeclaims/status", "persistentvolumes", "resourcequotas",
      "resourcequotas/status"
    ]
    verbs: ["get", "list", "watch", "proxy", "redirect"]
    nonResourceURLs: []

  - apiGroups: ["apps", "batch", "autoscaling", "policy"]
    resources: ["*"]
    verbs: ["*"]
    nonResourceURLs: []

  - apiGroups: ["extensions"]
    resources: [
      "daemonsets", "deployments", "horizontalpodautoscalers", "ingresses",
      "jobs", "replicasets", "replicationcontrollers",

      "daemonsets/status",
      "deployments/rollback",
      "deployments/scale",
      "deployments/status",
      "horizontalpodautoscalers/status",
      "ingresses/status",
      "jobs/status",
      "replicasets/scale",
      "replicasets/status",
      "replicationcontrollers/scale"
    ]
    verbs: ["*"]
    nonResourceURLs: []

  - apiGroups: ["extensions"]
    resources: ["networkpolicies", "thirdpartyresources"]
    verbs: ["get", "list", "watch", "proxy", "redirect"]
    nonResourceURLs: []

  - apiGroups: ["rbac.authorization.k8s.io"]
    resources: ["*"]
    verbs: ["get", "list", "watch", "proxy", "redirect"]
    nonResourceURLs: []
`
)
