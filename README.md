# go-k8s


Run Docker : docker build -t dachanh/go-deployments .
Start Minikube : minikube start 
kubectl create -f deployment.yaml
kubectl expose deployment go-app-simple --type=NodePort --name=go-app-svc



Note : 
Run locally docker images , you should set :
 minikube docker-env 
 eval $(minikube -p minikube docker-env)