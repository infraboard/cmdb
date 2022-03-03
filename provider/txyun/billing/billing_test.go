package billing_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

	op "github.com/infraboard/cmdb/provider/txyun/billing"
)

var (
	operater *op.BillingOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	req.Month = "2021-10"

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

	var secretID, secretKey string
	if secretID = os.Getenv("TX_CLOUD_SECRET_ID"); secretID == "" {
		panic("empty TX_CLOUD_SECRET_ID")
	}

	if secretKey = os.Getenv("TX_CLOUD_SECRET_KEY"); secretKey == "" {
		panic("empty TX_CLOUD_SECRET_KEY")
	}

	client := connectivity.NewTencentCloudClient(secretID, secretKey, regions.Shanghai)
	operater = op.NewBillingOperater(client.BillingClient())
}
