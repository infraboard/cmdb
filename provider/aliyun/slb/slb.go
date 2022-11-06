package slb

import (
	"context"

	slb "github.com/alibabacloud-go/slb-20140515/v3/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

func (o *SLBOperator) DescribeLoadBalancer(ctx context.Context, r *provider.DescribeRequest) (
	*lb.LoadBalancer, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := &slb.DescribeLoadBalancersRequest{
		LoadBalancerId: &r.Id,
		RegionId:       o.client.RegionId,
		PageNumber:     tea.Int32(1),
		PageSize:       tea.Int32(1),
	}

	set, err := o.QueryLoadBalancer(req)
	if err != nil {
		return nil, err
	}
	if set.Length() == 0 {
		return nil, exception.NewNotFound("lb %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *SLBOperator) PageQueryLoadBalancer(req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询已创建的负载均衡实例
// 参考: https://next.api.aliyun.com/api/Slb/2014-05-15/DescribeLoadBalancers?params={}
func (o *SLBOperator) QueryLoadBalancer(req *slb.DescribeLoadBalancersRequest) (*lb.LoadBalancerSet, error) {
	resp, err := o.client.DescribeLoadBalancers(req)
	if err != nil {
		return nil, err
	}

	set := lb.NewLoadBalancerSet()
	set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
	set.Items = o.transferLoadBalancerSet(resp.Body.LoadBalancers).Items
	return set, nil
}

func (o *SLBOperator) transferLoadBalancerSet(items *slb.DescribeLoadBalancersResponseBodyLoadBalancers) *lb.LoadBalancerSet {
	set := lb.NewLoadBalancerSet()
	for i := range items.LoadBalancer {
		set.Add(o.transferLoadBalancer(items.LoadBalancer[i]))
	}
	return set
}

func (o *SLBOperator) transferLoadBalancer(ins *slb.DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) *lb.LoadBalancer {
	r := lb.NewDefaultLoadBalancer()
	b := r.Resource.Meta
	b.CreateAt = tea.Int64Value(ins.CreateTimeStamp) / 1000
	b.Id = tea.StringValue(ins.LoadBalancerId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_ALIYUN
	info.Region = tea.StringValue(ins.RegionId)
	info.Zone = tea.StringValue(ins.MasterZoneId)
	info.Name = tea.StringValue(ins.LoadBalancerName)
	info.Type = tea.StringValue(ins.NetworkType)

	r.Resource.Status.PrivateIp = []string{tea.StringValue(ins.Address)}
	r.Resource.Status.Phase = praseSlbStatus(ins.LoadBalancerStatus)
	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(ins.PayType)

	desc := r.Describe
	desc.BandWidth = tea.Int32Value(ins.Bandwidth)

	return r
}
