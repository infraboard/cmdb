package ec2_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aws/connectivity"
	op "github.com/infraboard/cmdb/provider/aws/ec2"
)

var (
	operator *op.Ec2operator
)

func TestQuery(t *testing.T) {
	pager := operator.PageQueryHost(provider.NewQueryRequest())

	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func init() {
	var credentialID, credentialKey string
	if credentialID = os.Getenv("AMAZON_CLOUD_ACCESS_KEY"); credentialID == "" {
		panic("empty AMAZON_CLOUD_ACCESS_KEY")
	}

	if credentialKey = os.Getenv("AMAZON_CLOUD_ACCESS_SECRET"); credentialKey == "" {
		panic("empty AMAZON_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAwsCloudClient(credentialID, credentialKey, "ap-south-1")

	ec, err := client.Ec2Client()
	if err != nil {
		panic(err)
	}
	operator = op.NewEc2Operator(ec)
}
