package cvm_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.HostOperator
)

func TestQuery(t *testing.T) {
	pager := operator.QueryHost(provider.NewQueryHostRequest())

	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

// func TestInquiryPrice(t *testing.T) {
// 	should := assert.New(t)

// 	req := cvm.NewInquiryPriceRunInstancesRequest()
// 	req.Placement = &cvm.Placement{
// 		Zone: utils.StringPtr("ap-shanghai-2"),
// 	}
// 	req.ImageId = utils.StringPtr("img-l5eqiljn")
// 	req.InstanceType = utils.StringPtr("S4.SMALL1")
// 	req.InstanceChargeType = utils.StringPtr("SPOTPAID")
// 	err := operator.InquiryPrice(req)
// 	should.NoError(err)
// }

// func TestDescribeZones(t *testing.T) {
// 	operator.DescribeZones()
// }

// func TestDescribeInstanceType(t *testing.T) {
// 	operator.DescribeInstanceType()
// }

func TestCreate(t *testing.T) {
}

func init() {
	zap.DevelopmentSetup()

	err := txyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}

	operator = txyun.O().HostOperator()
}
