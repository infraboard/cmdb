package ecs

import (
	"strconv"
	"time"

	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperater(client *ecs.EcsClient) *EcsOperater {
	return &EcsOperater{
		client: client,
		log:    zap.L().Named("Huawei ECS"),
	}
}

type EcsOperater struct {
	client *ecs.EcsClient
	log    logger.Logger
}

func (o *EcsOperater) transferSet(list *[]model.ServerDetail) *host.HostSet {
	set := host.NewHostSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *EcsOperater) transferOne(ins model.ServerDetail) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.Vendor_HUAWEI
	h.Base.Zone = ins.OSEXTAZavailabilityZone
	h.Base.CreateAt = o.parseTime(ins.Created)
	h.Base.InstanceId = ins.Id

	h.Information.Category = ins.Flavor.Name
	h.Information.ExpireAt = o.parseTime(ins.AutoTerminateTime)
	h.Information.Name = ins.Name
	h.Information.Description = utils.PtrStrV(ins.Description)
	h.Information.Status = ins.Status
	h.Information.Tags = o.transferTags(ins.Tags)
	h.Information.PrivateIp, h.Information.PublicIp = o.parseIp(ins.Addresses)
	h.Information.PayType = ins.Metadata["charging_mode"]

	h.Describe.SerialNumber = ins.Id
	h.Describe.Cpu, _ = strconv.ParseInt(ins.Flavor.Vcpus, 10, 64)
	h.Describe.Memory, _ = strconv.ParseInt(ins.Flavor.Ram, 10, 64)
	h.Describe.OsType = ins.Metadata["os_type"]
	h.Describe.OsName = ins.Metadata["image_name"]
	h.Describe.ImageId = ins.Image.Id
	h.Describe.KeyPairName = []string{ins.KeyName}
	return h
}

func (o *EcsOperater) transferTags(tags *[]string) map[string]string {
	return nil
}

func (o *EcsOperater) parseTime(t string) int64 {
	if t == "" {
		return 0
	}

	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}

func (o *EcsOperater) parseIp(address map[string][]model.ServerAddress) (privateIps []string, publicIps []string) {
	for _, addrs := range address {
		for i := range addrs {
			switch *addrs[i].OSEXTIPStype {
			case model.GetServerAddressOSEXTIPStypeEnum().FIXED:
				privateIps = append(privateIps, addrs[i].Addr)
			case model.GetServerAddressOSEXTIPStypeEnum().FLOATING:
				publicIps = append(publicIps, addrs[i].Addr)
			}
		}
	}
	return
}
