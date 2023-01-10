package ec2_test

import (
	"context"
	"testing"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aws"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.HostOperator
	ctx      = context.Background()
)

func TestQuery(t *testing.T) {
	pager := operator.PageQueryHost(provider.NewQueryRequest())

	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(ctx, set); err != nil {
			panic(err)
		}
		t.Log(set)
	}
}

func init() {
	zap.DevelopmentSetup()
	err := aws.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aws.O().HostOperator()
}
