# TODO

Our RBAC rules are not being generated correctly. To work around this for now:

```
kubectl create clusterrolebinding cluster-api-superpowers --clusterrole=cluster-admin --user=system:serviceaccount:cluster-api-system:default
kubectl create clusterrolebinding cluster-api-provider-kubeadm-superpowers --clusterrole=cluster-admin --user=system:serviceaccount:cluster-api-provider-kubeadm-system:default

```

To deploy a sample Cluster and Machine to observe the webhook in action:

```
kubectl apply -f config/samples/cluster_v1alpha1_cluster.yaml 
kubectl apply -f config/samples/cluster_v1alpha1_machine.yaml 
```
