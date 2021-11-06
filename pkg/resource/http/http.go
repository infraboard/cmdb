package http

import (
	"fmt"

	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/pkg"
	"github.com/infraboard/cmdb/pkg/resource"
)

var (
	api = &handler{}
)

type handler struct {
	service resource.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Resource")
	if pkg.Resource == nil {
		return fmt.Errorf("dependence service resource not ready")
	}
	h.service = pkg.Resource
	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.GET("/search", api.SearchResource)
	r.GET("/vendors", api.ListVendor)
	r.GET("/regions", api.ListVendorRegion)
	r.GET("/resource_types", api.ListResourceType)
}
