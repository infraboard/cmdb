package syncer

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/mcube/crypto/cbc"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

var (
	validate = validator.New()
)

const (
	CrendentialAPIKey CrendentialType = iota
	CrendentialPassword
)

type CrendentialType int

func NewDefaultSecret() *Secret {
	return &Secret{
		CreateSecretRequest: NewCreateSecretRequest(),
	}
}

type Secret struct {
	Id       string `json:"id"`        // 全局唯一Id
	CreateAt int64  `json:"create_at"` // 创建时间

	*CreateSecretRequest
}

func (s *Secret) EncryptAPISecret(key string) error {
	// 判断文本是否已经加密
	if strings.HasPrefix(s.APISecret, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text has ciphered")
	}

	cipherText, err := cbc.Encrypt([]byte(s.APISecret), []byte(key))
	if err != nil {
		return err
	}

	base64Str := base64.StdEncoding.EncodeToString(cipherText)
	s.APISecret = fmt.Sprintf("%s%s", conf.CIPHER_TEXT_PREFIX, base64Str)
	return nil
}

func (s *Secret) DecryptAPISecret(key string) error {
	// 判断文本是否已经是明文
	if !strings.HasPrefix(s.APISecret, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text is plan text")
	}

	base64CipherText := strings.TrimPrefix(s.APISecret, conf.CIPHER_TEXT_PREFIX)

	cipherText, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return err
	}

	planText, err := cbc.Decrypt([]byte(cipherText), []byte(key))
	if err != nil {
		return err
	}

	s.APISecret = string(planText)
	return nil
}

func (s *Secret) Desense() {
	if s.APISecret != "" {
		s.APISecret = "******"
	}
}

func (s *Secret) IsAllowRegion(region string) bool {
	for i := range s.AllowRegions {
		if s.AllowRegions[i] == "*" {
			return true
		}

		if s.AllowRegions[i] == region {
			return true
		}
	}

	return false
}

func (s *Secret) AllowRegionString() string {
	return strings.Join(s.AllowRegions, ",")
}

func (s *Secret) LoadAllowRegionFromString(regions string) {
	if regions != "" {
		s.AllowRegions = strings.Split(regions, ",")
	}
}

type CreateSecretRequest struct {
	Description     string          `json:"description" validate:"required,lte=100"` // 描述
	Vendor          resource.Vendor `json:"vendor"`                                  // 厂商
	AllowRegions    []string        `json:"allow_regions"`                           // 允许同步的区域
	CrendentialType CrendentialType `json:"crendential_type"`                        // 凭证类型
	Address         string          `json:"address"`                                 // 服务地址, 云商不用填写
	APIKey          string          `json:"api_key" validate:"required,lte=100"`     // key
	APISecret       string          `json:"api_secret" validate:"required,lte=100"`  // secrete
	RequestRate     int             `json:"request_rate"`                            // 请求速率限制, 默认1秒5个
}

func NewSecretSet() *SecretSet {
	return &SecretSet{
		Items: []*Secret{},
	}
}

type SecretSet struct {
	Items []*Secret `json:"items"`
}

func (s *SecretSet) Add(item *Secret) {
	s.Items = append(s.Items, item)
}

func NewSecret(req *CreateSecretRequest) (*Secret, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Secret{
		Id:                  xid.New().String(),
		CreateAt:            ftime.Now().Timestamp(),
		CreateSecretRequest: req,
	}, nil
}

func NewCreateSecretRequest() *CreateSecretRequest {
	return &CreateSecretRequest{
		RequestRate: 5,
	}
}

func (req *CreateSecretRequest) Validate() error {
	if len(req.AllowRegions) == 0 {
		return fmt.Errorf("required less one allow_regions")
	}
	return validate.Struct(req)
}

func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{}
}

type QuerySecretRequest struct {
}

func NewDescribeSecretRequest(id string) *DescribeSecretRequest {
	return &DescribeSecretRequest{
		Id: id,
	}
}

type DescribeSecretRequest struct {
	Id string
}
