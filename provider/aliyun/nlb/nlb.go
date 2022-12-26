package nlb

import (
	"context"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"

	nlb "github.com/alibabacloud-go/nlb-20220430/client"
)

func (o *NLBOperator) DescribeLoadBalancer(ctx context.Context, r *provider.DescribeRequest) (
	*lb.LoadBalancer, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := &nlb.ListLoadBalancersRequest{
		LoadBalancerIds: []*string{&r.Id},
		MaxResults:      tea.Int32(1),
	}

	set, err := o.QueryLoadBalancer(req)
	if err != nil {
		return nil, err
	}
	if set.Length() == 0 {
		return nil, exception.NewNotFound("nlb %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *NLBOperator) PageQueryLoadBalancer(req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询指定地域的负载均衡，支持根据条件过滤
// 参考: https://next.api.aliyun.com/api/Nlb/2022-04-30/ListLoadBalancers?lang=GO
func (o *NLBOperator) QueryLoadBalancer(req *nlb.ListLoadBalancersRequest) (*lb.LoadBalancerSet, error) {
	resp, err := o.client.ListLoadBalancers(req)
	if err != nil {
		return nil, err
	}
	req.NextToken = resp.Body.NextToken

	set := lb.NewLoadBalancerSet()
	set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
	set.Items = o.transferLoadBalancerSet(resp.Body.LoadBalancers).Items
	return set, nil
}

func (o *NLBOperator) transferLoadBalancerSet(items []*nlb.ListLoadBalancersResponseBodyLoadBalancers) *lb.LoadBalancerSet {
	set := lb.NewLoadBalancerSet()
	for i := range items {
		set.Add(o.transferLoadBalancer(items[i]))
	}
	return set
}

func (o *NLBOperator) transferLoadBalancer(ins *nlb.ListLoadBalancersResponseBodyLoadBalancers) *lb.LoadBalancer {
	r := lb.NewDefaultLoadBalancer()
	b := r.Resource.Meta
	b.CreateAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.CreateTime))
	b.Id = tea.StringValue(ins.LoadBalancerId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_ALIYUN
	dnsName := tea.StringValue(ins.DNSName)
	dnsNameSlice := strings.Split(dnsName, ".")
	if len(dnsNameSlice) > 1 {
		info.Region = dnsNameSlice[1]
	}

	info.Name = tea.StringValue(ins.LoadBalancerName)
	info.Type = tea.StringValue(ins.AddressType)

	r.Resource.Status.PrivateAddress = []string{tea.StringValue(ins.DNSName)}
	r.Resource.Status.Phase = praseSlbStatus(ins.LoadBalancerStatus)
	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(ins.LoadBalancerBillingConfig.PayType)

	return r
}
