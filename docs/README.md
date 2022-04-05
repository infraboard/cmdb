



## MySQL安装

```sh
$ docker run -p 3306:3306 -itd -e MARIADB_USER=cmdb -e MARIADB_PASSWORD=123456 -e MARIADB_ROOT_PASSWORD=123456 --name mysql   mariadb:latest
```

详细文档请参考: [mariadb docker hub](https://hub.docker.com/_/mariadb)