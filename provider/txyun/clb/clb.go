package clb

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
)

func (o *CLBOperator) QueryLB(req *provider.QueryLBRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询一个地域的负载均衡实例列表。
// 参考: https://console.cloud.tencent.com/api/explorer?Product=clb&Version=2018-03-17&Action=DescribeLoadBalancers&SignVersion=
func (o *CLBOperator) queryCLB(ctx context.Context, req *clb.DescribeLoadBalancersRequest) (*lb.LBSet, error) {
	resp, err := o.client.DescribeLoadBalancersWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return o.transferCLBSet(resp.Response.LoadBalancerSet), nil
}

func (o *CLBOperator) transferCLBSet(items []*clb.LoadBalancer) *lb.LBSet {
	set := lb.NewLBSet()
	for i := range items {
		set.Add(o.transferLB(items[i]))
	}
	return set
}

func (o *CLBOperator) transferLB(ins *clb.LoadBalancer) *lb.LB {
	r := lb.NewDefaultLB()

	b := r.Base
	b.Vendor = resource.VENDOR_TENCENT
	b.CreateAt = utils.ParseTime("2006-01-02 15:04:05", utils.PtrStrV(ins.CreateTime))
	b.Region = tea.StringValue(ins.TargetRegionInfo.Region)
	b.Zone = tea.StringValue(ins.AnycastZone)
	b.Id = utils.PtrStrV(ins.LoadBalancerId)

	info := r.Information
	info.Name = tea.StringValue(ins.LoadBalancerName)
	info.Type = tea.StringValue(ins.LoadBalancerType)
	info.ExpireAt = utils.ParseTime("2006-01-02 15:04:05", utils.PtrStrV(ins.ExpireTime))
	info.PayType = tea.StringValue(ins.ChargeType)
	info.PrivateIp = tea.StringSliceValue(ins.LoadBalancerVips)
	info.Status = praseClbStatus(ins.Status)

	desc := r.Describe
	desc.Domain = tea.StringValue(ins.Domain)
	return r
}
