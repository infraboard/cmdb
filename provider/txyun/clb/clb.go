package clb

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
)

func (o *CLBOperator) DescribeLoadBalancer(ctx context.Context, r *provider.DescribeRequest) (
	*lb.LoadBalancer, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := clb.NewDescribeLoadBalancersRequest()
	req.LoadBalancerIds = tea.StringSlice([]string{r.Id})
	req.Limit = tea.Int64(1)

	set, err := o.QueryLoadBalancer(ctx, req)
	if err != nil {
		return nil, err
	}

	if set.Length() == 0 {
		return nil, exception.NewNotFound("lb %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *CLBOperator) PageQueryLoadBalancer(req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询一个地域的负载均衡实例列表。
// 参考: https://console.cloud.tencent.com/api/explorer?Product=clb&Version=2018-03-17&Action=DescribeLoadBalancers&SignVersion=
func (o *CLBOperator) QueryLoadBalancer(ctx context.Context, req *clb.DescribeLoadBalancersRequest) (*lb.LoadBalancerSet, error) {
	resp, err := o.client.DescribeLoadBalancersWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return o.transferCLBSet(resp.Response.LoadBalancerSet), nil
}

func (o *CLBOperator) transferCLBSet(items []*clb.LoadBalancer) *lb.LoadBalancerSet {
	set := lb.NewLoadBalancerSet()
	for i := range items {
		set.Add(o.transferLB(items[i]))
	}
	return set
}

func (o *CLBOperator) transferLB(ins *clb.LoadBalancer) *lb.LoadBalancer {
	r := lb.NewDefaultLoadBalancer()

	b := r.Resource.Meta
	b.CreateAt = utils.ParseTime("2006-01-02 15:04:05", utils.PtrStrV(ins.CreateTime))

	b.Id = utils.PtrStrV(ins.LoadBalancerId)

	info := r.Resource.Spec
	info.Region = tea.StringValue(ins.TargetRegionInfo.Region)
	info.Zone = tea.StringValue(ins.AnycastZone)
	info.Vendor = resource.VENDOR_TENCENT
	info.Name = tea.StringValue(ins.LoadBalancerName)
	info.Type = tea.StringValue(ins.LoadBalancerType)
	info.ExpireAt = utils.ParseTime("2006-01-02 15:04:05", utils.PtrStrV(ins.ExpireTime))
	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(tea.StringValue(ins.ChargeType))

	r.Resource.Status.PrivateIp = tea.StringSliceValue(ins.LoadBalancerVips)
	r.Resource.Status.Phase = praseClbStatus(ins.Status)

	desc := r.Describe
	desc.Domain = tea.StringValue(ins.Domain)
	return r
}
