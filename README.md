# go-k8s


Run Docker : <b>docker build -t dachanh/go-deployments .</b>

Start Minikube :<b> minikube start </b> 

Create Pod : <b>kubectl create -f deployment.yaml</b>

Developments:<b> kubectl expose deployment go-app-simple --type=NodePort --name=go-app-svc</b>



Note : 
Run locally docker images , you should set :
 
<b> minikube docker-env </b> 
 
<b> eval $(minikube -p minikube docker-env)</b>
