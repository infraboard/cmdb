package bss_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/aliyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/aliyun/bss"
)

var (
	operater *op.BssOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	req.Month = "2021-10"

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
	zap.DevelopmentSetup()

	var secretID, secretKey string
	if secretID = os.Getenv("AL_CLOUD_ACCESS_KEY"); secretID == "" {
		panic("empty AL_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("AL_CLOUD_ACCESS_SECRET"); secretKey == "" {
		panic("empty AL_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAliCloudClient(secretID, secretKey, "cn-zhangjiakou")

	ec, err := client.BssClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewBssOperater(ec)
}
