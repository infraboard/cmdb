package syncer

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
