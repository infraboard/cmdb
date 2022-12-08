package impl

import "github.com/infraboard/cmdb/apps/secret"

func NewSecret(s *secret.Secret) *Secret {
	return &Secret{
		s.Meta,
		s.Spec,
	}
}

type Secret struct {
	*secret.Meta
	*secret.CreateSecretRequest
}

func (s *Secret) TableName() string {
	return "secrets"
}

func (s *Secret) Secret() *secret.Secret {
	return &secret.Secret{
		Meta: s.Meta,
		Spec: s.CreateSecretRequest,
	}
}

func NewSecretSet() *SecretSet {
	return &SecretSet{
		Items: []*Secret{},
	}
}

type SecretSet struct {
	Total int64
	Items []*Secret
}

func (s *SecretSet) SecretSet() *secret.SecretSet {
	set := secret.NewSecretSet()
	set.Total = s.Total
	for i := range s.Items {
		item := s.Items[i].Secret()
		item.Spec.Desense()
		set.Add(item)
	}
	return set
}
