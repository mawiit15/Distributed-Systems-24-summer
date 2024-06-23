#!/bin/bash

#function used to forward all the ports
forward_ports() {
    #database
    kubectl port-forward service/mongo-db-service 27017:27017 &
    MONGO_PID=$!

    #backend
    kubectl port-forward service/backend-service 8080:8080 &
    BACKEND_PID=$!

    #frontend
    kubectl port-forward service/frontend-service 8090:8090 &
    FRONTEND_PID=$!

    #wait until background processes are finished
    wait $MONGO_PID $BACKEND_PID $FRONTEND_PID
}

#clean kill when exiting
cleanup() {
    kill $MONGO_PID $BACKEND_PID $FRONTEND_PID
    exit 0
}

#trap --> call cleanup when exiting (signal interrupt / signal terminate)
trap cleanup SIGINT SIGTERM

#start minikube if not running already
if ! minikube status &> /dev/null; then
    minikube start --driver=docker
fi

#kubectl uses minikube cotext
kubectl config use-context minikube

#apply k8s configs
kubectl apply -f ./k8s

#forward the ports
forward_ports

#infinite loop, keep processes alive
while :; do 
    sleep 1; 
done
