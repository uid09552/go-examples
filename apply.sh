#!/bin/bash
echo "deploy kafka"
cd kafka/no-persist
kubectl apply -f kafka-deploy.yml
kubectl apply -f kafka-service.yml
kubectl apply -f zoo-deploy.yml
cd ../..
