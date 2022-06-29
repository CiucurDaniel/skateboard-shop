# catalog microservice

```
app config

applogger init

get router

setup server

 graceful shutdown

listen and serve
```

# Docker

App

```
docker build -t skateshop-catalog-microservice:latest .

http://localhost:3030/

docker login

docker push <hub-user>/<repo-name>:<tag>

docker push ciucurdaniel/skateshop-catalog-microservice:latest

docker run -it -p 3030:3030 skateshop-catalog-microservice

```

DB

```
docker run --name dev-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-password -d mysql:latest
```

Db has to be created as well for now

```
CREATE DATABASE catalogdb;
```