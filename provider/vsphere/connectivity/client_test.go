package connectivity_test

import (
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/provider/vsphere/connectivity"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)

	err := connectivity.LoadClientFromEnv()
	if should.NoError(err) {
		vim, err := connectivity.C().VimClient()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(vim.Client.URL().Host)
	}
}
