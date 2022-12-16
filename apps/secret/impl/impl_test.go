package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/cmdb/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl secret.Service
	ctx  = context.Background()
)

func TestQuerySecret(t *testing.T) {
	req := secret.NewQuerySecretRequest()
	set, err := impl.QuerySecret(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func init() {
	tools.DevelopmentSetup()

	impl = app.GetInternalApp(secret.AppName).(secret.Service)
}
