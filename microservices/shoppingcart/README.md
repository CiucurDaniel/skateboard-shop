# shopping-cart microservice

The shopping cart microservice which is responsible for add item to cart, storing cart items and processing an order.

## Docker build image

```
docker build -t ciucurdaniel/skateshop-shoppingcart:latest .

docker login

docker push <hub-user>/<repo-name>:<tag>

docker push ciucurdaniel/skateshop-shoppingcart:latest

docker run -p 8033:8060 ciucurdaniel/skateshop-shoppingcart:latest
```