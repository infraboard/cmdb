package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/cmdb/apps/dict"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcube/app"
)

var (
	h = &handler{}
)

type handler struct {
	log logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(dict.AppName)
	return nil
}

func (h *handler) Name() string {
	return dict.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.GET("/crendential_types").To(h.CrendentialType).
		Doc("get all secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.ResourceLableKey, h.Name()).
		Metadata(label.ActionLableKey, label.List.Value()).
		Writes(response.NewData([]utils.EnumDescribe{})).
		Returns(200, "OK", []utils.EnumDescribe{}))
	ws.Route(ws.GET("/vendors").To(h.ListVendor).
		Doc("get all vendors").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.ResourceLableKey, "vendord").
		Metadata(label.ActionLableKey, label.List.Value()).
		Writes(response.NewData([]utils.EnumDescribe{})).
		Returns(200, "OK", []utils.EnumDescribe{}))

	ws.Route(ws.GET("/regions").To(h.ListVendorRegion).
		Doc("get all regions").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.ResourceLableKey, "regions").
		Metadata(label.ActionLableKey, label.List.Value()).
		Writes(response.NewData(map[string][]utils.EnumDescribe{})).
		Returns(200, "OK", map[string][]utils.EnumDescribe{}))

	ws.Route(ws.GET("/types").To(h.ListResourceType).
		Doc("get all resource types").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.ResourceLableKey, "types").
		Metadata(label.ActionLableKey, label.List.Value()).
		Writes(response.NewData(map[string][]utils.EnumDescribe{})).
		Returns(200, "OK", map[string][]utils.EnumDescribe{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
