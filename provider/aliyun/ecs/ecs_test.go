package ecs_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/aliyun/connectivity"

	op "github.com/infraboard/cmdb/provider/aliyun/ecs"
)

var (
	operater *op.EcsOperater
)

func TestQuery(t *testing.T) {
	pager := operater.PageQuery()

	hasNext := true
	for hasNext {
		p := pager.Next()
		hasNext = p.HasNext
		fmt.Println(p.Data)
	}

}

func init() {
	var secretID, secretKey string
	if secretID = os.Getenv("AL_CLOUD_ACCESS_KEY"); secretID == "" {
		panic("empty AL_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("AL_CLOUD_ACCESS_SECRET"); secretKey == "" {
		panic("empty AL_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAliCloudClient(secretID, secretKey, "cn-hangzhou")

	ec, err := client.EcsClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewEcsOperater(ec)
}
