package rpc_test

import (
	"context"
	"testing"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/client/rpc"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	client *rpc.ClientSet
	ctx    = context.Background()
)

func TestClient(t *testing.T) {
	rs, err := client.Resource().Search(ctx, resource.NewSearchRequest())
	if err != nil {
		if e, ok := err.(exception.APIException); ok {
			t.Fatal(e.ToJson())
		} else {
			t.Fatal(err)
		}
	}

	t.Log(rs)
}

func init() {
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	c, err := rpc.NewClientSet(conf.C().Mcenter)
	if err != nil {
		panic(err)
	}
	client = c
}
