package ecs

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

func (o *EcsOperator) DescribeHost(ctx context.Context, req *provider.DescribeHostRequest) (*host.Host, error) {
	resp, err := o.client.ShowServer(&model.ShowServerRequest{
		ServerId: req.Id,
	})
	if err != nil {
		if strings.Contains(err.Error(), "could not be found.") {
			return nil, exception.NewNotFound(err.Error())
		}
		return nil, err
	}

	if resp.Server.Status == "DELETED" {
		return nil, exception.NewNotFound("%s not found", req.Id)
	}

	if resp.Server.Id != req.Id {
		return nil, exception.NewNotFound("%s not found", req.Id)
	}

	h := o.transferInstance(*resp.Server)
	return h, nil
}

func (o *EcsOperator) PageQueryHost(req *provider.QueryHostRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询云服务器详情列表
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=ECS&api=ListServersDetails
func (o *EcsOperator) queryInstance(req *model.ListServersDetailsRequest) (*host.HostSet, error) {
	set := host.NewHostSet()

	resp, err := o.client.ListServersDetails(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.Count)
	set.Items = o.transferInstanceSet(resp.Servers).Items

	return set, nil
}

func (o *EcsOperator) transferInstanceSet(list *[]model.ServerDetail) *host.HostSet {
	set := host.NewHostSet()
	items := *list
	for i := range items {
		set.Add(o.transferInstance(items[i]))
	}
	return set
}

func (o *EcsOperator) transferInstance(ins model.ServerDetail) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.VENDOR_HUAWEI
	h.Base.Zone = ins.OSEXTAZavailabilityZone
	h.Base.CreateAt = o.parseTime(ins.Created)
	h.Base.Id = ins.Id

	h.Information.Category = ins.Flavor.Name
	h.Information.ExpireAt = o.parseTime(ins.AutoTerminateTime)
	h.Information.Name = ins.Name
	h.Information.Description = utils.PtrStrV(ins.Description)
	h.Information.Status = praseEcsStatus(ins.Status)
	if ins.Tags != nil {
		h.Information.Tags = o.transferTags(*ins.Tags)
	}
	h.Information.PrivateIp, h.Information.PublicIp = o.parseIp(ins.Addresses)
	h.Information.PayMode = o.ParseChangeMode(ins.Metadata["charging_mode"])
	h.Information.Owner = o.GetAccountId()

	h.Describe.SerialNumber = ins.Id
	h.Describe.Cpu, _ = strconv.ParseInt(ins.Flavor.Vcpus, 10, 64)
	h.Describe.Memory, _ = strconv.ParseInt(ins.Flavor.Ram, 10, 64)
	h.Describe.OsType = ins.Metadata["os_type"]
	h.Describe.OsName = ins.Metadata["image_name"]
	h.Describe.ImageId = ins.Image.Id
	h.Describe.KeyPairName = []string{ins.KeyName}
	return h
}

func (o *EcsOperator) parseTime(t string) int64 {
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

//  1. charging_mode 云服务器的计费类型。  - “0”：按需计费（即postPaid-后付费方式）。 - “1”：按包年包月计费（即prePaid-预付费方式）。\"2\"：
func (o *EcsOperator) ParseChangeMode(mode string) resource.PayMode {
	switch mode {
	case "0":
		return resource.PayMode_POST_PAY
	case "1":
		return resource.PayMode_PRE_PAY
	case "2":
		return resource.PayMode_POST_PAY
	}

	return resource.PayMode_NULL
}

func (o *EcsOperator) parseIp(address map[string][]model.ServerAddress) (privateIps []string, publicIps []string) {
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

func (o *EcsOperator) transferTags(tags []string) (ret []*resource.Tag) {
	for _, t := range tags {
		kv := strings.Split(t, "=")
		if len(kv) == 2 {
			ret = append(ret, resource.NewThirdTag(kv[0], kv[1]))
		} else {
			ret = append(ret, resource.NewThirdTag("ecs", t))
		}
	}

	return
}
