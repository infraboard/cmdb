# cmdb

多云资产管理平台, 支持厂商:
+ 阿里云
+ 腾讯云
+ 华为云
+ AWS
+ VMware

## SDK使用

```go
package main

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/client"
	"github.com/infraboard/cmdb/app/resource"
)

func main() {
    // 配置cmdb grpc服务调用地址和凭证
	conf := client.NewConfig("localhost:18060")
	conf.WithClientCredentials("xx", "xx")

    // 创建CMDB客户端
	cmdb, err := client.NewClient(conf)
	if err != nil {
		panic(err)
	}

    // 服务调用
	rs, err := cmdb.Resource().Search(context.Background(), resource.NewSearchRequest())
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)
}
```

## 开发环境

grpc 环境准备
```sh
# 1.安装protoc编译器,  项目使用版本: v3.19.1
# 下载预编译包安装: https://github.com/protocolbuffers/protobuf/releases

# 2.protoc-gen-go go语言查询, 项目使用版本: v1.27.1   
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 3.安装protoc-gen-go-grpc插件, 项目使用版本: 1.1.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 4.安装自定义proto tag插件
go install github.com/favadi/protoc-go-inject-tag@latest
```

运行程序:
```sh
make run
```

protobuf代码生成:
```
make gen
```



