package txyun

import (
	"github.com/caarlos0/env/v6"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/billing"
	"github.com/infraboard/cmdb/provider/txyun/cdb"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/cmdb/provider/txyun/cos"
	"github.com/infraboard/cmdb/provider/txyun/cvm"
	"github.com/infraboard/cmdb/provider/txyun/redis"
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
	conf := &connectivity.TencentCloudClient{}
	if err := env.Parse(conf); err != nil {
		return err
	}
	op, err := NewOperator(conf.SecretID, conf.SecretKey, conf.Region)
	if err != nil {
		return err
	}
	operator = op
	return nil
}

func NewOperator(credentialID, credentialKey, region string) (*Operator, error) {
	client := connectivity.NewTencentCloudClient(credentialID, credentialKey, region)

	account, err := client.Account()
	if err != nil {
		return nil, err
	}

	return &Operator{
		account: account,
		client:  client,
	}, nil
}

type Operator struct {
	account string
	client  *connectivity.TencentCloudClient
}

func (o *Operator) Client() *connectivity.TencentCloudClient {
	return o.client
}

func (o *Operator) HostOperator() provider.HostOperator {
	op := cvm.NewCVMOperator(o.client.CvmClient(), o.client.CBSClient())
	op.WithAccountId(o.account)
	return op
}

func (o *Operator) RdsOperator() provider.RdsOperator {
	return cdb.NewCDBOperator(o.client.CDBClient())
}

func (o *Operator) RedisOperator() provider.RedisOperator {
	op := redis.NewRedisOperator(o.client.RedisClient())
	return op
}

func (o *Operator) OssOperator() provider.OssOperator {
	op := cos.NewCosOperator(o.client.CosClient())
	return op
}

func (o *Operator) BillOperator() provider.BillOperator {
	return billing.NewBillingoperator(o.client.BillingClient())
}
