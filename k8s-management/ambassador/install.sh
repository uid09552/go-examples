#kubectl create clusterrolebinding my-cluster-admin-binding --clusterrole=cluster-admin --user=$(gcloud info --format="value(config.account)")
kubectl apply -f https://getambassador.io/yaml/ambassador/ambassador-rbac.yaml
echo <<EOL
---
apiVersion: v1
kind: Service
metadata:
  name: ambassador
spec:
  type: NodePort
  ports:
  - name: http
    port: 8088
    targetPort: 8080
    nodePort: 30036  # Optional: Define the port you would like exposed
    protocol: TCP
  selector:
    service: ambassador

EOL | kubectl apply -f -