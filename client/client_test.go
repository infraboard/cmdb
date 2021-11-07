package client_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/client"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)
	conf := client.NewDefaultConfig()
	conf.WithClientCredentials("VYizVq1fsK7olinqVHrBvFOl", "qS9FGBoFGRaVfbgeqFVDRcgH7nNJi9fp")
	c, err := client.NewClient(conf)
	if should.NoError(err) {
		rs, err := c.Resource().Search(context.Background(), resource.NewSearchRequest())
		should.NoError(err)
		fmt.Println(rs)
	}
}
