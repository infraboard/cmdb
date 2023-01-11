package cms_test

import (
	"context"
	"testing"

	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.CmsOperator
	ctx      = context.Background()
)

// 云产品监控项: https://help.aliyun.com/document_detail/163515.html?spm=api-workbench.API%20Explorer.0.0.25386468wUNrOt
// ECS监控项: https://cms.console.aliyun.com/metric-meta/acs_ecs_dashboard/ecs?spm=a2c4g.11186623.0.0.595a76ab2VoGQ3
func TestDescribeMetric(t *testing.T) {
	req := provider.NewDescribeMetricLastRequeset("acs_rds_dashboard", "CpuUsage")
	set, err := operator.DescribeMetricLast(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func init() {
	zap.DevelopmentSetup()

	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().CmsOperator()
}
