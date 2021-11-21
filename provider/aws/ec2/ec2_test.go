package ec2_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/aws/connectivity"
	op "github.com/infraboard/cmdb/provider/aws/ec2"
)

var (
	operater *op.Ec2Operater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	pager := operater.PageQuery(req)

	hasNext := true
	for hasNext {
		p := pager.Next()
		hasNext = p.HasNext
		fmt.Println(p.Data)
	}
}

func init() {
	var secretID, secretKey string
	if secretID = os.Getenv("AMAZON_CLOUD_ACCESS_KEY"); secretID == "" {
		panic("empty AMAZON_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("AMAZON_CLOUD_ACCESS_SECRET"); secretKey == "" {
		panic("empty AMAZON_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAwsCloudClient(secretID, secretKey, "ap-south-1")

	ec, err := client.Ec2Client()
	if err != nil {
		panic(err)
	}
	operater = op.NewEc2Operator(ec)
}
