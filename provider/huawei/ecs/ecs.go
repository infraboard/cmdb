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

func (o *EcsOperator) DescribeHost(ctx context.Context, req *provider.DescribeRequest) (*host.Host, error) {
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

func (o *EcsOperator) PageQueryHost(req *provider.QueryRequest) pager.Pager {
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
	r := host.NewDefaultHost()
	b := r.Resource.Meta
	b.CreateAt = o.parseTime(ins.Created)
	b.Id = ins.Id
	b.SerialNumber = ins.Id

	i := r.Resource.Spec
	i.Vendor = resource.VENDOR_HUAWEI
	i.Zone = ins.OSEXTAZavailabilityZone
	i.Owner = o.GetAccountId()
	i.Category = ins.Flavor.Name
	i.ExpireAt = o.parseTime(ins.AutoTerminateTime)
	i.Name = ins.Name
	i.Description = utils.PtrStrV(ins.Description)
	r.Resource.Cost.PayMode = o.ParseChangeMode(ins.Metadata["charging_mode"])

	r.Resource.Status.PrivateIp, r.Resource.Status.PublicIp = o.parseIp(ins.Addresses)
	r.Resource.Status.Phase = praseEcsStatus(ins.Status)

	cpu, _ := strconv.ParseInt(ins.Flavor.Vcpus, 10, 64)
	i.Cpu = int32(cpu)
	mem, _ := strconv.ParseInt(ins.Flavor.Ram, 10, 64)
	i.Memory = int32(mem)

	if ins.Tags != nil {
		r.Resource.Tags = o.transferTags(*ins.Tags)
	}

	r.Describe.OsType = ins.Metadata["os_type"]
	r.Describe.OsName = ins.Metadata["image_name"]
	r.Describe.ImageId = ins.Image.Id
	r.Describe.KeyPairName = []string{ins.KeyName}
	return r
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

// 1. charging_mode 云服务器的计费类型。  - “0”：按需计费（即postPaid-后付费方式）。 - “1”：按包年包月计费（即prePaid-预付费方式）。\"2\"：
func (o *EcsOperator) ParseChangeMode(mode string) resource.PAY_MODE {
	switch mode {
	case "0":
		return resource.PAY_MODE_POST_PAY
	case "1":
		return resource.PAY_MODE_PRE_PAY
	case "2":
		return resource.PAY_MODE_POST_PAY
	}

	return resource.PAY_MODE_NULL
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
			ret = append(ret, resource.NewGroupTag(kv[0], kv[1]))
		} else {
			ret = append(ret, resource.NewGroupTag("ecs", t))
		}
	}

	return
}
