package cms_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.CmsOperator
	ctx      = context.Background()
)

func TestDescribeMetric(t *testing.T) {
	req := provider.NewDescribeMetricLastRequeset("acs_rds_dashboard", "CpuUsage")
	set, err := operator.DescribeMetricLast(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for i := range set.Items {
		fmt.Println(set.Items[i])
	}
}

func init() {
	zap.DevelopmentSetup()

	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().CmsOperator()
}
