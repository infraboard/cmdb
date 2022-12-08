package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/cmdb/apps/dict"
	"github.com/infraboard/mcube/http/label"
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
	tags := []string{"数据字典"}

	ws.Route(ws.GET("/crendential_types").To(h.CrendentialType).
		Doc("凭证类型").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "crendential_types").
		Metadata(label.Action, label.List.Value()).
		Writes(dict.CrendentialTypes).
		Returns(200, "OK", dict.CrendentialTypes))

	ws.Route(ws.GET("/vendors").To(h.Vendor).
		Doc("资源厂商类型").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "vendord").
		Metadata(label.Action, label.List.Value()).
		Writes(dict.Vendors).
		Returns(200, "OK", dict.Vendors))

	ws.Route(ws.GET("/regions").To(h.VendorRegion).
		Doc("资源厂商地域(Region)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "regions").
		Metadata(label.Action, label.List.Value()).
		Writes(dict.Regions).
		Returns(200, "OK", dict.Regions))

	ws.Route(ws.GET("/resource_types").To(h.ResourceType).
		Doc("资源类型").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "types").
		Metadata(label.Action, label.List.Value()).
		Writes(dict.ResourceTypes).
		Returns(200, "OK", dict.ResourceTypes))
}

func init() {
	app.RegistryRESTfulApp(h)
}
