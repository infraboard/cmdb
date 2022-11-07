package impl_test

import (
	"context"
	"testing"
	"time"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl resource.Service
	ctx  context.Context
)

func TestPut(t *testing.T) {
	res := resource.NewDefaultResource(resource.TYPE_HOST)
	res.Meta.Id = "test01"
	res.Meta.Domain = "default"
	res.Meta.Namespace = "default"
	res.Meta.SyncAt = time.Now().Unix()
	ins, err := impl.Put(ctx, res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(resource.AppName).(resource.Service)
}
