package vm

import (
	"context"
	"errors"
	"fmt"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
)

func (o *VMOperater) Query() (*host.HostSet, error) {
	set := host.NewHostSet()

	// 查询DC
	dcs, err := o.finder.DatacenterList(context.Background(), "*")
	if err != nil {
		fmt.Println("xxx")
		return nil, err
	}

	if len(dcs) == 0 {
		return nil, errors.New("not datacenter found")
	}

	for i := range dcs {
		dc := dcs[i]
		o.log.Debugf("query dc %s vms ...", dc.Name())

		vms, err := o.queryVM(dc)
		if err != nil {
			o.log.Errorf("query dc %s error, %s", dc.Name())
		}

		fmt.Println(vms)
	}

	return set, nil
}

func (o *VMOperater) queryVM(dc *object.Datacenter) (*host.HostSet, error) {
	dcFinder := o.finder.SetDatacenter(dc)

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()
	items, err := dcFinder.VirtualMachineList(ctx, "*")
	if err != nil {
		return nil, err
	}

	set := host.NewHostSet()

	o.log.Debugf("total vms in dc %s, %d", dc.Name(), len(items))
	for _, item := range items {
		vmp, err := o.properties(item)
		if err != nil {
			return nil, err
		}

		// 排除模板
		if !vmp.Config.Template {
			set.Add(o.transferOne(vmp, dc.Name()))
		}
	}

	return set, nil
}

// Properties is a convenience method that wraps fetching the
// VirtualMachine MO from its higher-level object.
func (o *VMOperater) properties(vm *object.VirtualMachine) (*mo.VirtualMachine, error) {
	o.log.Debugf("Fetching properties for VM %q", vm.InventoryPath)
	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	var props mo.VirtualMachine
	if err := vm.Properties(ctx, vm.Reference(), nil, &props); err != nil {
		return nil, err
	}

	return &props, nil
}
