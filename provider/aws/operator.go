package aws

import (
	"github.com/caarlos0/env/v6"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aws/connectivity"
	"github.com/infraboard/cmdb/provider/aws/ec2"
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
	conf := &connectivity.AwsCloudClient{}
	if err := env.Parse(conf); err != nil {
		return err
	}
	operator = NewOperator(conf.AccessKey, conf.AccessSecret, conf.Region)
	return nil
}

func NewOperator(ak, sk, region string) *Operator {
	client := connectivity.NewAwsCloudClient(ak, sk, region)
	return &Operator{
		client: client,
	}
}

type Operator struct {
	client *connectivity.AwsCloudClient
}

func (o *Operator) HostOperator() provider.HostOperator {
	c, err := o.client.Ec2Client()
	if err != nil {
		panic(err)
	}
	return ec2.NewEc2Operator(c)
}
