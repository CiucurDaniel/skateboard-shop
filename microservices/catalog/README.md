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
docker build -t ciucurdaniel/skateshop-catalog-microservice:latest .

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

## Sometimes there are some problems with the DB

To solve use the following:

```
ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


mysql> GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'my-password' WITH GRANT OPTION;
Query OK, 0 rows affected (0.01 sec)

mysql> CREATE DATABASE catalogdb;
Query OK, 1 row affected (0.00 sec)

mysql> USE catalogdb;
```