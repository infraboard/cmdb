package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
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
	tags := []string{"资源检索"}

	ws.Route(ws.GET("/search").To(h.SearchResource).
		Doc("检索资源").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Reads(resource.SearchRequest{}).
		Writes(resource.ResourceSet{}).
		Returns(200, "OK", resource.ResourceSet{}))

	// 资源发现
	ws.Route(ws.GET("/discovery/prometheus").To(h.DiscoveryPrometheus).
		Doc("监控发现(Prometheus)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "prometheus_resource").
		Metadata(label.Action, label.List.Value()).
		Reads(resource.SearchRequest{}).
		Writes(resource.ResourceSet{}).
		Returns(200, "OK", resource.ResourceSet{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
