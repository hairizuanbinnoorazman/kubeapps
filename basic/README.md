# Playing around with basic kubernetes

- [Basic commands](#basic-commands)
- [Running nginx](#running-nginx)

## Basic commands

```
kubectl get pods
kubectl get services
kubectl get configmaps
kubectl get secrets
kubectl get deployments

kubectl describe deployments nginx
```

## Running nginx

Install docker desktop on your computer. Docker now comes with Kubernetes which can be enabled

```bash
kubectl run nginx --image nginx
kubectl expose --port 80 --type NodePort deployments nginx
```

**Cleaning up**

Removing the docker services

```bash
kubectl delete services nginx
kubectl delete deployments nginx
```

## Grab data from custom registry and running it

Create a temporary docker registry to try things out locally  
https://hub.docker.com/_/registry/

```bash
# Create a fake registry and then store the result in the registry
docker run -d -p 5000:5000 --name registry registry:2
docker image tag nginx localhost:5000/myfirstimage
docker push localhost:5000/myfirstimage
```

Request that kubernetes pull from the fake registry

```bash
kubectl run localnginx  --image localhost:5000/myfirstimage
kubectl expose --port 80 --type NodePort deployments localnginx
```

**Cleaning up**

```bash
kubectl delete services localnginx
kubectl delete deployments localnginx
docker stop regisry
docker rm registry
```
