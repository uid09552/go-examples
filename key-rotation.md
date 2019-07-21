https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/#rotating-a-decryption-key
kubectl get secrets --all-namespaces -o json | kubectl replace -f -
