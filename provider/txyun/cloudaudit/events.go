package cloudaudit

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
	cloudaudit "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudaudit/v20190319"
)

func (o *Operator) PageQueryEvents(req *provider.QueryEventRequest) pager.Pager {
	p := newPager(o, req)
	p.SetRate(float64(req.Rate))
	return p
}

// 用于对操作日志进行检索，便于用户进行查询相关的操作信息
// 参考: https://console.cloud.tencent.com/api/explorer?Product=cloudaudit&Version=2019-03-19&Action=DescribeEvents&SignVersion=
func (o *Operator) QueryEvents(ctx context.Context, req *cloudaudit.DescribeEventsRequest) (*lb.LoadBalancerSet, error) {
	resp, err := o.client.DescribeEventsWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	req.NextToken = resp.Response.NextToken
	fmt.Println(resp.ToJsonString())
	return nil, nil
}
