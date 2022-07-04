# auth microservice

The shopping cart microservice which is responsible for add item to cart, storing cart items and processing an order.

## Docker build image

```
docker build -t ciucurdaniel/skateshop-auth:latest .

docker login

docker push <hub-user>/<repo-name>:<tag>

docker push ciucurdaniel/skateshop-auth:latest

docker run -p 8034:8070 ciucurdaniel/skateshop-auth:latest
```