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
	PAY_TYPE_STATUS_MAP = map[string]resource.PayMode{
		"SubscriptionOrder": resource.PayMode_PRE_PAY,
		"Subscription":      resource.PayMode_PRE_PAY,
		"PrePaid":           resource.PayMode_PRE_PAY,
		"PrePay":            resource.PayMode_PRE_PAY,
		"PayAsYouGo":        resource.PayMode_POST_PAY,
		"PayAsYouGoBill":    resource.PayMode_POST_PAY,
		"PostPaid":          resource.PayMode_POST_PAY,
		"PostPay":           resource.PayMode_POST_PAY,
	}
)

func PrasePayMode(s *string) resource.PayMode {
	if s == nil {
		return resource.PayMode_NULL
	}
	if v, ok := PAY_TYPE_STATUS_MAP[*s]; ok {
		return v
	}

	return resource.PayMode_NULL
}
