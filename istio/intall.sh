curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.3.1 sh -
kubectl apply -f install/kubernetes/istio-demo.yaml