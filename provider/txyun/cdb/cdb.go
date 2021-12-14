package cdb

import (
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/infraboard/cmdb/app/rds"
)

func NewCDBOperater(client *cdb.Client) *CDBOperater {
	return &CDBOperater{
		client: client,
	}
}

type CDBOperater struct {
	client *cdb.Client
}

func (o *CDBOperater) transferSet(items []*cdb.InstanceInfo) *rds.Set {
	set := rds.NewSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *CDBOperater) transferOne(ins *cdb.InstanceInfo) *rds.RDS {
	return nil
}
