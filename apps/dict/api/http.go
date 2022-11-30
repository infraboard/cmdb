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
	tags := []string{h.Name()}

	ws.Route(ws.GET("/crendential_types").To(h.CrendentialType).
		Doc("get all credential").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "crendential_types").
		Metadata(label.Action, label.List.Value()).
		Writes(dict.CrendentialTypes).
		Returns(200, "OK", dict.CrendentialTypes))

	ws.Route(ws.GET("/vendors").To(h.Vendor).
		Doc("get all vendors").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "vendord").
		Metadata(label.Action, label.List.Value()).
		Writes(dict.Vendors).
		Returns(200, "OK", dict.Vendors))

	ws.Route(ws.GET("/regions").To(h.VendorRegion).
		Doc("get all regions").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "regions").
		Metadata(label.Action, label.List.Value()).
		Writes(dict.Regions).
		Returns(200, "OK", dict.Regions))

	ws.Route(ws.GET("/resource_types").To(h.ResourceType).
		Doc("get all resource types").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "types").
		Metadata(label.Action, label.List.Value()).
		Writes(dict.ResourceTypes).
		Returns(200, "OK", dict.ResourceTypes))
}

func init() {
	app.RegistryRESTfulApp(h)
}
