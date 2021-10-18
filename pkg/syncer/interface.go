package syncer

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/cmdb/pkg/host"
)

var (
	validate = validator.New()
)

type Service interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error)
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

type CreateSecretRequest struct {
	Description     string          `json:"description" validate:"required,lte=100"` // 描述
	Vendor          host.Vendor     `json:"vendor"`                                  // 厂商
	Region          string          `json:"region" validate:"required,lte=100"`      // 区域
	CrendentialType CrendentialType `json:"crendential_type"`                        // 凭证类型
	APIKey          string          `json:"api_key" validate:"required,lte=100"`     // key
	APISecret       string          `json:"api_secret" validate:"required,lte=100"`  // secrete
}

func (req *CreateSecretRequest) Validate() error {
	return validate.Struct(req)
}

func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{}
}

type QuerySecretRequest struct {
}
