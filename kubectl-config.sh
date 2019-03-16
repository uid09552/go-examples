#!/bin/bash
echo $server
kubectl config set-cluster rancher --server=$VARIABLE_SERVER
kubectl config set-context default --cluster rancher --user admin
kubectl config set-credentials admin --token $VARIABLE_TOKEN
kubectl config use-context default
