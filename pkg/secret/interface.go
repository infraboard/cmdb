package secret

import (
	"context"
)

type Service interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error)
	DescribeSecret(context.Context, *DescribeSecretRequest) (*Secret, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*Secret, error)
}
