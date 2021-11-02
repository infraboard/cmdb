package http

import (
	"fmt"

	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/pkg"
	"github.com/infraboard/cmdb/pkg/syncer"
)

var (
	api = &handler{}
)

type handler struct {
	service syncer.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Syncer")
	if pkg.Syncer == nil {
		return fmt.Errorf("dependence service syncer not ready")
	}
	h.service = pkg.Syncer
	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.POST("/secrets", api.CreateSecret)
	r.GET("/secrets", api.QuerySecret)
	r.GET("/secrets/:id", api.DescribeSecret)
	r.DELETE("/secrets/:id", api.DeleteSecret)
	r.POST("/secrets/:id/sync", api.Sync)
	r.GET("/crendential_types", api.ListCrendentialType)
}
