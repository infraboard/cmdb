package http

import (
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/app/secret"
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

func (h *handler) Registry(r router.SubRouter) {
	sr := r.ResourceRouter("secret")
	sr.Permission(true)
	sr.Handle("POST", "/secrets", h.CreateSecret).AddLabel(label.Create)
	sr.Handle("GET", "/secrets", h.QuerySecret).AddLabel(label.List)
	sr.Handle("GET", "/secrets/:id", h.DescribeSecret).AddLabel(label.Get)
	sr.Handle("DELETE", "/secrets/:id", h.DeleteSecret).AddLabel(label.Delete)
	sr.Handle("GET", "/crendential_types", h.ListCrendentialType).DisablePermission()
}

func init() {
	app.RegistryHttpApp(h)
}
