package region

import "github.com/infraboard/cmdb/utils"

var Regions = []utils.EnumDescribe{
	{Name: "Bangkok", Value: "ap-bangkok", Describe: "曼谷"},
	{Name: "Beijing", Value: "ap-beijing", Describe: "北京"},
	{Name: "Chengdu", Value: "ap-chengdu", Describe: "成都"},
	{Name: "Chongqing", Value: "ap-chongqing", Describe: "重庆"},
	{Name: "Guangzhou", Value: "ap-guangzhou", Describe: "广州"},
	{Name: "GuangzhouOpen", Value: "ap-guangzhou-open", Describe: "广州Open"},
	{Name: "HongKong", Value: "ap-hongkong", Describe: "中国香港"},
	{Name: "Mumbai", Value: "ap-mumbai", Describe: "孟买"},
	{Name: "Seoul", Value: "ap-seoul", Describe: "首尔"},
	{Name: "Shanghai", Value: "ap-shanghai", Describe: "上海"},
	{Name: "Nanjing", Value: "ap-nanjing", Describe: "南京"},
	{Name: "ShanghaiFSI", Value: "ap-shanghai-fsi", Describe: "上海金融"},
	{Name: "ShenzhenFSI", Value: "ap-shenzhen-fsi", Describe: "深圳金融"},
	{Name: "Singapore", Value: "ap-singapore", Describe: "新加坡"},
	{Name: "Tokyo", Value: "ap-tokyo", Describe: "东京"},
	{Name: "Frankfurt", Value: "eu-frankfurt", Describe: "法兰克福"},
	{Name: "Moscow", Value: "eu-moscow", Describe: "莫斯科"},
	{Name: "Ashburn", Value: "na-ashburn", Describe: "阿什本"},
	{Name: "SiliconValley", Value: "na-siliconvalley", Describe: "硅谷"},
	{Name: "Toronto", Value: "na-toronto", Describe: "多伦多"},
}
