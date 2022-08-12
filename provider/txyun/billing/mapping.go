package billing

import (
	"fmt"

	"github.com/infraboard/cmdb/apps/order"
)

// modifyNetworkMode 调整带宽模式
// modifyNetworkSize 调整带宽大小
// refund 退款
// downgrade 降配
// upgrade 升配
// renew 续费
// purchase 购买
// preMoveOut 包年包月迁出资源
// preMoveIn 包年包月迁入资源
// preToPost 预付费转后付费
// postMoveOut 按量计费迁出资源
// postMoveIn 按量计费迁入资源
var (
	ORDER_TYPE_MAP = map[string]order.ORDER_TYPE{
		"modifyNetworkMode": order.ORDER_TYPE_CONVERT,
		"modifyNetworkSize": order.ORDER_TYPE_CONVERT,
		"refund":            order.ORDER_TYPE_REFUND,
		"downgrade":         order.ORDER_TYPE_DOWNGRADE,
		"upgrade":           order.ORDER_TYPE_UPGRADE,
		"renew":             order.ORDER_TYPE_RENEW,
		"purchase":          order.ORDER_TYPE_NEW,
		"preMoveOut":        order.ORDER_TYPE_CONVERT,
		"preMoveIn":         order.ORDER_TYPE_CONVERT,
		"preToPost":         order.ORDER_TYPE_CONVERT,
		"postMoveOut":       order.ORDER_TYPE_CONVERT,
		"postMoveIn":        order.ORDER_TYPE_CONVERT,
	}
)

func praseOrderType(s *string) string {
	if s == nil {
		return ""
	}
	if v, ok := ORDER_TYPE_MAP[*s]; ok {
		return v.String()
	}

	return *s
}

// 1：未支付
// 2：已支付
// 3：发货中
// 4：已发货
// 5：发货失败
// 6：已退款
// 7：已关单
// 8：订单过期
// 9：订单已失效
// 10：产品已失效
// 11：代付拒绝
// 12：支付中
var (
	ORDER_STATUS_MAP = map[int64]order.ORDER_STATUS{
		1:  order.ORDER_STATUS_UNPAID,
		2:  order.ORDER_STATUS_PAID,
		3:  order.ORDER_STATUS_PAID,
		4:  order.ORDER_STATUS_PAID,
		5:  order.ORDER_STATUS_PAID,
		6:  order.ORDER_STATUS_PAID,
		7:  order.ORDER_STATUS_PAID,
		8:  order.ORDER_STATUS_EXPIRED,
		9:  order.ORDER_STATUS_EXPIRED,
		10: order.ORDER_STATUS_EXPIRED,
		11: order.ORDER_STATUS_UNPAID,
		12: order.ORDER_STATUS_PAYING,
	}
)

func praseOrderStatus(s *int64) string {
	if s == nil {
		return ""
	}
	if v, ok := ORDER_STATUS_MAP[*s]; ok {
		return v.String()
	}

	return fmt.Sprintf("%d", s)
}
