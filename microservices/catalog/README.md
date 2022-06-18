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

docker run -it -p 3030:3030 skateshop-catalog-microservice

http://localhost:3030/

```

DB
```
docker run --name dev-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-password -d mysql:latest
```