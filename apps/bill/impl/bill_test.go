package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"

	"github.com/infraboard/cmdb/apps/bill"
	op "github.com/infraboard/cmdb/provider/txyun/billing"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"

	_ "github.com/infraboard/cmdb/apps/bill/impl"
)

var (
	svc      bill.ServiceServer
	operater *op.BillingOperater
)

func TestSyncBill(t *testing.T) {
	should := assert.New(t)

	req := op.NewPageQueryRequest()
	req.Month = "2022-04"

	pager := operater.PageQuery(req)

	ctx := context.Background()
	for pager.Next() {
		set := bill.NewBillSet()

		err := pager.Scan(ctx, set)
		if should.NoError(err) {
			for _, item := range set.Items {
				_, err = svc.SyncBill(ctx, item)
				if err != nil {
					t.Fatal(err)
					return
				}
			}
		}
	}
}

func init() {
	zap.DevelopmentSetup()
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
	svc = app.GetGrpcApp(bill.AppName).(bill.ServiceServer)

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
