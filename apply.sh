#!/bin/bash
echo "deploy kafka"
cd kafka/no-persist
kubectl apply -f kafka-deploy.yml
kubectl apply -f kafka-service.yml
kubectl apply -f zoo-deploy.yml
cd ../..

echo "deploy openfaas"
cd openfaas
kubectl apply -f https://raw.githubusercontent.com/openfaas/faas-netes/master/namespaces.yml
kubectl apply -f ./yaml
cd ..