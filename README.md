# cmdb
cmdb



## 开发环境

protobuf 环境准备
```sh
# 1.安装protoc编译器

# 2.protoc-gen-go 插件之前已经安装
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 3.安装protoc-gen-go-grpc插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 4.安装自定义proto tag插件
go install github.com/favadi/protoc-go-inject-tag@latest
```
