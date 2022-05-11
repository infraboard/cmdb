package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/mcube/app"
)

var (
	h = &handler{}
)

type handler struct {
	service secret.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(secret.AppName)
	h.service = app.GetGrpcApp(secret.AppName).(secret.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return secret.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/").To(h.CreateSecret).
		Doc("create a secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(secret.CreateSecretRequest{}).
		Writes(response.NewData(secret.Secret{})))

	ws.Route(ws.GET("/").To(h.QuerySecret).
		Doc("get all secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(secret.QuerySecretRequest{}).
		Writes(response.NewData(secret.SecretSet{})).
		Returns(200, "OK", secret.SecretSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeSecret).
		Doc("describe an secret").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(secret.Secret{})).
		Returns(200, "OK", response.NewData(secret.Secret{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteSecret).
		Doc("delete a secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")))

	// sr.Handle("GET", "/crendential_types", h.ListCrendentialType).DisablePermission()
}

func init() {
	app.RegistryRESTfulApp(h)
}
