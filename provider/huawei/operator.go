package huawei

import (
	"github.com/caarlos0/env/v6"
	"github.com/infraboard/cmdb/provider/huawei/bss"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/cmdb/provider/huawei/ecs"
	"github.com/infraboard/cmdb/provider/huawei/rds"
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
	client := &connectivity.HuaweiCloudClient{}
	if err := env.Parse(client); err != nil {
		return err
	}
	operator = NewOperator(client)
	return nil
}

func NewOperator(client *connectivity.HuaweiCloudClient) *Operator {
	return &Operator{
		client: client,
	}
}

type Operator struct {
	client *connectivity.HuaweiCloudClient
}

func (o *Operator) EcsOperator() *ecs.EcsOperator {
	c, err := o.client.EcsClient()
	if err != nil {
		panic(err)
	}
	return ecs.NewEcsOperator(c)
}

func (o *Operator) BssOperator() *bss.BssOperator {
	c, err := o.client.BssClient()
	if err != nil {
		panic(err)
	}
	return bss.NewBssOperator(c)
}

func (o *Operator) RdsOperator() *rds.RdsOperator {
	c, err := o.client.RdsClient()
	if err != nil {
		panic(err)
	}
	return rds.NewRdsOperator(c)
}
