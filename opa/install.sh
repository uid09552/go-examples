kubectl delete configmap example-policy 
kubectl create configmap example-policy --from-file example.rego
kubectl apply -f deployment-opa.yaml
kubectl apply -f service-opa.yaml