package connectivity

import (
	"github.com/caarlos0/env/v6"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	bss "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"
	dcs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2"
	dcs_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/region"
	dns "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2"
	dns_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/region"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecs_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	eip "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2"
	eip_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/region"
	elb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2"
	elb_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/region"
	evs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2"
	evs_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/region"
	iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	rds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3"
	rds_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/region"

	iam_model "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	iam_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
)

var (
	client *HuaweiCloudClient
)

func C() *HuaweiCloudClient {
	if client == nil {
		panic("please load config first")
	}
	return client
}

func LoadClientFromEnv() error {
	client = &HuaweiCloudClient{}
	if err := env.Parse(client); err != nil {
		return err
	}

	return nil
}

// NewHuaweiCloudClient client
func NewHuaweiCloudClient(ak, sk, region string) *HuaweiCloudClient {
	return &HuaweiCloudClient{
		Region:       region,
		AccessKey:    ak,
		AccessSecret: sk,
	}
}

type HuaweiCloudClient struct {
	Region       string `env:"HW_CLOUD_REGION"`
	AccessKey    string `env:"HW_CLOUD_ACCESS_KEY"`
	AccessSecret string `env:"HW_CLOUD_ACCESS_SECRET"`

	ecsConn *ecs.EcsClient
	rdsConn *rds.RdsClient
	dcsConn *dcs.DcsClient
	bssConn *bss.BssClient
	evsConn *evs.EvsClient
	eipConn *eip.EipClient
	elbConn *elb.ElbClient
	dnsConn *dns.DnsClient
	iamConn *iam.IamClient
}

func (c *HuaweiCloudClient) Credentials() basic.Credentials {
	auth := basic.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.AccessSecret).
		Build()
	return auth
}

func (c *HuaweiCloudClient) GlobalCredentials() global.Credentials {
	auth := global.NewCredentialsBuilder().
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
		WithRegion(ecs_region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.ecsConn = ecs.NewEcsClient(client)

	return c.ecsConn, nil
}

// EcsClient 客户端
func (c *HuaweiCloudClient) EvsClient() (*evs.EvsClient, error) {
	if c.evsConn != nil {
		return c.evsConn, nil
	}

	client := evs.EvsClientBuilder().
		WithRegion(evs_region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.evsConn = evs.NewEvsClient(client)

	return c.evsConn, nil
}

// EipClient 客户端
func (c *HuaweiCloudClient) EipClient() (*eip.EipClient, error) {
	if c.eipConn != nil {
		return c.eipConn, nil
	}

	client := evs.EvsClientBuilder().
		WithRegion(eip_region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.eipConn = eip.NewEipClient(client)

	return c.eipConn, nil
}

// ElbClient 客户端
func (c *HuaweiCloudClient) ElbClient() (*elb.ElbClient, error) {
	if c.elbConn != nil {
		return c.elbConn, nil
	}

	client := evs.EvsClientBuilder().
		WithRegion(elb_region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.elbConn = elb.NewElbClient(client)

	return c.elbConn, nil
}

// RdsClient 客户端
func (c *HuaweiCloudClient) RdsClient() (*rds.RdsClient, error) {
	if c.rdsConn != nil {
		return c.rdsConn, nil
	}

	client := rds.RdsClientBuilder().
		WithRegion(rds_region.ValueOf(c.Region)).
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
		WithRegion(dcs_region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.dcsConn = dcs.NewDcsClient(client)

	return c.dcsConn, nil
}

// DcsClient 客户端
func (c *HuaweiCloudClient) BssClient() (*bss.BssClient, error) {
	if c.bssConn != nil {
		return c.bssConn, nil
	}

	client := bss.BssClientBuilder().
		WithEndpoint("bss.myhuaweicloud.com").
		WithCredential(c.GlobalCredentials()).
		Build()

	c.bssConn = bss.NewBssClient(client)
	return c.bssConn, nil
}

// EcsClient 客户端
func (c *HuaweiCloudClient) DnsClient() (*dns.DnsClient, error) {
	if c.dnsConn != nil {
		return c.dnsConn, nil
	}

	auth := basic.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.AccessSecret).
		Build()

	c.dnsConn = dns.NewDnsClient(
		dns.DnsClientBuilder().
			WithRegion(dns_region.ValueOf(c.Region)).
			WithCredential(auth).
			Build())

	return c.dnsConn, nil
}

// IamClient 客户端
func (c *HuaweiCloudClient) IamClient() (*iam.IamClient, error) {
	if c.iamConn != nil {
		return c.iamConn, nil
	}
	client := iam.IamClientBuilder().
		WithRegion(iam_region.ValueOf(c.Region)).
		WithCredential(c.GlobalCredentials()).
		Build()

	c.iamConn = iam.NewIamClient(client)
	return c.iamConn, nil
}

// IamClient 客户端
func (c *HuaweiCloudClient) Account() (string, error) {
	iamConn, err := c.IamClient()
	if err != nil {
		return "", err
	}

	req := &iam_model.ShowUserRequest{}
	resp, err := iamConn.ShowUser(req)
	if err != nil {
		return "", err
	}

	return resp.User.Id, nil
}
