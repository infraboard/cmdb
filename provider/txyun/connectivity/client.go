package connectivity

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/cmdb/utils"
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	redis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
)

var (
	client *TencentCloudClient
)

func C() *TencentCloudClient {
	if client == nil {
		panic("please load config first")
	}
	return client
}

func LoadClientFromEnv() error {
	client = &TencentCloudClient{}
	if err := env.Parse(client); err != nil {
		return err
	}

	return nil
}

// NewTencentCloudClient client
func NewTencentCloudClient(credentialID, credentialKey, region string) *TencentCloudClient {
	return &TencentCloudClient{
		Region:    region,
		SecretID:  credentialID,
		SecretKey: credentialKey,
	}
}

// TencentCloudClient client for all TencentCloud service
type TencentCloudClient struct {
	Region    string `env:"TX_CLOUD_REGION"`
	SecretID  string `env:"TX_CLOUD_SECRET_ID"`
	SecretKey string `env:"TX_CLOUD_SECRET_KEY"`

	cvmConn   *cvm.Client
	cdbConn   *cdb.Client
	redisConn *redis.Client
	billConn  *billing.Client
}

// UseCvmClient cvm
func (me *TencentCloudClient) CvmClient() *cvm.Client {
	if me.cvmConn != nil {
		return me.cvmConn
	}

	credential := common.NewCredential(
		me.SecretID,
		me.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	cvmConn, _ := cvm.NewClient(credential, me.Region, cpf)
	me.cvmConn = cvmConn
	return me.cvmConn
}

// UseBillingClient billing客户端
func (me *TencentCloudClient) BillingClient() *billing.Client {
	if me.billConn != nil {
		return me.billConn
	}
	credential := common.NewCredential(
		me.SecretID,
		me.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	billConn, _ := billing.NewClient(credential, me.Region, cpf)
	me.billConn = billConn

	return me.billConn
}

// CDBClient cdb
func (me *TencentCloudClient) CDBClient() *cdb.Client {
	if me.cdbConn != nil {
		return me.cdbConn
	}

	credential := common.NewCredential(
		me.SecretID,
		me.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	cdbConn, _ := cdb.NewClient(credential, me.Region, cpf)
	me.cdbConn = cdbConn
	return me.cdbConn
}

// RedisClient cdb
func (me *TencentCloudClient) RedisClient() *redis.Client {
	if me.redisConn != nil {
		return me.redisConn
	}

	credential := common.NewCredential(
		me.SecretID,
		me.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	conn, _ := redis.NewClient(credential, me.Region, cpf)
	me.redisConn = conn
	return me.redisConn
}

// 获取客户端账号ID
func (me *TencentCloudClient) Account() (string, error) {
	credential := common.NewCredential(
		me.SecretID,
		me.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	stsConn, _ := sts.NewClient(credential, me.Region, cpf)

	req := sts.NewGetCallerIdentityRequest()

	resp, err := stsConn.GetCallerIdentity(req)
	if err != nil {
		return "", fmt.Errorf("unable to initialize the STS client: %#v", err)
	}

	return utils.PtrStrV(resp.Response.AccountId), nil
}
