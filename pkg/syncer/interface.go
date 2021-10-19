package syncer

import (
	"context"
)

type Service interface {
	SecretService
	SyncService
}

type SecretService interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error)
	DescribeSecret(context.Context, *DescribeSecretRequest) (*Secret, error)
}

type SyncService interface {
	Sync(context.Context, *SyncRequest) (*SyncReponse, error)
}
