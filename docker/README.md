# build docker image
'''
sudo docker build .
sudo docker tag <imageid> <remote-tag>
'''
# create pod
kubectl create -f pod.yaml
# create deployment
kubectl create -f deployment.yaml
kubectl get deployments
