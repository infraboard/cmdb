package cvm_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/txyun/cvm"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	operator *op.CVMOperator
)

func TestQueryCVM(t *testing.T) {
	pager := operator.QueryHost(provider.NewQueryHostRequest())

	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestQueryDisk(t *testing.T) {
	pager := operator.QueryDisk(provider.NewQueryDiskRequest())

	for pager.Next() {
		set := disk.NewDiskSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeEcs(t *testing.T) {
	req := provider.NewDescribeHostRequest("ins-1wxeveeb")
	ins, err := operator.DescribeHost(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestInquiryPrice(t *testing.T) {
	req := cvm.NewInquiryPriceRunInstancesRequest()
	req.Placement = &cvm.Placement{
		Zone: utils.StringPtr("ap-shanghai-2"),
	}
	req.ImageId = utils.StringPtr("img-l5eqiljn")
	req.InstanceType = utils.StringPtr("S4.SMALL1")
	req.InstanceChargeType = utils.StringPtr("SPOTPAID")
	if err := operator.InquiryPrice(req); err != nil {

	}
}

func TestDescribeZones(t *testing.T) {
	operator.DescribeZones()
}

func TestDescribeInstanceType(t *testing.T) {
	operator.DescribeInstanceType()
}

func TestCreate(t *testing.T) {
}

func init() {
	zap.DevelopmentSetup()

	err := txyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}

	c := txyun.O().Client()
	operator = op.NewCVMOperator(c.CvmClient(), c.CBSClient())
}
