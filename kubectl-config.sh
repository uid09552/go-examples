#!/bin/bash
echo $server
kubectl config set-cluster rancher --server=$server
kubectl config set-context default --cluster rancher --user admin
kubectl config set-credentials admin --token $token
kubectl config use-context default
