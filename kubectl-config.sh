#!/bin/bash
echo $server
kubectl config set-cluster rancher --server=$VARIABLE_SERVER
kubectl config set-context default --cluster rancher --user admin-user
kubectl config set-credentials admin-user --token $VARIABLE_TOKEN
kubectl config use-context default
kubectl get nodes
#Rancher has own tokens
