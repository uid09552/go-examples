#!/bin/bash
echo "common deploy functions"
kubectl version
kubectl apply -f $1
kubectl describe deployment $2
kubectl label --overwrite deployment $2 version=$RELEASE_ENVIRONMENTID
kubectl label --overwrite deployment $2 build="$BUILD_BUILDNUMBER"
echo "Deploymentdate"$RELEASE_DEPLOYMENT_STARTTIME
kubectl annotate --overwrite deployment $2 deploymenttime="$RELEASE_DEPLOYMENT_STARTTIME"
kubectl describe deployment $2
