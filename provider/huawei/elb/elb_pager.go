package elb

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/model"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/utils"
)

func newPager(operator *ELBOperator) pager.Pager {
	req := &model.ListLoadbalancersRequest{}

	return &ecsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("huawei.elb"),
	}
}

type ecsPager struct {
	*pager.BasePager
	operator *ELBOperator
	req      *model.ListLoadbalancersRequest
	log      logger.Logger
}

func (p *ecsPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryLoadBalancer(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *ecsPager) nextReq() *model.ListLoadbalancersRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())

	// 注意: 华为云的Offse表示的是页码
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	return p.req
}
