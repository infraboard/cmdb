package billing

import (
	"context"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newOrderPager(operator *BillOperator, r *provider.QueryOrderRequest) pager.Pager {
	req := billing.NewDescribeDealsByCondRequest()
	req.StartTime = tea.String(r.StartTime.Format(utils.TIME_SECOND_FORMAT_MOD1))
	req.EndTime = tea.String(r.EndTime.Format(utils.TIME_SECOND_FORMAT_MOD1))
	fmt.Println(req.ToJsonString())
	return &orderPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("tx.order"),
	}
}

type orderPager struct {
	*pager.BasePager
	operator *BillOperator
	req      *billing.DescribeDealsByCondRequest
	log      logger.Logger
}

func (p *orderPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.doQueryOrder(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d order", resp.Length())

	// 由于账单接口并没有返回Total总数量, 无法通过梳理判断是否数据已经拉起完成, 改而判断是否有数据
	p.CheckHasNext(set)
	return nil
}

func (p *orderPager) nextReq() *billing.DescribeDealsByCondRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Int64Ptr(p.Offset())
	p.req.Limit = common.Int64Ptr(p.PageSize())
	return p.req
}
