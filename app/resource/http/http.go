package http

import (
	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/app"
	"github.com/infraboard/cmdb/app/resource"
)

var (
	api = &handler{}
)

type handler struct {
	service resource.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(resource.AppName)
	h.service = app.GetGrpcApp(resource.AppName).(resource.ServiceServer)
	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.GET("/search", api.SearchResource)
	r.GET("/vendors", api.ListVendor)
	r.GET("/regions", api.ListVendorRegion)
	r.GET("/resource_types", api.ListResourceType)
}
