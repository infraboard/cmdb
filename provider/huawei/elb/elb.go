package elb

import (
	"context"
	"strings"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/model"
	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *ELBOperator) DescribeLoadBalancer(ctx context.Context, req *provider.DescribeRequest) (
	*lb.LoadBalancer, error) {
	return nil, nil
}

func (o *ELBOperator) PageQueryLoadBalancer(req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询负载均衡器
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=ELB&api=ListLoadbalancers&version=v2
func (o *ELBOperator) queryELB(req *model.ListLoadbalancersRequest) (*lb.LoadBalancerSet, error) {
	resp, err := o.client.ListLoadbalancers(req)
	if err != nil {
		return nil, err
	}

	set := lb.NewLoadBalancerSet()
	set.Items = o.transferELoadBalancerSet(resp.Loadbalancers).Items

	last := set.GetLast()
	if last != nil {
		req.Marker = &last.Base.Id
	}

	return set, nil
}

func (o *ELBOperator) transferELoadBalancerSet(list *[]model.LoadbalancerResp) *lb.LoadBalancerSet {
	set := lb.NewLoadBalancerSet()

	if list == nil {
		return set
	}

	items := *list

	for i := range items {
		set.Add(o.transferELB(items[i]))
	}
	return set
}

func (o *ELBOperator) transferELB(ins model.LoadbalancerResp) *lb.LoadBalancer {
	r := lb.NewDefaultLoadBalancer()
	b := r.Base
	b.Vendor = resource.VENDOR_HUAWEI
	b.CreateAt = utils.ParseTime("2006-01-02T15:04:05", ins.CreatedAt)
	b.Id = ins.Id

	info := r.Information
	info.Name = ins.Name
	st, _ := ins.OperatingStatus.MarshalJSON()
	info.Status = praseElbStatus(strings.Trim(strings.TrimSpace(string(st)), `"`))
	info.PrivateIp = []string{ins.VipAddress}
	return r
}
