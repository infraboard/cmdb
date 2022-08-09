package billing

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
)

func (o *BillOperator) QueryOrder(req *provider.QueryOrderRequest) pager.Pager {
	p := newOrderPager(o, req)
	p.SetRate(req.Rate)

	return p
}

// 查询订单数据
// 参考: https://console.cloud.tencent.com/api/explorer?Product=billing&Version=2018-07-09&Action=DescribeDealsByCond&SignVersion=
func (o *BillOperator) doQueryOrder(ctx context.Context, req *billing.DescribeDealsByCondRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()
	resp, err := o.client.DescribeDealsByCondWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.ToJsonString())

	return set, nil
}
