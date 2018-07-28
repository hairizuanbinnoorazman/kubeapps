# Playing around with basic kubernetes

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

## Cleaning up

Removing the docker services

```bash
kubectl delete services nginx
kubectl delete deployments nginx
```
