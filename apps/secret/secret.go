package secret

import "encoding/json"

func (s *Secret) TableName() string {
	return "secrets"
}

func (c *Secret) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*Meta
		*CreateSecretRequest
	}{c.Meta, c.Spec})
}
