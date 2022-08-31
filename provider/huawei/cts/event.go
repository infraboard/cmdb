package cts

import (
	"context"
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cts/v3/model"
	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *Operator) PageQueryEvent(req *provider.QueryEventRequest) pager.Pager {
	p := newPager(o, req, o.log)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询事件列表
// 参考: https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=CTS&api=ListTraces
func (o *Operator) QueryEvents(ctx context.Context, req *model.ListTracesRequest) (*lb.LoadBalancerSet, error) {
	resp, err := o.client.ListTraces(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())
	return nil, nil
}
