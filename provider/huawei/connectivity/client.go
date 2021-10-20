package connectivity

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	dcs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	rds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3"
)

// NewHuaweiCloudClient client
func NewHuaweiCloudClient(ak, sk, region string) *HuaweiCloudClient {
	return &HuaweiCloudClient{
		Region:       region,
		AccessKey:    ak,
		AccessSecret: sk,
	}
}

type HuaweiCloudClient struct {
	Region       string
	AccessKey    string
	AccessSecret string
	accountId    string
	ecsConn      *ecs.EcsClient
	rdsConn      *rds.RdsClient
	dcsConn      *dcs.DcsClient
}

func (c *HuaweiCloudClient) Credentials() basic.Credentials {
	auth := basic.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.AccessSecret).
		Build()
	return auth
}

// EcsClient 客户端
func (c *HuaweiCloudClient) EcsClient() (*ecs.EcsClient, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client := ecs.EcsClientBuilder().
		WithRegion(region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.ecsConn = ecs.NewEcsClient(client)

	return c.ecsConn, nil
}

// RdsClient 客户端
func (c *HuaweiCloudClient) RdsClient() (*rds.RdsClient, error) {
	if c.rdsConn != nil {
		return c.rdsConn, nil
	}

	client := rds.RdsClientBuilder().
		WithRegion(region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.rdsConn = rds.NewRdsClient(client)

	return c.rdsConn, nil
}

// DcsClient 客户端
func (c *HuaweiCloudClient) DcsClient() (*dcs.DcsClient, error) {
	if c.dcsConn != nil {
		return c.dcsConn, nil
	}

	client := dcs.DcsClientBuilder().
		WithRegion(region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.dcsConn = dcs.NewDcsClient(client)

	return c.dcsConn, nil
}
