package cvm_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	op "github.com/infraboard/cmdb/provider/txyun/cvm"
)

var (
	operater *op.CVMOperater
)

func TestQuery(t *testing.T) {
	pager := operater.Query()

	hasNext := true
	for hasNext {
		p := pager.Next()
		hasNext = p.HasNext
		fmt.Println(p.Data)
	}
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
