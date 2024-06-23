# Distributed systems summer 2024 lab

This is the Lab for the course `Parallele und Verteilte Systeme` in summer 2024.\
For this, a backend was implemented according to specs provided in `openapi-spec.json`. The code is written in `go` using a `MongoDB` databse service.\
The frontend component used here is not of our own creation and was provided by the lecturers of this course. The image can be found on [Dockerhub](https://hub.docker.com/r/maeddes/todoui-thymeleaf).

## Starting the application

To start the application, simply clone the repositoy and navigate to its root directory

```bash
cd path/to/the/repository
```

and then start up the application.

```bash
docker compose up --build
```

The applications frontend will be available at `http://localhost.com:8090`.

## Docker image

If you wish to just use the docker image of the implemented backend, you can pull it from dockerhub using

```bash
docker pull gtorfo/todobackend:latest
```

or use it in your `docker-compose.yml` directly (excerpt shown):

```yaml
services:
  backend:
    image: gtorfo/todobackend:latest
```

## Documentation

Documentation on the API endpoints can be found under `http://localhost.com:8080/swagger/index.html`.

## Kubernetes setup

### Prerequisites

To deploy this application as a Kubernetes cluster, two prerequisites are needed:\
`Minikube`, a tool that allows one to run a local Kubernetes cluster and\
`kubectl`, a command line tool needed to interact with the Kubernetes cluster created by `Minikube`.

To install `Minikube`, run the following commands:
```bash
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube /usr/local/bin/
rm minikube
```
More information about `Minikube` can be found [here](https://minikube.sigs.k8s.io/docs/).

To install `kubectl`, run the following commands:
```bash
curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/
```
More information about `kubectl` can be found [here](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands).

### Deployment

To simplify the deployment process, a script is used to automate the setup. Navigate to the projects root directory and execute the following command:
```bash
./cluster_setup.sh
```
This script
 - starts a Kubernetes cluster using `Minikube`
 - applies the Kubernetes manifests using `kubectl` and
 - forwards all necessary ports.

 After this, the applications frontend can be accessed on port 8090, and the backend can be accessed on port 8080 as usual.
