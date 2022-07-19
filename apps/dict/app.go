package dict

import (
	"github.com/infraboard/cmdb/apps/credential"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"

	ali_region "github.com/infraboard/cmdb/provider/aliyun/region"
	aws_region "github.com/infraboard/cmdb/provider/aws/region"
	hw_region "github.com/infraboard/cmdb/provider/huawei/region"
	tx_region "github.com/infraboard/cmdb/provider/txyun/region"
)

const (
	AppName = "dict"
)

var (
	CrendentialTypes = []utils.EnumDescribe{
		{Value: credential.Type_API_KEY.String(), Describe: "API凭证"},
		{Value: credential.Type_PASSWORD.String(), Describe: "用户名密码"},
	}

	Vendors = []utils.EnumDescribe{
		{Value: resource.Vendor_ALIYUN.String(), Describe: "阿里云"},
		{Value: resource.Vendor_TENCENT.String(), Describe: "腾讯云"},
		{Value: resource.Vendor_HUAWEI.String(), Describe: "华为云"},
		{Value: resource.Vendor_AMAZON.String(), Describe: "AWS"},
		{Value: resource.Vendor_VSPHERE.String(), Describe: "Vsphere"},
	}

	ResourceTypes = map[string][]utils.EnumDescribe{
		resource.Vendor_ALIYUN.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "ECS", Meta: utils.ParamType("region")},
			{Name: "关系型数据库", Value: resource.Type_RDS.String(), Describe: "RDS", Meta: utils.ParamType("region")},
			{Name: "月账单", Value: resource.Type_BILL.String(), Describe: "月账单", Meta: utils.ParamType("month")},
		},
		resource.Vendor_TENCENT.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "CVM", Meta: utils.ParamType("region")},
			{Name: "关系型数据库", Value: resource.Type_RDS.String(), Describe: "CDB", Meta: utils.ParamType("region")},
			{Name: "月账单", Value: resource.Type_BILL.String(), Describe: "月账单", Meta: utils.ParamType("month")},
		},
		resource.Vendor_HUAWEI.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "ECS", Meta: utils.ParamType("region")},
			{Name: "关系型数据库", Value: resource.Type_RDS.String(), Describe: "RDS", Meta: utils.ParamType("region")},
			{Name: "月账单", Value: resource.Type_BILL.String(), Describe: "月账单", Meta: utils.ParamType("month")},
		},
		resource.Vendor_AMAZON.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "EC2", Meta: utils.ParamType("region")},
		},
		resource.Vendor_VSPHERE.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "VM"},
		},
	}

	Regions = map[string][]utils.EnumDescribe{
		resource.Vendor_ALIYUN.String():  ali_region.Regions,
		resource.Vendor_TENCENT.String(): tx_region.Regions,
		resource.Vendor_HUAWEI.String():  hw_region.Regions,
		resource.Vendor_AMAZON.String():  aws_region.Regions,
		resource.Vendor_VSPHERE.String(): hw_region.Regions,
	}
)
