package actiontrail

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"

	actiontrail "github.com/alibabacloud-go/actiontrail-20200706/v2/client"
)

func (o *Operator) PageQueryEvent(req *provider.QueryEventRequest) pager.Pager {
	p := newPager(o, req, o.log)
	p.SetRate(float64(req.Rate))
	return p
}

// 检索详细历史事件
// 参考: https://next.api.aliyun.com/api/Actiontrail/2020-07-06/LookupEvents?params={}&sdkStyle=old&lang=GO&tab=DEMO
func (o *Operator) QueryEvents(ctx context.Context, req *actiontrail.LookupEventsRequest) (*lb.LoadBalancerSet, error) {
	resp, err := o.client.LookupEvents(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())
	return nil, nil
}
