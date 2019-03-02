
[![Build Status](https://dev.azure.com/maxrgbg/maxrgbg/_apis/build/status/uid09552.go-examples?branchName=master)](https://dev.azure.com/maxrgbg/maxrgbg/_build/latest?definitionId=1&branchName=master)
## Example Service written in go.
Application is used to test a K8S Cluster.
CI-Steps
1. Build Docker image
2. Push Docker image to hub.docker.com
3. Deploy Image with rolling upgrade to K8S

The "main" app is called from outsite the cluster. The main app is then calling a K8S Service App "Frontend". The "Frontend" App is then going to call a second app "Backend" via localhost, because they are deployed as a Pod.

