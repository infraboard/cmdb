package slb

import (
	slb "github.com/alibabacloud-go/slb-20140515/v3/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *SLBOperator) QueryLB(req *provider.QueryLBRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询已创建的负载均衡实例
// 参考: https://next.api.aliyun.com/api/Slb/2014-05-15/DescribeLoadBalancers?params={}
func (o *SLBOperator) Query(req *slb.DescribeLoadBalancersRequest) (*lb.LBSet, error) {
	resp, err := o.client.DescribeLoadBalancers(req)
	if err != nil {
		return nil, err
	}
	set := lb.NewLBSet()
	set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
	set.Items = o.transferLBSet(resp.Body.LoadBalancers).Items
	return nil, nil
}

func (o *SLBOperator) transferLBSet(items *slb.DescribeLoadBalancersResponseBodyLoadBalancers) *lb.LBSet {
	set := lb.NewLBSet()
	for i := range items.LoadBalancer {
		set.Add(o.transferLB(items.LoadBalancer[i]))
	}
	return set
}

func (o *SLBOperator) transferLB(ins *slb.DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) *lb.LB {
	r := lb.NewDefaultLB()
	b := r.Base
	b.Vendor = resource.VENDOR_ALIYUN
	b.Region = tea.StringValue(ins.RegionId)
	b.Zone = tea.StringValue(ins.MasterZoneId)
	b.CreateAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.CreateTime))
	b.Id = tea.StringValue(ins.LoadBalancerId)

	info := r.Information
	info.Name = tea.StringValue(ins.LoadBalancerName)
	info.Type = tea.StringValue(ins.NetworkType)
	info.Status = tea.StringValue(ins.LoadBalancerStatus)
	info.PayType = tea.StringValue(ins.PayType)
	return r
}
