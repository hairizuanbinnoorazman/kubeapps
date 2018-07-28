# Playing around with basic kubernetes

Install docker desktop on your computer. Docker now comes with Kubernetes which can be enabled

```bash
kubectl run nginx --image nginx
kubectl expose --port 80 --type NodePort deployments nginx
```
