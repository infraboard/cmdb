package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/apps/task"
	"github.com/infraboard/mcube/app"
)

var (
	h = &handler{}
)

type handler struct {
	task task.ServiceServer
	log  logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(task.AppName)
	h.task = app.GetGrpcApp(task.AppName).(task.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return task.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/").To(h.CreatTask).
		Doc("create a task").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "task").
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(task.CreateTaskRequst{}).
		Writes(response.NewData(task.Task{})))

	ws.Route(ws.GET("/").To(h.QueryTask).
		Doc("get all task").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "task").
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(task.QueryTaskRecordRequest{}).
		Writes(response.NewData(task.TaskSet{})).
		Returns(200, "OK", task.TaskSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeTask).
		Doc("describe an task").
		Param(ws.PathParameter("id", "identifier of the task").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "task").
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(response.NewData(task.Task{})).
		Returns(200, "OK", response.NewData(task.Task{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}/records").To(h.DescribeTaskRecord).
		Doc("get task records").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, "task_records").
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(task.QueryTaskRecordRequest{}).
		Writes(response.NewData(task.RecordSet{})).
		Returns(200, "OK", task.RecordSet{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
