package bss

import (
	"fmt"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/infraboard/cmdb/apps/bill"
)

// 查询用户或者分销客户订单列表情况。默认查询当前时间最近1小时范围内订单，如需查询更长时间范围的订单数据，
// 请设**CreateTimeStart** 和**CreateTimeEnd**参数
// 参考文档: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/QueryOrders?params={}
func (o *BssOperator) doQueryOrder(req *bssopenapi.QueryOrdersRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()
	resp, err := o.client.QueryOrders(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	data := resp.Body.Data
	set.Total = int64(*data.TotalCount)
	return set, nil
}
