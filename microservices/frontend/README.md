# frontend

the frontend component from `Skateboard shop`.


# docker steps

The component needs to be packaged into an image.
Then pushed to docker registry. Images are pushed into my personal public registry in order to avoid the need for `imagePullSecrets` in minikube cluster.

```
docker build -t ciucurdaniel/skateshop-frontend:latest .

docker login

docker push <hub-user>/<repo-name>:<tag>

docker push ciucurdaniel/skateshop-frontend:latest

docker run -p 8080:9000 ciucurdaniel/skateshop-frontend:latest
```
