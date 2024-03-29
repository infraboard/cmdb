package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/cmdb/provider"
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
	operator provider.BillOperator
)

func TestSyncBill(t *testing.T) {
	should := assert.New(t)

	req := provider.NewQueryBillRequest()
	req.Month = "2022-05"

	pager := operator.PageQueryBill(req)

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
	operator = op.NewBillingoperator(client.BillingClient())
}
