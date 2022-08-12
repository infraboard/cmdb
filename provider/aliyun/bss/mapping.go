package bss

import (
	"github.com/infraboard/cmdb/apps/order"
)

// New：新购。
// Renew：续费。
// Upgrade：升级。
// Refund：退款。
// Convert： 付费类型转换。
// Downgrade：降配。
// ResizeDisk：磁盘扩容
var (
	ORDER_TYPE_MAP = map[string]order.ORDER_TYPE{
		"New":        order.ORDER_TYPE_NEW,
		"Renew":      order.ORDER_TYPE_RENEW,
		"Upgrade":    order.ORDER_TYPE_UPGRADE,
		"Downgrade":  order.ORDER_TYPE_DOWNGRADE,
		"Refund":     order.ORDER_TYPE_REFUND,
		"Convert":    order.ORDER_TYPE_CONVERT,
		"ResizeDisk": order.ORDER_TYPE_UPGRADE,
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

// Unpaid：未支付。
// Paid：已支付。
// Cancelled：已作废。
var (
	ORDER_STATUS_MAP = map[string]order.ORDER_STATUS{
		"Unpaid":    order.ORDER_STATUS_UNPAID,
		"Paid":      order.ORDER_STATUS_PAID,
		"Cancelled": order.ORDER_STATUS_CANCELLED,
	}
)

func praseOrderStatus(s *string) string {
	if s == nil {
		return ""
	}
	if v, ok := ORDER_STATUS_MAP[*s]; ok {
		return v.String()
	}

	return *s
}
