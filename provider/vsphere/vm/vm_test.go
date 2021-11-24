package vm_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/app/host"
	"github.com/infraboard/cmdb/provider/vsphere/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/vsphere/vm"
)

var (
	operater *op.VMOperater
)

func TestQuery(t *testing.T) {
	err := operater.Query(func(h *host.Host) {
		fmt.Println(h)
	})
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	zap.DevelopmentSetup()

	var host, username, password string
	if host = os.Getenv("VS_HOST"); host == "" {
		panic("empty VS_HOST")
	}

	if username = os.Getenv("VS_USERNAME"); username == "" {
		panic("empty VS_USERNAME")
	}

	if password = os.Getenv("VS_PASSWORD"); password == "" {
		panic("empty VS_PASSWORD")
	}

	client := connectivity.NewVsphereClient(host, username, password)
	vim, err := client.VimClient()
	if err != nil {
		panic(err)
	}

	operater = op.NewVmOperater(vim)
}
