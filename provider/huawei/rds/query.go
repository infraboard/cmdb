package rds

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"

	"github.com/infraboard/cmdb/app/rds"
)

func (o *RdsOperater) Query(req *model.ListInstancesRequest) (*rds.Set, error) {
	set := rds.NewSet()

	resp, err := o.client.ListInstances(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.TotalCount)
	set.Items = o.transferSet(resp.Instances).Items

	return set, nil
}

func (o *RdsOperater) PageQuery() rds.Pager {
	return newPager(20, o)
}
