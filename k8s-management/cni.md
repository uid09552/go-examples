# Overview
* All Pods can communicatate with each other
* Each pods has its own IP Address
* no need for mapping of container ports

# CNI is responsible for
* Inserting the Network interface to Containter Namespace (veth pair)
* making neccessary changes to host (attach to bridge e.g.)
* Assign IP Address to the Network Interface 
* Setup Routes consistent with IP Address management

# Kubelet
* Default network plugin
* probes on startup for plugins

## Links
https://www.level-up.one/kubernetes-networking-3-level-up/

# Flannel
* flanneld
* VXLan
* Subents per host
* no network policies but with calico
