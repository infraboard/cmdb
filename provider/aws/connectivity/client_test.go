package connectivity_test

import (
	"testing"

	"github.com/infraboard/cmdb/provider/aws/connectivity"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)

	err := connectivity.LoadClientFromEnv()
	if should.NoError(err) {
		_, err := connectivity.C().Ec2Client()
		should.NoError(err)
	}
}
