package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type Ec2operator struct {
	client *ec2.Client
	log    logger.Logger
}

// NewEc2Operator Ec2Operator
func NewEc2Operator(client *ec2.Client) *Ec2operator {
	return &Ec2operator{
		client: client,
		log:    zap.L().Named("AWS EC2"),
	}
}
