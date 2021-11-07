package http

import (
	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/app"
	"github.com/infraboard/cmdb/app/host"
)

var (
	api = &handler{}
)

type handler struct {
	service host.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(host.AppName)
	h.service = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.GET("/hosts", api.QueryHost)
	r.POST("/hosts", api.CreateHost)
	r.GET("/hosts/:id", api.DescribeHost)
	r.DELETE("/hosts/:id", api.DeleteHost)
	r.PUT("/hosts/:id", api.PutHost)
	r.PATCH("/hosts/:id", api.PatchHost)
}
