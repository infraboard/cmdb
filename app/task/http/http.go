package http

import (
	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	pkg "github.com/infraboard/cmdb/app"
	"github.com/infraboard/cmdb/app/task"
)

var (
	api = &handler{}
)

type handler struct {
	task task.ServiceServer
	log  logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(task.AppName)
	h.task = pkg.GetGrpcApp(task.AppName).(task.ServiceServer)
	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.GET("/tasks", api.QueryTask)
	r.POST("/tasks", api.CreatTask)
}
