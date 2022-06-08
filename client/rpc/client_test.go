package rpc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/client/rpc"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)

	c, err := rpc.NewClientSet(conf.C().Mcenter)
	if should.NoError(err) {
		rs, err := c.Resource().Search(context.Background(), resource.NewSearchRequest())
		should.NoError(err)
		fmt.Println(rs)
	}
}

func init() {
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
}
