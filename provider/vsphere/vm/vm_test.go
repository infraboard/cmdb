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

func TestRangeString(t *testing.T) {
	str := "architecture='X86' bitness='64' distroName='CentOS Stream' distroVersion='8' familyName='Linux' kernelVersion='4.18.0-365.el8.x86_64' prettyName='CentOS Stream 8'"
	m := op.ParseExtraConfigValue(str)
	for k, v := range m {
		fmt.Println(k, v)
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
