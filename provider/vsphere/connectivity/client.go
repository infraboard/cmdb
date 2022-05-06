package connectivity

import (
	"context"
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/vim25"
)

var (
	client *VsphereClient
)

func C() *VsphereClient {
	if client == nil {
		panic("please load config first")
	}
	return client
}

func LoadClientFromEnv() error {
	client = &VsphereClient{}
	if err := env.Parse(client); err != nil {
		return err
	}

	return nil
}

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewVsphereClient client
func NewVsphereClient(host, username, password string) *VsphereClient {
	return &VsphereClient{
		Host:     host,
		Username: username,
		Password: password,
		Insecure: true,
	}
}

type VsphereClient struct {
	Host     string `validate:"required" env:"VS_HOST"`
	Username string `validate:"required" env:"VS_USERNAME"`
	Password string `validate:"required" env:"VS_PASSWORD"`
	Insecure bool

	client *vim25.Client
}

// 参考官方样例: https://github.com/vmware/govmomi/blob/master/examples/examples.go#L104
func (v *VsphereClient) VimClient() (*vim25.Client, error) {
	if v.client != nil {
		return v.client, nil
	}

	// 校验参数
	if err := v.validate(); err != nil {
		return nil, err
	}

	// 构建soap client
	u, err := v.URL()
	if err != nil {
		return nil, err
	}

	// Share govc's session cache
	s := &cache.Session{
		URL:      u,
		Insecure: v.Insecure,
	}

	v.client = new(vim25.Client)
	err = s.Login(context.Background(), v.client, nil)
	if err != nil {
		return nil, err
	}

	return v.client, nil
}

func (v *VsphereClient) validate() error {
	return validate.Struct(v)
}

func (v *VsphereClient) URL() (*url.URL, error) {
	tempURL := fmt.Sprintf("https://%s:%s@%s", v.Username, url.PathEscape(v.Password), v.Host)
	return url.Parse(tempURL + "/sdk")
}
