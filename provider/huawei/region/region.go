package region

import "github.com/infraboard/cmdb/utils"

var Regions = []utils.EnumDescribe{
	{Name: "CN_NORTH_1", Value: "cn-north-1", Describe: "华北-北京一"},
	{Name: "CN_NORTH_4", Value: "cn-north-4", Describe: "华北-北京四"},
	{Name: "CN_NORTH_9", Value: "cn-north-9", Describe: "华北-乌兰察布一"},
	{Name: "CN_EAST_3", Value: "cn-east-3", Describe: "华东-上海一"},
	{Name: "CN_SOUTH_1", Value: "cn-south-1", Describe: "华南-广州"},
	{Name: "CN_SOUTH_4", Value: "cn-south-4", Describe: "华南-广州-友好用户环境"},
	{Name: "CN_EAST_2", Value: "cn-east-2", Describe: "华东-上海二"},
	{Name: "CN_SOUTHWEST_2", Value: "cn-southwest-2", Describe: "西南-贵阳一"},
	{Name: "AP_SOUTHEAST_1", Value: "ap-southeast-1", Describe: "中国-香港"},
	{Name: "AP_SOUTHEAST_2", Value: "ap-southeast-2", Describe: "亚太-曼谷"},
	{Name: "AP_SOUTHEAST_3", Value: "ap-southeast-3", Describe: "亚太-新加坡"},
	{Name: "AF_SOUTH_1", Value: "af-south-1", Describe: "非洲-约翰内斯堡"},
	{Name: "NA_MEXICO_1", Value: "na-mexico-1", Describe: "拉美-墨西哥城一"},
	{Name: "LA_NORTH_1", Value: "la-north-2", Describe: "拉美-墨西哥城二"},
	{Name: "SA_BRAZIL_1", Value: "sa-brazil-1", Describe: "拉美-圣保罗一"},
	{Name: "LA_SOUTH_2", Value: "la-south-2", Describe: "拉美-圣地亚哥"},
}
