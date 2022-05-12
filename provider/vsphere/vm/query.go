package vm

import (
	"context"
	"errors"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
)

type QueryCallback func(*host.Host)

func (o *VMOperator) Query(cb QueryCallback) error {
	// 查询DC
	dcs, err := o.finder.DatacenterList(context.Background(), "*")
	if err != nil {
		return err
	}

	if len(dcs) == 0 {
		return errors.New("not datacenter found")
	}

	for i := range dcs {
		dc := dcs[i]
		o.log.Debugf("query dc %s vms ...", dc.Name())

		err := o.queryVM(dc, cb)
		if err != nil {
			o.log.Errorf("query dc %s error, %s", dc.Name())
		}
	}

	return nil
}

func (o *VMOperator) queryVM(dc *object.Datacenter, cb QueryCallback) error {
	dcFinder := o.finder.SetDatacenter(dc)

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()
	items, err := dcFinder.VirtualMachineList(ctx, "*")
	if err != nil {
		return err
	}

	o.log.Debugf("total vms in dc %s, %d", dc.Name(), len(items))
	for _, item := range items {
		vmp, err := o.properties(item)
		if err != nil {
			return err
		}

		// 排除模板
		if !vmp.Config.Template {
			cb(o.transferOne(vmp, dc.Name()))
		}
	}

	return nil
}

// Properties is a convenience method that wraps fetching the
// VirtualMachine MO from its higher-level object.
func (o *VMOperator) properties(vm *object.VirtualMachine) (*mo.VirtualMachine, error) {
	o.log.Debugf("Fetching properties for VM %q", vm.InventoryPath)
	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	var props mo.VirtualMachine
	if err := vm.Properties(ctx, vm.Reference(), nil, &props); err != nil {
		return nil, err
	}

	return &props, nil
}
