package impl_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/task"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"

	_ "github.com/infraboard/cmdb/apps/all"
)

var (
	svc task.ServiceServer
)

func TestSyncBill(t *testing.T) {
	req := task.NewCreateTaskRequst()
	req.Type = task.Type_RESOURCE_SYNC
	req.ResourceType = resource.Type_BILL
	req.SecretId = "c5pcffua0bro7e7a05j0"
	ins, err := svc.CreatTask(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ins)
	time.Sleep(5 * time.Second)
}

func init() {
	zap.DevelopmentSetup()
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
	svc = app.GetGrpcApp(task.AppName).(task.ServiceServer)
}
