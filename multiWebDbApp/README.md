# Multi Web applications with database attached

This set of applcations consists of the following:

- A Golang Web Application
  - Contains database storage component (data stored in MySQL)
- A MySQL DB
  - Maintained by Kubernetes
- A Python Web Application
  - Contains database storage component (data stored in Redis)
- A Redis DB
  - Contains Redis storage component
- A ConfigMaps to contain certain configuration data for the Golang and Python applications
- A Secrets to contain secret data to connect Golang and Python to redis

# Information on kubernetes volumes

Kubernetes volumes worked quite different from normal dockerfiles.

- In a deployment pod
  - Specify the volumes via the usage of persistant volume claims
- Specify a persistent volume claim
- Specify a persistent volume

Locally, it has to be manually done for now, no dynamic provisioning of storage is available locally

```bash
kubectl apply -f pv.yaml
kubectl apply -f mysql.yaml
kubectl expose --port 3306 --type NodePort deployments mysql
```

To clean up

```bash
kubectl delete deployments mysql
kubectl delete pvc mysql-pv-claim
kubectl delete pv mysql-pv
kubectl delete pv mysql-pv-2
```

Note during mounting, each application is required to have the data to be loaded on the appropriate folder. Do check each application to view where the folder has to be loaded up

https://estl.tech/deploying-redis-with-persistence-on-google-kubernetes-engine-c1d60f70a043
