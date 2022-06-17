package huawei

import (
	"github.com/caarlos0/env/v6"
	"github.com/infraboard/cmdb/provider"
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
	conf := &connectivity.HuaweiCloudClient{}
	if err := env.Parse(conf); err != nil {
		return err
	}
	op, err := NewOperator(conf.AccessKey, conf.AccessSecret, conf.Region)
	if err != nil {
		return err
	}
	operator = op
	return nil
}

func NewOperator(ak, sk, region string) (*Operator, error) {
	client := connectivity.NewHuaweiCloudClient(ak, sk, region)
	return &Operator{
		client: client,
	}, nil
}

type Operator struct {
	client *connectivity.HuaweiCloudClient
}

func (o *Operator) HostOperator() provider.HostOperator {
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

func (o *Operator) RdsOperator() provider.RdsOperator {
	c, err := o.client.RdsClient()
	if err != nil {
		panic(err)
	}
	return rds.NewRdsOperator(c)
}
