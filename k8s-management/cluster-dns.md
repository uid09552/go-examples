# Resolve DNS of K8S Services

```
kubecl exec -it busybox -- nslookup kubernetes.default
kubecl exec -it busybox -- cat /etc/resolf.conf
kubectl get pods -n kube-system -l kube-dns
kubectl get endpoints kube-dns -n kube-system
```


# Change Cluster DNS

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-dns
  namespace: kube-system
data:
  upstreamNameservers: |
    ["172.16.0.1"]
```

# Pod DNS (overwrite cluster dns)
```
apiVersion: v1
kind: Pod
metadata:
  namespace: default
  name: dns-example
spec:
  containers:
    - name: test
      image: nginx
  dnsPolicy: "None"
  dnsConfig:
    nameservers:
      - 1.2.3.4
    searches:
      - ns1.svc.cluster.local
      - my.dns.search.suffix
    options:
      - name: ndots
        value: "2"
      - name: edns0
```
# Not using kube-dns only custom
```
apiVersion: v1
kind: Pod
metadata:
  name: busyboxdns
  labels:
    name: busybox
spec:
  hostname: busybox-2
  subdomain: default-subdomain
  containers:
  - image: busybox:1.28
    command:
      - sleep
      - "3600"
    name: busybox
  dnsConfig:
    nameservers:
      - 8.8.8.8
  dnsPolicy: None
```