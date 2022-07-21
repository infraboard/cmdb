package vm

import (
	"regexp"
	"time"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/vmware/govmomi/find"
)

func NewVMOperator(client *vim25.Client) *VMOperator {
	return &VMOperator{
		masterIpRegx: "",
		client:       client,
		finder:       find.NewFinder(client, false),
		log:          zap.L().Named("Vsphere VM"),
		Timeout:      time.Second * 60,
	}
}

type VMOperator struct {
	masterIpRegx string
	client       *vim25.Client
	log          logger.Logger
	finder       *find.Finder
	Timeout      time.Duration
}

func (o *VMOperator) transferOne(ins *mo.VirtualMachine, dcName string) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.Vendor_VSPHERE
	h.Base.Region = o.client.URL().Host
	h.Base.Zone = dcName
	h.Base.CreateAt = ins.Config.CreateDate.UnixMilli()
	h.Base.Id = ins.Config.Uuid

	h.Information.Name = ins.Name
	h.Information.Status = string(ins.Summary.Runtime.PowerState)

	h.Describe.Cpu = int64(ins.Config.Hardware.NumCPU)
	h.Describe.Memory = int64(ins.Config.Hardware.MemoryMB)
	h.Describe.OsType = ins.Guest.GuestFamily
	h.Describe.OsName = ins.Guest.GuestFullName
	h.Describe.SerialNumber = ins.Config.Uuid

	// 获取主Ip
	privateIP := o.GetMasterIp(ins.Guest.Net)
	if privateIP == "" {
		privateIP = ins.Guest.IpAddress
	}
	h.Information.PrivateIp = []string{privateIP}
	return h
}

func (o *VMOperator) GetMasterIp(nics []types.GuestNicInfo) string {
	ips := []string{}
	for i := range nics {
		for j := range nics[i].IpAddress {
			ips = append(ips, nics[i].IpAddress[j])
		}
	}

	if o.masterIpRegx != "" {
		expr, _ := regexp.Compile(o.masterIpRegx)
		for _, ip := range ips {
			if expr.MatchString(ip) {
				return ip
			}
		}
	}

	if len(ips) > 0 {
		return ips[0]
	}
	return ""
}
