package bss

import (
	"fmt"

	"github.com/infraboard/cmdb/apps/order"
	"github.com/infraboard/cmdb/apps/resource"
)

// 1:开通
// 2:续订
// 3:变更
// 4:退订
// 11:按需转包年/包月
// 13:试用
// 14:转商用
// 15:费用调整
var (
	ORDER_TYPE_MAP = map[int32]order.ORDER_TYPE{
		1:  order.ORDER_TYPE_NEW,
		2:  order.ORDER_TYPE_RENEW,
		3:  order.ORDER_TYPE_MODIFY,
		4:  order.ORDER_TYPE_REFUND,
		11: order.ORDER_TYPE_MODIFY,
		13: order.ORDER_TYPE_NEW,
		14: order.ORDER_TYPE_NEW,
		15: order.ORDER_TYPE_NEW,
	}
)

func praseOrderType(s *int32) string {
	if s == nil {
		return ""
	}
	if v, ok := ORDER_TYPE_MAP[*s]; ok {
		return v.String()
	}

	return fmt.Sprintf("%d", s)
}

//  1:待审核
//  3:处理中
//  4:已取消
//  5:已完成
//  6:待支付
//  9:待确认
var (
	ORDER_STATUS_MAP = map[int32]order.ORDER_STATUS{
		1: order.ORDER_STATUS_UNPAID,
		6: order.ORDER_STATUS_UNPAID,
		9: order.ORDER_STATUS_PAYING,
		3: order.ORDER_STATUS_PAYING,
		4: order.ORDER_STATUS_CANCELLED,
		5: order.ORDER_STATUS_PAID,
	}
)

func praseOrderStatus(s *int32) string {
	if s == nil {
		return ""
	}
	if v, ok := ORDER_STATUS_MAP[*s]; ok {
		return v.String()
	}

	return fmt.Sprintf("%d", s)
}

// hws.service.type.rds
// hws.service.type.ec2
// hws.service.type.dcs
// hws.service.type.cce
// hws.service.type.sfs
// hws.service.type.dds
// hws.service.type.cfw
// hws.service.type.ebs
// hws.service.type.natgateway
// hws.service.type.ec2,hws.service.type.vpc
// hws.service.type.cbr
// hws.service.type.ec2,hws.service.type.ebs
// hws.service.type.vpc
// hws.service.type.dcaas
// hws.service.type.vpn
// hws.service.type.taurus
var (
	RESOURCE_TYPE_MAP = map[string]resource.TYPE{
		"hws.service.type.ec2": resource.TYPE_HOST,
		"hws.service.type.rds": resource.TYPE_RDS,
		"hws.service.type.dcs": resource.TYPE_REDIS,
		"hws.service.type.dds": resource.TYPE_MONGODB,
		"hws.service.type.ebs": resource.TYPE_DISK,
		"hws.service.type.cce": resource.TYPE_HOST,
		"hws.service.type.vpc": resource.TYPE_EIP,
	}
)

func praseResourceType(s *string) resource.TYPE {
	if s == nil {
		return resource.TYPE_OTHER
	}
	if v, ok := RESOURCE_TYPE_MAP[*s]; ok {
		return v
	}

	return resource.TYPE_OTHER
}
