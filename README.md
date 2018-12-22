# TODO

1) I've run into the cert bootstraping problem @alvaroaleman mentioned here:

https://github.com/kubernetes-sigs/cluster-api/issues/147#issuecomment-430349301

2) Our RBAC rules are not being generated correctly. To work around this for now:

```
kubectl create clusterrolebinding cluster-api-superpowers --clusterrole=cluster-admin --user=system:serviceaccount:cluster-api-system:default
kubectl create clusterrolebinding cluster-api-provider-kubeadm-superpowers --clusterrole=cluster-admin --user=system:serviceaccount:cluster-api-provider-kubeadm-system:default
```

3) To deploy a sample Cluster and Machine to observe the webhook in action:

```
kubectl apply -f config/samples/cluster_v1alpha1_cluster.yaml 
kubectl apply -f config/samples/cluster_v1alpha1_machine.yaml 
```

Obviously this will not work unless 1) above is fixed (or hacked around
as documented in the comment...).
