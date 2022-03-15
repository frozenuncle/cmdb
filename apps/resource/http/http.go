package http

import (
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/app"
)

var (
	h = &handler{}
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

func (h *handler) Name() string {
	return resource.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	rr := r.ResourceRouter("resource")
	rr.Handle("GET", "/search", h.SearchResource).AddLabel(label.List)
	rr.Handle("GET", "/vendors", h.ListVendor).DisablePermission()
	rr.Handle("GET", "/regions", h.ListVendorRegion).DisablePermission()
	rr.Handle("GET", "/resource_types", h.ListResourceType).DisablePermission()

	// 资源标签管理
	rr.Handle("POST", "/resources/:id/tags", h.AddTag).AddLabel(label.Update)
	rr.Handle("DELETE", "/resources/:id/tags", h.RemoveTag).AddLabel(label.Update)

	// 资源发现
	rr.Handle("GET", "/discovery/prometheus", h.DiscoveryPrometheus).DisableAuth()
}

func init() {
	app.RegistryHttpApp(h)
}
