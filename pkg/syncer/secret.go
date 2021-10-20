package syncer

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/cmdb/pkg/resource"
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

type CreateSecretRequest struct {
	Description     string          `json:"description" validate:"required,lte=100"` // 描述
	Vendor          resource.Vendor `json:"vendor"`                                  // 厂商
	Region          string          `json:"region" validate:"required,lte=100"`      // 区域
	CrendentialType CrendentialType `json:"crendential_type"`                        // 凭证类型
	APIKey          string          `json:"api_key" validate:"required,lte=100"`     // key
	APISecret       string          `json:"api_secret" validate:"required,lte=100"`  // secrete
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
	return &CreateSecretRequest{}
}

func (req *CreateSecretRequest) Validate() error {
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
