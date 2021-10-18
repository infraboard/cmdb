package connectivity

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
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
	accountId    string
	ecsConn      *ecs.Client
}

// EcsClient 客户端
func (c *AliCloudClient) EcsClient() (*ecs.Client, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client, err := ecs.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	if err != nil {
		return nil, err
	}

	c.ecsConn = client

	return client, nil
}

// 获取客户端账号ID
func (c *AliCloudClient) AccountID() (string, error) {
	if c.accountId != "" {
		return c.accountId, nil
	}

	args := sts.CreateGetCallerIdentityRequest()

	stsClient, err := sts.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	stsClient.GetConfig().WithScheme("HTTPS")

	if err != nil {
		return "", fmt.Errorf("unable to initialize the STS client: %#v", err)
	}
	stsClient.AppendUserAgent("Infraboard", "1.0")
	identity, err := stsClient.GetCallerIdentity(args)
	if err != nil {
		return "", err
	}
	if identity == nil {
		return "", fmt.Errorf("caller identity not found")
	}
	c.accountId = identity.AccountId

	return c.accountId, nil
}
