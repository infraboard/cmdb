package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	cmdbRds "github.com/infraboard/cmdb/pkg/rds"
)

func NewRdsOperater(client *rds.Client) *RdsOperater {
	return &RdsOperater{
		client: client,
	}
}

type RdsOperater struct {
	client *rds.Client
}

func (o *RdsOperater) transferSet(items []rds.DBInstance) *cmdbRds.RdsSet {
	set := cmdbRds.NewRdsSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *RdsOperater) transferOne(ins rds.DBInstance) *cmdbRds.Rds {
	return nil
}
