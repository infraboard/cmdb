package mapping

import "github.com/infraboard/cmdb/apps/resource"

var (
	// Subscription：预付费。
	// PayAsYouGo：后付费。
	// PrePaid：包年包月。
	// PostPaid：按量付费
	// 账单:
	// 	 SubscriptionOrder (预付订单)。
	// 	 PayAsYouGoBill (后付账单)。
	// 	 Refund (退款)。
	// 	 Adjustment (调账)。
	// 订单
	// 	 Subscription：预付费。
	// 	 PayAsYouGo：按量付费。
	// MongoDB:
	//	 PrePaid：预付费，包年包月。
	// 	 PostPaid：按量付费。
	PAY_TYPE_STATUS_MAP = map[string]resource.PAY_MODE{
		"SubscriptionOrder": resource.PAY_MODE_PRE_PAY,
		"Subscription":      resource.PAY_MODE_PRE_PAY,
		"PrePaid":           resource.PAY_MODE_PRE_PAY,
		"PrePay":            resource.PAY_MODE_PRE_PAY,
		"PayAsYouGo":        resource.PAY_MODE_POST_PAY,
		"PayAsYouGoBill":    resource.PAY_MODE_POST_PAY,
		"PostPaid":          resource.PAY_MODE_POST_PAY,
		"PostPay":           resource.PAY_MODE_POST_PAY,
	}
)

func PrasePAY_MODE(s *string) resource.PAY_MODE {
	if s == nil {
		return resource.PAY_MODE_NULL
	}
	if v, ok := PAY_TYPE_STATUS_MAP[*s]; ok {
		return v
	}

	return resource.PAY_MODE_NULL
}
