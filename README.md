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

Apply everything by running

```
kubectl apply -f kubernetes
```

You can watch all pods coming alive with

```
kubectl get pods -w
```

To find the application URL exposed in our cluster by the service, type

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

## Production environment

To do

## Endpoints

### POST /users

#### Request

```bash
curl -XPOST http://localhost:3000/v1/users --header "Content-Type: application/json"  --data '{
    "name": "Jack",
    "email": "jack@email.com"
}'
```

#### Response

```json
{
  "id": "5fd6ac5c6884b412d6ec1475",
  "name": "Jack",
  "email": "jack@email.com",
  "updated_at": 1607904348,
  "created_at": 1607904348
}
```

### GET /users/:id

#### Request

```bash
curl http://localhost:3000/v1/users/1
```

#### Response

```json
{
  "id": "5fd6ac5c6884b412d6ec1475",
  "name": "Jack",
  "email": "jack@email.com",
  "updated_at": 1607904348,
  "created_at": 1607904348
}
```

### POST /short_urls

#### Request

```bash
curl -XPOST http://localhost:3000/v1/short_urls --header "Content-Type: application/json"  --data '{
    "original_url": "https://really-long-website-url.com",
    "user_id": "5fd6ac5c6884b412d6ec1475",
    "expires_at": 100
}'
```

#### Response

```json
{
  "hash": "ckinsrq9400000108qln5h3ah",
  "original_url": "https://really-long-website-url.com",
  "user_id": "5fd6ac5c6884b412d6ec1475",
  "expires_at": 100,
  "created_at": 1607904409
}
```

### GET /short_urls/:hash

#### Request

```bash
curl http://localhost:3000/v1/short_urls/ckinsrq9400000108qln5h3ah
```

#### Response

```json
{ "original_url": "https://really-long-website-url.com" }
```
