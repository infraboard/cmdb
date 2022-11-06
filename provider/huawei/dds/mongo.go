package dds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dds/v3/model"
	"github.com/infraboard/cmdb/apps/mongodb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/huawei/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *DdsOperator) PageQueryMongo(req *provider.QueryRequest) pager.Pager {
	return newPager(o)
}

// 根据指定条件查询实例列表和详情
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=DDS&api=ListInstances
func (o *DdsOperator) Query(req *model.ListInstancesRequest) (*mongodb.MongoDBSet, error) {
	set := mongodb.NewMongoDBSet()

	resp, err := o.client.ListInstances(req)
	if err != nil {
		return nil, err
	}

	// 华为云返回的TotalCount不准确
	set.Total = int64(tea.Int32Value(resp.TotalCount))
	set.Items = o.transferMongoDBSet(resp.Instances).Items

	return set, nil
}

func (o *DdsOperator) transferMongoDBSet(list *[]model.QueryInstanceResponse) *mongodb.MongoDBSet {
	set := mongodb.NewMongoDBSet()
	items := *list
	for i := range items {
		set.Add(o.transferMongoDBOne(items[i]))
	}
	return set
}

func (o *DdsOperator) transferMongoDBOne(ins model.QueryInstanceResponse) *mongodb.MongoDB {
	r := mongodb.NewDefaultMongoDB()
	b := r.Resource.Meta
	b.CreateAt = utils.ParseTime("2006-01-02T15:04:05", ins.Created)
	b.Id = ins.Id

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_HUAWEI
	info.Region = ins.Region
	info.Name = ins.Name

	r.Resource.Status.Phase = praseStatus(ins.Status)

	desc := r.Describe
	desc.Engine = ins.Datastore.Type
	desc.EngineVersion = ins.Datastore.Version

	for _, g := range ins.Groups {
		if g.Type == strings.ToLower(ins.Mode) {
			size, _ := strconv.ParseInt(g.Volume.Size, 10, 32)
			info.Storage = int32(size)
		}

		info.Type = fmt.Sprintf("%s|%s|%d", ins.Mode, g.Nodes[0].SpecCode, len(g.Nodes))

		for _, node := range g.Nodes {
			if strings.ToLower(node.Role) != "hidden" {
				r.Resource.Status.PrivateIp = append(r.Resource.Status.PrivateIp, fmt.Sprintf("%s:%s", node.PrivateIp, ins.Port))
			}
		}
	}

	pm, _ := strconv.ParseInt(tea.StringValue(ins.PayMode), 10, 32)
	m := int32(pm)
	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(&m)
	return r
}
