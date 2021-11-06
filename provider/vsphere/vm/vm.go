package vm

import (
	"time"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/vmware/govmomi/find"
)

func NewVmOperater(client *vim25.Client) *VMOperater {
	return &VMOperater{
		client:  client,
		finder:  find.NewFinder(client, false),
		log:     zap.L().Named("Vsphere VM"),
		Timeout: time.Second * 60,
	}
}

type VMOperater struct {
	client  *vim25.Client
	log     logger.Logger
	finder  *find.Finder
	Timeout time.Duration
}

func (o *VMOperater) transferOne(ins *mo.VirtualMachine, dcName string) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.Vendor_VSPHERE
	h.Base.Region = o.client.URL().Host
	h.Base.Zone = dcName
	h.Base.CreateAt = ins.Config.CreateDate.UnixNano() / 1000000
	h.Base.InstanceId = ins.Config.Uuid

	h.Information.Name = ins.Name
	h.Information.Status = string(ins.Summary.Runtime.PowerState)
	h.Information.PrivateIp = []string{ins.Guest.IpAddress}

	h.Describe.CPU = int64(ins.Config.Hardware.NumCPU)
	h.Describe.Memory = int64(ins.Config.Hardware.MemoryMB)
	h.Describe.OSType = ins.Guest.GuestFamily
	h.Describe.OSName = ins.Guest.GuestFullName
	h.Describe.SerialNumber = ins.Config.Uuid
	return h
}
