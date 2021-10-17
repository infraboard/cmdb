package cvm

import (
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/pkg/host"
)

func NewCVMOperater(client *cvm.Client) *CVMOperater {
	return &CVMOperater{
		client: client,
	}
}

type CVMOperater struct {
	client *cvm.Client
}

func (o *CVMOperater) transferSet(items []*cvm.Instance) *host.HostSet {
	set := host.NewHostSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *CVMOperater) transferOne(ins *cvm.Instance) *host.Host {
	return nil
}
