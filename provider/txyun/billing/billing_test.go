package billing_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/txyun/billing"
)

var (
	operater *op.BillingOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	req.Month = "2022-04"

	pager := operater.PageQuery(req)
	for pager.Next() {
		set := bill.NewBillSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()

	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	client := connectivity.C()
	err = client.Check()
	if err != nil {
		panic(err)
	}

	operater = op.NewBillingOperater(client.BillingClient())
}
