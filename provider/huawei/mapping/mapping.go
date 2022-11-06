package mapping

import "github.com/infraboard/cmdb/apps/resource"

var (
	// 计费模式，0表示按需计费，1表示包年/包月计费。
	PAY_TYPE_STATUS_MAP = map[int32]resource.PAY_MODE{
		0: resource.PAY_MODE_POST_PAY,
		1: resource.PAY_MODE_PRE_PAY,
	}
)

func PrasePAY_MODE(s *int32) resource.PAY_MODE {
	if s == nil {
		return resource.PAY_MODE_NULL
	}
	if v, ok := PAY_TYPE_STATUS_MAP[*s]; ok {
		return v
	}

	return resource.PAY_MODE_NULL
}
