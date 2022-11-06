package mapping

import "github.com/infraboard/cmdb/apps/resource"

var (
	// CDB付费类型，可能的返回值：0-包年包月；1-按量计费
	// 负载均衡实例的计费类型，PREPAID：包年包月，POSTPAID_BY_HOUR：按量计费
	// 实例计费模式。取值范围：<br><li>`PREPAID`：表示预付费，即包年包月<br><li>`POSTPAID_BY_HOUR`：表示后付费，即按量计费<br><li>`CDHPAID`：`CDH`付费，即只对`CDH`计费，不对`CDH`上的实例计费。<br><li>`SPOTPAID`：表示竞价实例付费。
	// 磁盘付费模式。取值范围：<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：后付费，即按量计费。
	// 弹性公网IP的网络计费模式。注意，传统账户类型账户的弹性公网IP没有网络计费模式属性，值为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 包括：
	// <li><strong>BANDWIDTH_PREPAID_BY_MONTH</strong></li>
	// <p style="padding-left: 30px;">表示包月带宽预付费。</p>
	// <li><strong>TRAFFIC_POSTPAID_BY_HOUR</strong></li>
	// <p style="padding-left: 30px;">表示按小时流量后付费。</p>
	// <li><strong>BANDWIDTH_POSTPAID_BY_HOUR</strong></li>
	// <p style="padding-left: 30px;">表示按小时带宽后付费。</p>
	// <li><strong>BANDWIDTH_PACKAGE</strong></li>
	// <p style="padding-left: 30px;">表示共享带宽包。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。

	// 订单付费模式：prePay 预付费 postPay后付费 riPay预留实例
	PAY_TYPE_STATUS_MAP = map[string]resource.PAY_MODE{
		"包年包月":             resource.PAY_MODE_PRE_PAY,
		"0":                resource.PAY_MODE_PRE_PAY,
		"PREPAID":          resource.PAY_MODE_PRE_PAY,
		"prePay":           resource.PAY_MODE_PRE_PAY,
		"1":                resource.PAY_MODE_POST_PAY,
		"POSTPAID_BY_HOUR": resource.PAY_MODE_POST_PAY,
		"postPay":          resource.PAY_MODE_POST_PAY,
		"SPOTPAID":         resource.PAY_MODE_POST_PAY,
		"按量计费":             resource.PAY_MODE_POST_PAY,
		"riPay":            resource.PAY_MODE_RESERVED_PAY,
	}
)

func PrasePAY_MODE(s string) resource.PAY_MODE {
	if v, ok := PAY_TYPE_STATUS_MAP[s]; ok {
		return v
	}

	return resource.PAY_MODE_NULL
}
