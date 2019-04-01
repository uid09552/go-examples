# Resolve DNS of K8S Services
```
kubecl exec -it busybox -- nslookup kubernetes.default
kubecl exec -it busybox -- cat /etc/resolf.conf
kubectl get pods -n kube-system -l kube-dns
kubectl get endpoints kube-dns -n kube-system
```
