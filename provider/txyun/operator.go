package txyun

import (
	"github.com/caarlos0/env/v6"
	"github.com/infraboard/cmdb/provider/txyun/billing"
	"github.com/infraboard/cmdb/provider/txyun/cdb"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/cmdb/provider/txyun/cvm"
)

var (
	operator *Operator
)

func O() *Operator {
	if operator == nil {
		panic("please load config first")
	}
	return operator
}

func LoadOperatorFromEnv() error {
	client := &connectivity.TencentCloudClient{}
	if err := env.Parse(client); err != nil {
		return err
	}
	operator = NewOperator(client)
	return nil
}

func NewOperator(client *connectivity.TencentCloudClient) *Operator {
	return &Operator{
		client: client,
	}
}

type Operator struct {
	client *connectivity.TencentCloudClient
}

func (o *Operator) CvmOperator() *cvm.CVMOperator {
	return cvm.NewCVMOperator(o.client.CvmClient())
}

func (o *Operator) CdbOperator() *cdb.CDBOperator {
	return cdb.NewCDBOperator(o.client.CDBClient())
}

func (o *Operator) BillOperator() *billing.BillingOperater {
	return billing.NewBillingOperater(o.client.BillingClient())
}
