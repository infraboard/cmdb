package connectivity_test

import (
	"fmt"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/cmdb/provider/aliyun/connectivity"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)

	client := &connectivity.AliCloudClient{}
	if err := env.Parse(client); err != nil {
		if should.NoError(err) {
			client.Check()
			fmt.Println(client.AccountID())
		}
	}

}
