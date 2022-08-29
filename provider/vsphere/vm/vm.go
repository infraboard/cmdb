package vm

import (
	"regexp"
	"strings"
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
		log:          zap.L().Named("vsphere.vm"),
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
	b := h.Resource.Base
	b.Vendor = resource.VENDOR_VSPHERE
	b.Region = o.client.URL().Host
	b.Zone = dcName
	b.CreateAt = ins.Config.CreateDate.UnixMilli()
	b.Id = ins.Config.Uuid

	i := h.Resource.Information
	i.Name = ins.Name
	i.Status = praseStatus(string(ins.Summary.Runtime.PowerState))
	i.Cpu = ins.Config.Hardware.NumCPU
	i.Memory = ins.Config.Hardware.MemoryMB
	i.SerialNumber = ins.Config.Uuid

	h.Describe.OsType = strings.TrimSuffix(ins.Guest.GuestFamily, "Guest")
	h.Describe.OsName = ins.Guest.GuestFullName

	// 获取主Ip
	privateIP := o.GetMasterIp(ins.Guest.Net)
	if privateIP == "" {
		privateIP = ins.Guest.IpAddress
	}
	h.Resource.Information.PrivateIp = []string{privateIP}
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

func ParseExtraConfigValue(v string) map[string]string {
	conf := map[string]string{}

	key := []rune{}
	value := []rune{}
	count := 0
	parsekey := true
	for _, c := range v {
		if parsekey {
			// 解析key
			if c != '=' {
				key = append(key, c)
			} else {
				parsekey = false
			}
		} else {
			// 解析value
			if c == '\'' {
				count++
				if count%2 == 0 {
					conf[strings.TrimSpace(string(key))] = strings.TrimSpace(string(value))
					key = []rune{}
					value = []rune{}
					parsekey = true
				}
			} else {
				value = append(value, c)
			}
		}
	}

	return conf
}
