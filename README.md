# URL Shortener

A highly scalable URL shortener API written in golang. 

## Tech Stack
- Golang
- MongoDB
- Redis
- Terraform
- EKS


## Local environment

### Run with docker-compose 

For local development, you can start everything up using docker-compose by running

```bash
make all
```

When it's done, three containers will be up and running: a Redis database for caching, a MongoDB as the main general database and the URL shortener API itself.

### Minikube

You can also run the application locally with kubernetes using minikube.

First of all, start the minikube cluster

```
minikube start
```

Apply everything with

```
kubectl apply -f kubernetes
```

You can watch all pods coming alive with

```
kubectl get pods -w
```

To find the application URL exposed in our cluster by the service, you can go with

```
minikube service url-shortener
```


When you're done, delete the application with

```
kubectl delete -f kubernetes
```

And also delete the minikube cluster

```
minikube delete --all
```


## Production envinroment

To do