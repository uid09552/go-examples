#!/bin/bash
echo "common deploy functions"
kubectl version
kubectl apply -f $1
kubectl label deployment $2 version=$RELEASE_ENVIRONMENTID
kubectl label deployment $2 deploymenttime=$RELEASE_DEPLOYMENT_STARTTIME
kubectl describe deployment $2
