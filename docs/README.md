



## MySQL安装

```sh
$ docker run -p 3306:3306 -itd -e MARIADB_USER=cmdb -e MARIADB_PASSWORD=123456 -e MARIADB_ROOT_PASSWORD=123456 --name mysql   mariadb:latest
```

详细文档请参考: [mariadb docker hub](https://hub.docker.com/_/mariadb)


## 应用中间件

配置

// 存储
Rds       instance  db    username   password
Redis     instance  db    username   password
Mongodb   instance  db    username   password


Bucket              bucket client_id client_credential

// 功能
Sms                 client_id client_credential
Vms                 client_id client_credential