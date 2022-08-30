package connectivity

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	actiontrail "github.com/alibabacloud-go/actiontrail-20200706/v2/client"
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	cms "github.com/alibabacloud-go/cms-20190101/v7/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dds "github.com/alibabacloud-go/dds-20151201/v3/client"
	domain "github.com/alibabacloud-go/domain-20180129/v3/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	redis "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	rds "github.com/alibabacloud-go/rds-20140815/v2/client"
	slb "github.com/alibabacloud-go/slb-20140515/v3/client"
	"github.com/alibabacloud-go/tea/tea"
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
	AccessKey    string `env:"AL_CLOUD_ACCESS_KEY"`
	AccessSecret string `env:"AL_CLOUD_ACCESS_SECRET"`
	Region       string `env:"AL_CLOUD_REGION"`

	ecsConn   *ecs.Client
	rdsConn   *rds.Client
	cmsConn   *cms.Client
	redisConn *redis.Client
	ddsConn   *dds.Client
	domConn   *domain.Client
	dnsConn   *dns.Client
	slbConn   *slb.Client
	atConn    *actiontrail.Client
	bssConn   *bssopenapi.Client
}

// EcsClient 客户端
func (c *AliCloudClient) EcsClient() (*ecs.Client, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client, err := ecs.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("ecs.aliyuncs.com"),
		RegionId:        tea.String(c.Region),
	})
	if err != nil {
		return nil, err
	}

	c.ecsConn = client
	return client, nil
}

func (c *AliCloudClient) CmsClient() (*cms.Client, error) {
	if c.cmsConn != nil {
		return c.cmsConn, nil
	}

	client, err := cms.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("metrics.aliyuncs.com"),
	})
	if err != nil {
		return nil, err
	}

	c.cmsConn = client
	return client, nil
}

func (c *AliCloudClient) BssClient() (*bssopenapi.Client, error) {
	if c.bssConn != nil {
		return c.bssConn, nil
	}

	client, err := bssopenapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("business.aliyuncs.com"),
	})
	if err != nil {
		return nil, err
	}

	c.bssConn = client
	return client, nil
}

func (c *AliCloudClient) DomainClient() (*domain.Client, error) {
	if c.domConn != nil {
		return c.domConn, nil
	}

	client, err := domain.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("domain.aliyuncs.com"),
	})
	if err != nil {
		return nil, err
	}

	c.domConn = client
	return client, nil
}

func (c *AliCloudClient) DnsClient() (*dns.Client, error) {
	if c.dnsConn != nil {
		return c.dnsConn, nil
	}

	client, err := dns.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("alidns.aliyuncs.com"),
	})
	if err != nil {
		return nil, err
	}

	c.dnsConn = client
	return client, nil
}

func (c *AliCloudClient) RdsClient() (*rds.Client, error) {
	if c.rdsConn != nil {
		return c.rdsConn, nil
	}

	client, err := rds.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("rds.aliyuncs.com"),
		RegionId:        tea.String(c.Region),
	})
	if err != nil {
		return nil, fmt.Errorf("new rds client error, %s", err)
	}

	c.rdsConn = client
	return client, nil
}

func (c *AliCloudClient) RedisClient() (*redis.Client, error) {
	if c.redisConn != nil {
		return c.redisConn, nil
	}

	client, err := redis.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("r-kvstore.aliyuncs.com"),
		RegionId:        tea.String(c.Region),
	})
	if err != nil {
		return nil, fmt.Errorf("new rds client error, %s", err)
	}

	c.redisConn = client
	return client, nil
}

func (c *AliCloudClient) MongoDBClient() (*dds.Client, error) {
	if c.ddsConn != nil {
		return c.ddsConn, nil
	}

	client, err := dds.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("mongodb.aliyuncs.com"),
		RegionId:        tea.String(c.Region),
	})
	if err != nil {
		return nil, fmt.Errorf("new rds client error, %s", err)
	}

	c.ddsConn = client
	return client, nil
}

func (c *AliCloudClient) ActionTrailClient() (*actiontrail.Client, error) {
	if c.atConn != nil {
		return c.atConn, nil
	}

	client, err := actiontrail.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		RegionId:        tea.String(c.Region),
		Endpoint:        tea.String("actiontrail." + c.Region + ".aliyuncs.com"),
	})
	if err != nil {
		return nil, err
	}

	c.atConn = client
	return client, nil
}

func (c *AliCloudClient) OssClient() (*oss.Client, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	ep := fmt.Sprintf("https://oss-%s.aliyuncs.com", c.Region)
	return oss.New(ep, c.AccessKey, c.AccessSecret)
}

func (c *AliCloudClient) SLBClient() (*slb.Client, error) {
	if c.slbConn != nil {
		return c.slbConn, nil
	}

	client, err := slb.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("slb.aliyuncs.com"),
		RegionId:        tea.String(c.Region),
	})
	if err != nil {
		return nil, fmt.Errorf("new slb client error, %s", err)
	}

	c.slbConn = client
	return client, nil
}

// 获取客户端账号ID
// 参考: https://next.api.aliyun.com/api/Sts/2015-04-01/GetCallerIdentity
func (c *AliCloudClient) Account() (string, error) {
	args := sts.CreateGetCallerIdentityRequest()

	stsClient, err := sts.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	if err != nil {
		return "", fmt.Errorf("unable to initialize the STS client: %#v", err)
	}

	stsClient.GetConfig().WithScheme("HTTPS")
	stsClient.AppendUserAgent("Infraboard", "1.0")
	identity, err := stsClient.GetCallerIdentity(args)
	if err != nil {
		return "", err
	}
	if identity == nil {
		return "", fmt.Errorf("caller identity not found")
	}

	return identity.AccountId, nil
}
