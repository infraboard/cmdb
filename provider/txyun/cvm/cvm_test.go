package cvm_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	op "github.com/infraboard/cmdb/provider/txyun/cvm"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger/zap"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	operater *op.CVMOperater
)

func TestQuery(t *testing.T) {
	pager := operater.PageQuery(op.NewPageQueryRequest(5))

	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestInquiryPrice(t *testing.T) {
	should := assert.New(t)

	req := cvm.NewInquiryPriceRunInstancesRequest()
	req.Placement = &cvm.Placement{
		Zone: utils.StringPtr("ap-shanghai-2"),
	}
	req.ImageId = utils.StringPtr("img-l5eqiljn")
	req.InstanceType = utils.StringPtr("S4.SMALL1")
	req.InstanceChargeType = utils.StringPtr("SPOTPAID")
	err := operater.InquiryPrice(req)
	should.NoError(err)

}

func TestDescribeZones(t *testing.T) {
	operater.DescribeZones()
}

func TestDescribeInstanceType(t *testing.T) {
	operater.DescribeInstanceType()
}

func TestCreate(t *testing.T) {
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
	err := client.Check()
	if err != nil {
		panic(err)
	}
	operater = op.NewCVMOperater(client.CvmClient())
	operater.WithAccountId(client.AccountID())
}
