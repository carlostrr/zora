package errparse

const kubescapeBigErrMsg = `failed to get resource: apps/v1, Resource=replicasets, namespace: , labelSelector: , reason: replicasets.apps is forbidden: User "system:anonymous" cannot list resource "replicasets" in API group "apps" at the cluster scope; failed to get resource: apps/v1, Resource=deployments, namespace: , labelSelector: , reason: deployments.apps is forbidden: User "system:anonymous" cannot list resource "deployments" in API group "apps" at the cluster scope; failed to get resource: /v1, Resource=pods, namespace: , labelSelector: , reason: pods is forbidden: User "system:anonymous" cannot list resource "pods" in API group "" at the cluster scope; failed to get resource: /v1, Resource=configmaps, namespace: , labelSelector: , reason: configmaps is forbidden: User "system:anonymous" cannot list resource "configmaps" in API group "" at the cluster scope; failed to get resource: /v1, Resource=services, namespace: , labelSelector: , reason: services is forbidden: User "system:anonymous" cannot list resource "services" in API group "" at the cluster scope; failed to get resource: batch/v1beta1, Resource=cronjobs, namespace: , labelSelector: , reason: cronjobs.batch is forbidden: User "system:anonymous" cannot list resource "cronjobs" in API group "batch" at the cluster scope; failed to get resource: apps/v1, Resource=daemonsets, namespace: , labelSelector: , reason: daemonsets.apps is forbidden: User "system:anonymous" cannot list resource "daemonsets" in API group "apps" at the cluster scope; failed to get resource: rbac.authorization.k8s.io/v1, Resource=rolebindings, namespace: , labelSelector: , reason: rolebindings.rbac.authorization.k8s.io is forbidden: User "system:anonymous" cannot list resource "rolebindings" in API group "rbac.authorization.k8s.io" at the cluster scope; failed to get resource: apps/v1, Resource=statefulsets, namespace: , labelSelector: , reason: statefulsets.apps is forbidden: User "system:anonymous" cannot list resource "statefulsets" in API group "apps" at the cluster scope; failed to get resource: rbac.authorization.k8s.io/v1, Resource=clusterrolebindings, namespace: , labelSelector: , reason: clusterrolebindings.rbac.authorization.k8s.io is forbidden: User "system:anonymous" cannot list resource "clusterrolebindings" in API group "rbac.authorization.k8s.io" at the cluster scope; failed to get resource: networking.k8s.io/v1, Resource=networkpolicies, namespace: , labelSelector: , reason: networkpolicies.networking.k8s.io is forbidden: User "system:anonymous" cannot list resource "networkpolicies" in API group "networking.k8s.io" at the cluster scope; failed to get resource: policy/v1beta1, Resource=podsecuritypolicies, namespace: , labelSelector: , reason: podsecuritypolicies.policy is forbidden: User "system:anonymous" cannot list resource "podsecuritypolicies" in API group "policy" at the cluster scope; failed to get resource: /v1, Resource=namespaces, namespace: , labelSelector: , reason: namespaces is forbidden: User "system:anonymous" cannot list resource "namespaces" in API group "" at the cluster scope; failed to get resource: /v1, Resource=serviceaccounts, namespace: , labelSelector: , reason: serviceaccounts is forbidden: User "system:anonymous" cannot list resource "serviceaccounts" in API group "" at the cluster scope; failed to get resource: rbac.authorization.k8s.io/v1, Resource=clusterroles, namespace: , labelSelector: , reason: clusterroles.rbac.authorization.k8s.io is forbidden: User "system:anonymous" cannot list resource "clusterroles" in API group "rbac.authorization.k8s.io" at the cluster scope; failed to get resource: batch/v1, Resource=jobs, namespace: , labelSelector: , reason: jobs.batch is forbidden: User "system:anonymous" cannot list resource "jobs" in API group "batch" at the cluster scope; failed to get resource: admissionregistration.k8s.io/v1, Resource=validatingwebhookconfigurations, namespace: , labelSelector: , reason: validatingwebhookconfigurations.admissionregistration.k8s.io is forbidden: User "system:anonymous" cannot list resource "validatingwebhookconfigurations" in API group "admissionregistration.k8s.io" at the cluster scope; failed to get resource: rbac.authorization.k8s.io/v1, Resource=roles, namespace: , labelSelector: , reason: roles.rbac.authorization.k8s.io is forbidden: User "system:anonymous" cannot list resource "roles" in API group "rbac.authorization.k8s.io" at the cluster scope; failed to get resource: admissionregistration.k8s.io/v1, Resource=mutatingwebhookconfigurations, namespace: , labelSelector: , reason: mutatingwebhookconfigurations.admissionregistration.k8s.io is forbidden: User "system:anonymous" cannot list resource "mutatingwebhookconfigurations" in API group "admissionregistration.k8s.io" at the cluster scope; failed to get resource: /v1, Resource=nodes, namespace: , labelSelector: , reason: nodes is forbidden: User "system:anonymous" cannot list resource "nodes" in API group "" at the cluster scope`
