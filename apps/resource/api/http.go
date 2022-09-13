package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/app"
)

var (
	h = &handler{}
)

type handler struct {
	service resource.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(resource.AppName)
	h.service = app.GetGrpcApp(resource.AppName).(resource.Service)
	return nil
}

func (h *handler) Name() string {
	return resource.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.GET("/search").To(h.SearchResource).
		Doc("get all resources").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Reads(resource.SearchRequest{}).
		Writes(response.NewData(resource.ResourceSet{})).
		Returns(200, "OK", resource.ResourceSet{}))

	// 资源标签管理
	ws.Route(ws.POST("/tags").To(h.AddTag).
		Doc("add resource tags").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "tags").
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Reads([]*resource.Tag{}).
		Writes(response.NewData(resource.Resource{})))
	ws.Route(ws.DELETE("/tags").To(h.RemoveTag).
		Doc("remove resource tags").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "tags").
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Reads([]*resource.Tag{}).
		Writes(response.NewData(resource.Resource{})))

	// 资源发现
	ws.Route(ws.GET("/discovery/prometheus").To(h.DiscoveryPrometheus).
		Doc("discovery resoruce for prometheus").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "prometheus_resource").
		Metadata(label.Action, label.List.Value()).
		Reads(resource.SearchRequest{}).
		Writes(response.NewData(resource.ResourceSet{})).
		Returns(200, "OK", resource.ResourceSet{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
