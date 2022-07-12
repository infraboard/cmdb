package cms

import (
	"context"
	"fmt"

	cms "github.com/alibabacloud-go/cms-20190101/v7/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/provider"
)

// 查询指定监控项的最新监控数据
// 参考: https://next.api.aliyun.com/api/Cms/2019-01-01/DescribeMetricLast
func (o *CmsOperator) DescribeMetricLast(ctx context.Context, req *provider.DescribeMetricLastRequeset) (*provider.DatapointSet, error) {
	describeMetricLastRequest := &cms.DescribeMetricLastRequest{
		Length:     tea.String("2000"),
		Namespace:  tea.String("acs_rds_dashboard"),
		MetricName: tea.String("CpuUsage"),
	}

	resp, err := o.client.DescribeMetricLast(describeMetricLastRequest)

	if !*resp.Body.Success {
		o.log.Debug(resp.String())
		return nil, fmt.Errorf(*resp.Body.Message)
	}
	return nil, err
}
