package connectivity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// AwsCloudClient client for all Amazon Cloud service
type AwsCloudClient struct {
	Region       string
	AccessKey    string
	AccessSecret string
	ec2Conn      *ec2.Client
}

// NewAwsCloudClient client
func NewAwsCloudClient(ak, sk, region string) *AwsCloudClient {
	return &AwsCloudClient{
		Region:       region,
		AccessKey:    ak,
		AccessSecret: sk,
	}
}

func (as *AwsCloudClient) Ec2Client() (*ec2.Client, error) {
	if as.ec2Conn != nil {
		return as.ec2Conn, nil
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(as.AccessKey, as.AccessSecret, "")),
		config.WithRegion(as.Region),
	)
	if err != nil {
		return nil, err
	}
	as.ec2Conn = ec2.NewFromConfig(cfg)
	return as.ec2Conn, nil
}
