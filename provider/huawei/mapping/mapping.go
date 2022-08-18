package mapping

import "github.com/infraboard/cmdb/apps/resource"

var (
	// 计费模式，0表示按需计费，1表示包年/包月计费。
	PAY_TYPE_STATUS_MAP = map[int32]resource.PayMode{
		0: resource.PayMode_POST_PAY,
		1: resource.PayMode_PRE_PAY,
	}
)

func PrasePayMode(s *int32) resource.PayMode {
	if s == nil {
		return resource.PayMode_NULL
	}
	if v, ok := PAY_TYPE_STATUS_MAP[*s]; ok {
		return v
	}

	return resource.PayMode_NULL
}
