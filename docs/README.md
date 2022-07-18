



## MySQL安装

```sh
$ docker run -p 3306:3306 -itd -e MARIADB_USER=cmdb -e MARIADB_PASSWORD=123456 -e MARIADB_ROOT_PASSWORD=123456 --name mysql   mariadb:latest
```

详细文档请参考: [mariadb docker hub](https://hub.docker.com/_/mariadb)



Resource


// 创建一个部署
1. Privider (阿里云/腾讯云/华为云/IDC)


// 计算
ECS     instance         username   password



配置

// 存储
Rds       instance  db    username   password
Redis     instance  db    username   password
Mongodb   instance  db    username   password


Bucket              bucket client_id client_secret

// 功能
Sms                 client_id client_secret
Vms                 client_id client_secret