package vm_test

import (
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider/vsphere/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/vsphere/vm"
)

var (
	operator *op.VMOperator
)

func TestQuery(t *testing.T) {
	err := operator.QueryHost(func(h *host.Host) {
		fmt.Println(h)
	})
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	zap.DevelopmentSetup()

	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	vim, err := connectivity.C().VimClient()
	if err != nil {
		panic(err)
	}

	operator = op.NewVMOperator(vim)
}
