# Upgrade nodes
```
kubectl get nodes
kubectl drain $node --ignore-deamonsets
kubectl delete node $node
kubeadm token list
#create new token
kubeadm token generate
kubeadm token create $token --ttl 3h --discovery-token-ca-cert-hash sha256$hash
kubeadm join --token --discovery-token-ca-cert-hash sha256$hash

```
