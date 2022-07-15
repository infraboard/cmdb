package disk

import "github.com/infraboard/cmdb/apps/disk"

// 查询一块或多块已经创建的块存储（包括云盘以及本地盘）
// 参考: https://next.api.aliyun.com/api/Ecs/2014-05-26/DescribeDisks?params={}
func (o *DiskOperator) Query() (*disk.Set, error) {
	return nil, nil
}
