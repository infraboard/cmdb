package connectivity

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// NewAliCloudClient client
func NewAliCloudClient(ak, sk, region string) *AliCloudClient {
	return &AliCloudClient{
		Region:       region,
		AccessKey:    ak,
		AccessSecret: sk,
	}
}

type AliCloudClient struct {
	Region       string
	AccessKey    string
	AccessSecret string
	ecsConn      *ecs.Client
}

// EcsClient 客户端
func (c *AliCloudClient) EcsClient() (*ecs.Client, error) {
	client, err := ecs.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	if err != nil {
		return nil, err
	}
	return client, nil
}
