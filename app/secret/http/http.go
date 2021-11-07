package http

import (
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/app"
	"github.com/infraboard/cmdb/app/secret"
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
	r.Handle("POST", "/secrets", h.CreateSecret)
	r.Handle("GET", "/secrets", h.QuerySecret)
	r.Handle("GET", "/secrets/:id", h.DescribeSecret)
	r.Handle("DELETE", "/secrets/:id", h.DeleteSecret)
	r.Handle("GET", "/crendential_types", h.ListCrendentialType)
}

func init() {
	app.RegistryHttpApp(h)
}
