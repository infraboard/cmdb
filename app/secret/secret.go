package secret

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/crypto/cbc"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

const (
	AppName = "Secret"
)

var (
	validate = validator.New()
)

func NewDefaultSecret() *Secret {
	return &Secret{
		RequestRate: 5,
	}
}

func (s *Secret) EncryptAPISecret(key string) error {
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

func (s *Secret) DecryptAPISecret(key string) error {
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

func (s *Secret) Desense() {
	if s.ApiSecret != "" {
		s.ApiSecret = "******"
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

func (s *Secret) ShortDesc() string {
	return fmt.Sprintf("%s[%s]", s.Description, s.DensenseKey())
}

func (s *Secret) DensenseKey() string {
	if s.ApiKey == "" {
		return ""
	}
	total := len(s.ApiKey)
	if total > 8 {
		return fmt.Sprintf("%s****%s", s.ApiKey[:4], s.ApiKey[total-4:])
	}
	return s.ApiKey
}

func (s *Secret) AllowRegionString() string {
	return strings.Join(s.AllowRegions, ",")
}

func (s *Secret) LoadAllowRegionFromString(regions string) {
	if regions != "" {
		s.AllowRegions = strings.Split(regions, ",")
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

func NewSecret(req *CreateSecretRequest) (*Secret, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Secret{
		Id:              xid.New().String(),
		CreateAt:        ftime.Now().Timestamp(),
		Description:     req.Description,
		Vendor:          req.Vendor,
		AllowRegions:    req.AllowRegions,
		CrendentialType: req.CrendentialType,
		Address:         req.Address,
		ApiKey:          req.ApiKey,
		ApiSecret:       req.ApiSecret,
		RequestRate:     req.RequestRate,
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

func NewQuerySecretRequestFromHTTP(r *http.Request) *QuerySecretRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	kw := qs.Get("keywords")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}
	return &QuerySecretRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{
		PageSize:   20,
		PageNumber: 1,
		Keywords:   "",
	}
}

func (req *QuerySecretRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
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
