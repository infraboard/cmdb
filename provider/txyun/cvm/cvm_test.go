package cvm_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	op "github.com/infraboard/cmdb/provider/txyun/cvm"
	"github.com/infraboard/cmdb/utils"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var (
	operater *op.CVMOperater
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

func TestInquiryPrice(t *testing.T) {
	should := assert.New(t)

	req := cvm.NewInquiryPriceRunInstancesRequest()
	req.Placement = &cvm.Placement{
		Zone: utils.StringPtr("ap-shanghai-1"),
	}
	req.ImageId = utils.StringPtr("img-l5eqiljn")
	req.InstanceType = utils.StringPtr("S1.SMALL1")
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
	var secretID, secretKey string
	if secretID = os.Getenv("TX_CLOUD_SECRET_ID"); secretID == "" {
		panic("empty TX_CLOUD_SECRET_ID")
	}

	if secretKey = os.Getenv("TX_CLOUD_SECRET_KEY"); secretKey == "" {
		panic("empty TX_CLOUD_SECRET_KEY")
	}

	client := connectivity.NewTencentCloudClient(secretID, secretKey, regions.Shanghai)
	operater = op.NewCVMOperater(client.CvmClient())
}
