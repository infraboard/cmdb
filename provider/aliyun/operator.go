package aliyun

import (
	"github.com/caarlos0/env/v6"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/bss"
	"github.com/infraboard/cmdb/provider/aliyun/connectivity"
	"github.com/infraboard/cmdb/provider/aliyun/ecs"
	"github.com/infraboard/cmdb/provider/aliyun/rds"
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
	conf := &connectivity.AliCloudClient{}
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
	client := connectivity.NewAliCloudClient(ak, sk, region)

	account, err := client.Account()
	if err != nil {
		return nil, err
	}

	return &Operator{
		client:  client,
		account: account,
	}, nil
}

type Operator struct {
	client  *connectivity.AliCloudClient
	account string
}

func (o *Operator) HostOperator() provider.HostOperator {
	c, err := o.client.EcsClient()
	if err != nil {
		panic(err)
	}
	op := ecs.NewEcsOperator(c)
	op.WithAccountId(o.account)
	return op
}

func (o *Operator) BillOperator() provider.BillOperator {
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
