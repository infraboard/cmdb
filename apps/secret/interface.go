package secret

import (
	context "context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/crypto/cbc"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
)

const (
	AppName = "secrets"
)

var (
	validate = validator.New()
)

type Service interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*Secret, error)
	RPCServer
}

func NewDefaultSecret() *Secret {
	return &Secret{
		Meta: &Meta{},
		Spec: &CreateSecretRequest{
			RequestRate: 5,
		},
	}
}

func (s *CreateSecretRequest) EncryptAPISecret(key string) error {
	// 判断文本是否已经加密
	if strings.HasPrefix(s.ApiSecret, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text has ciphered")
	}

	cipherText, err := cbc.Encrypt([]byte(s.ApiSecret), []byte(key))
	if err != nil {
		return err
	}

	base64Str := base64.StdEncoding.EncodeToString(cipherText)
	s.ApiSecret = fmt.Sprintf("%s%s", conf.CIPHER_TEXT_PREFIX, base64Str)
	return nil
}

func (s *CreateSecretRequest) DecryptAPISecret(key string) error {
	// 判断文本是否已经是明文
	if !strings.HasPrefix(s.ApiSecret, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text is plan text")
	}

	base64CipherText := strings.TrimPrefix(s.ApiSecret, conf.CIPHER_TEXT_PREFIX)

	cipherText, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return err
	}

	planText, err := cbc.Decrypt([]byte(cipherText), []byte(key))
	if err != nil {
		return err
	}

	s.ApiSecret = string(planText)
	return nil
}

func (s *CreateSecretRequest) Desense() {
	if s.ApiSecret != "" {
		s.ApiSecret = "******"
	}
}

func (s *CreateSecretRequest) IsAllowRegion(region string) bool {
	for i := range s.Regions {
		if s.Regions[i] == "*" {
			return true
		}

		if s.Regions[i] == region {
			return true
		}
	}

	return false
}

func (s *CreateSecretRequest) ShortDesc() string {
	return fmt.Sprintf("%s[%s]", s.Description, s.DensenseKey())
}

func (s *CreateSecretRequest) DensenseKey() string {
	if s.ApiKey == "" {
		return ""
	}
	total := len(s.ApiKey)
	if total > 8 {
		return fmt.Sprintf("%s****%s", s.ApiKey[:4], s.ApiKey[total-4:])
	}
	return s.ApiKey
}

func (s *CreateSecretRequest) AllowRegionString() string {
	return strings.Join(s.Regions, ",")
}

func (s *CreateSecretRequest) LoadAllowRegionFromString(regions string) {
	if regions != "" {
		s.Regions = strings.Split(regions, ",")
	}
}

func NewSecretSet() *SecretSet {
	return &SecretSet{
		Items: []*Secret{},
	}
}

func (s *SecretSet) Add(item *Secret) {
	s.Items = append(s.Items, item)
}

func (s *SecretSet) Desense() {
	for i := range s.Items {
		s.Items[i].Spec.Desense()
	}
}

func NewSecret(req *CreateSecretRequest) (*Secret, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Secret{
		Meta: NewMeta(),
		Spec: req,
	}, nil
}

func NewMeta() *Meta {
	return &Meta{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
	}
}

func NewCreateSecretRequest() *CreateSecretRequest {
	return &CreateSecretRequest{
		RequestRate: 5,
	}
}

func (req *CreateSecretRequest) SetOwner(tk *token.Token) {
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace
}

func (req *CreateSecretRequest) Validate() error {
	if len(req.Regions) == 0 {
		return fmt.Errorf("required less one allow_regions")
	}
	return validate.Struct(req)
}

func NewQuerySecretRequestFromHTTP(r *http.Request) *QuerySecretRequest {
	qs := r.URL.Query()

	return &QuerySecretRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{
		Page:     request.NewDefaultPageRequest(),
		Keywords: "",
	}
}

func (req *QuerySecretRequest) WithNamespace(tk *token.Token) {
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace
}

func NewDescribeSecretRequest(id string) *DescribeSecretRequest {
	return &DescribeSecretRequest{
		Id: id,
	}
}

func NewDeleteSecretRequestWithID(id string) *DeleteSecretRequest {
	return &DeleteSecretRequest{
		Id: id,
	}
}
