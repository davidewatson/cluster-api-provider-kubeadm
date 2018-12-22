# TODO

Our RBAC rules are not being generated correctly. To work around this for now:

```
kubectl create clusterrolebinding cluster-api-superpowers --clusterrole=cluster-admin --user=system:serviceaccount:cluster-api-system:default
kubectl create clusterrolebinding cluster-api-provider-kubeadm-superpowers --clusterrole=cluster-admin --user=system:serviceaccount:cluster-api-provider-kubeadm-system:default

```
