package ecs

import (
	"fmt"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/infraboard/cmdb/apps/disk"
)

// 仅支持查询包年包月资源的续费价格
// 参考文档: https://next.api.aliyun.com/api/Ecs/2014-05-26/DescribeRenewalPrice?params={}&tab=DEMO&lang=GO
func (o *EcsOperator) InquiryRenewPrice(req *ecs.DescribeRenewalPriceRequest) (*disk.DiskSet, error) {
	set := disk.NewDiskSet()

	resp, err := o.client.DescribeRenewalPrice(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	return set, nil
}
