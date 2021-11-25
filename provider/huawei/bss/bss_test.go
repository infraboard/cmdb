package bss_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/huawei/connectivity"

	op "github.com/infraboard/cmdb/provider/huawei/bss"
)

var (
	operater *op.BssOperater
)

func TestQuery(t *testing.T) {
	pager := operater.PageQuery()

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
	var secretID, secretKey string
	if secretID = os.Getenv("HW_CLOUD_ACCESS_KEY"); secretID == "" {
		panic("empty HW_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("HW_CLOUD_ACCESS_SECRET"); secretKey == "" {
		panic("empty HW_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewHuaweiCloudClient(secretID, secretKey, "cn-north-4")

	ec, err := client.BssClient()
	if err != nil {
		panic(err)
	}

	operater = op.NewBssOperater(ec)
}
