package rds_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"

	op "github.com/infraboard/cmdb/provider/huawei/rds"
)

var (
	operater *op.RdsOperater
)

func TestQuery(t *testing.T) {
	pager := operater.PageQuery()

	for pager.Next() {
		set := rds.NewSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
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

	ec, err := client.RdsClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewEcsOperater(ec)
}
