package http

import (
	"net/http"

	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"

	ali_region "github.com/infraboard/cmdb/provider/aliyun/region"
	hw_region "github.com/infraboard/cmdb/provider/huawei/region"
	tx_region "github.com/infraboard/cmdb/provider/txyun/region"
)

func (h *handler) ListVendor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := []utils.EnumDescribe{
		{Value: resource.VendorAliYun.String(), Describe: "阿里云"},
		{Value: resource.VendorTencent.String(), Describe: "腾讯云"},
		{Value: resource.VendorHuaWei.String(), Describe: "华为云"},
		{Value: resource.VendorVsphere.String(), Describe: "Vsphere"},
	}
	response.Success(w, resp)
}

func (h *handler) ListVendorRegion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := map[string][]utils.EnumDescribe{
		resource.VendorAliYun.String():  ali_region.Regions,
		resource.VendorTencent.String(): tx_region.Regions,
		resource.VendorHuaWei.String():  hw_region.Regions,
	}
	response.Success(w, resp)
}
