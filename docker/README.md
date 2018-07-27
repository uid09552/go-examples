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

# do rolling upgrade
kubectl replace -f deployment.yaml --record