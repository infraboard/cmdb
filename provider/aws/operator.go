package aws

import (
	"github.com/caarlos0/env/v6"
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
	client := &connectivity.AwsCloudClient{}
	if err := env.Parse(client); err != nil {
		return err
	}
	operator = NewOperator(client)
	return nil
}

func NewOperator(client *connectivity.AwsCloudClient) *Operator {
	return &Operator{
		client: client,
	}
}

type Operator struct {
	client *connectivity.AwsCloudClient
}

func (o *Operator) EcsOperator() *ec2.Ec2Operater {
	c, err := o.client.Ec2Client()
	if err != nil {
		panic(err)
	}
	return ec2.NewEc2Operator(c)
}
