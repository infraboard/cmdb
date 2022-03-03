package ecs_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/aliyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/aliyun/ecs"
)

var (
	operater *op.EcsOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	pager := operater.PageQuery(req)
	for pager.HasNext() {
		p := pager.Next()
		if p.Err != nil {
			panic(p.Err)
		}
		fmt.Println(p.Data)
	}
}

func init() {
	zap.DevelopmentSetup()
	var ak, sk string
	if ak = os.Getenv("AL_CLOUD_ACCESS_KEY"); ak == "" {
		panic("empty AL_CLOUD_ACCESS_KEY")
	}

	if sk = os.Getenv("AL_CLOUD_ACCESS_SECRET"); sk == "" {
		panic("empty AL_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAliCloudClient(ak, sk, "cn-hangzhou")

	ec, err := client.EcsClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewEcsOperater(ec)
}
