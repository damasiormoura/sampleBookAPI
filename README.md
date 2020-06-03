# Sample Booking API

Proof of Concept of a go service deployed in Kubernetes (Cloud Native).

## Prerequisites

In order to build and run this showcase you need to have a couple of things installed:

* The Go SDK (e.g. using Brew)
* An IDE for the development, like Visual Studio Code or IntelliJ
* The Docker Toolbox or native Docker, whatever you prefer

## Building the Go microservices

```bash
$ go build
$ ./samplebooksapi.exe
```

## Containerization with Docker

```bash
$ GOOS=linux GOARCH=amd64 go build
$ docker build -t samplebooksapi:1.0.0 .
$ docker images
$ docker tag samplebooksapi:1.0.0 damasiormoura/samplebooksapi:1.0.0
$ docker push damasiormoura/samplebooksapi:1.0.0
```

## Running Docker image locally

```bash
$ docker run -it -p 8080:8080 samplebooksapi:1.0.0
$ docker run -it -e "PORT=9090" -p 9090:9090 samplebooksapi:1.0.0
$ docker ps --all

$ docker run --name=cloud-native-go -d -p 8080:8080 samplebooksapi:1.0.0
$ docker ps
$ docker stats
$ docker stop
$ docker start
```

## Docker Compose

```bash
$ docker-compose build
$ docker images
$ docker-compose up -d
$ docker-compose kill
$ docker-compose rm
```

== Kubernetes and Pods

```bash
$ kubectl get pods
$ kubectl create -f k8s-pod.yml
$ kubectl get pods
$ kubectl describe pod sample-books

$ kubectl port-forward sample-books 8080:8080

$ kubectl get pods --show-labels
$ kubectl get pods -o wide -L env
$ kubectl label pod sample-books hello=world
$ kubectl get pods -o wide -L hello
$ kubectl label pod sample-books env=prod --overwrite
$ kubectl get pods -o wide -L env

$ kubectl get ns
$ kubectl get pods --namespace kube-system
$ kubectl create -f k8s-namespace.yml
$ kubectl get ns
$ kubectl create -f k8s-pod.yml --namespace sample-books

$ kubectl delete -f k8s-pod.yml
```

== Kubernetes Deployments and Services

```bash
$ kubectl get services,deployments,pods

$ kubectl create -f k8s-deployment.yml
$ kubectl get deployments,pods

$ kubectl apply -f k8s-deployment.yml

$ kubectl create -f k8s-service.yml
$ kubectl get services
$ kubectl describe service sample-books

$ kubectl delete -f k8s-deployment.yml
$ kubectl delete -f k8s-service.yml
```

== Kubernetes Scaling and Rolling Updates

```bash
$ kubectl create -f k8s-deployment.yml --record=true

$ kubectl scale deployment sample-books --replicas=5
$ kubectl scale deployment sample-books --replicas=3

$ kubectl rollout history deployment sample-books

$ kubectl apply -f k8s-deployment.yml

$ kubectl edit deployment sample-books

$ kubectl set image deployment sample-books sample-books=samplebooksapi:1.0.1

$ kubectl rollout history deployment sample-books
$ kubectl rollout undo deployment sample-books --to-revision=0

$ kubectl delete -f k8s-deployment.yml
```

## Maintainer and Author

Rodrigo Moura (@damasiormoura)


## Reference
Based on Linkedin course: https://www.linkedin.com/learning/getting-started-with-cloud-native-go

M.-Leander Reimer (@lreimer)
