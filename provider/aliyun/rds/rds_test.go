package rds_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/aliyun/connectivity"

	op "github.com/infraboard/cmdb/provider/aliyun/rds"
)

var (
	operater *op.RdsOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	pager := operater.PageQuery(req)

	hasNext := true
	for hasNext {
		p := pager.Next()
		if p.Err != nil {
			panic(p.Err)
		}
		hasNext = p.HasNext
		fmt.Println(p.Data)
	}
}

func init() {
	var ak, sk string
	if ak = os.Getenv("AL_CLOUD_ACCESS_KEY"); ak == "" {
		panic("empty AL_CLOUD_ACCESS_KEY")
	}

	if sk = os.Getenv("AL_CLOUD_ACCESS_SECRET"); sk == "" {
		panic("empty AL_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAliCloudClient(ak, sk, "cn-hangzhou")

	ec, err := client.RdsClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewRdsOperater(ec)
}
