package connectivity

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	cms "github.com/alibabacloud-go/cms-20190101/v7/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	redis "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	rds "github.com/alibabacloud-go/rds-20140815/v2/client"
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

func (c *AliCloudClient) OssClient() (*oss.Client, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	ep := fmt.Sprintf("https://oss-%s.aliyuncs.com", c.Region)
	return oss.New(ep, c.AccessKey, c.AccessSecret)
}

// 获取客户端账号ID
// 参考: https://next.api.aliyun.com/api/Sts/2015-04-01/GetCallerIdentity
func (c *AliCloudClient) Account() (string, error) {
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

	return identity.AccountId, nil
}
