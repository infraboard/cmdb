package client

// NewDefaultConfig todo
func NewConfig(address string) *Config {
	return &Config{
		address:        address,
		Authentication: &Authentication{},
	}
}

// Config 客户端配置
type Config struct {
	address   string
	enableTLS bool
	*Authentication
}

// SetAddress todo
func (c *Config) SetAddress(addr string) {
	c.address = addr
}

// Address 地址
func (c *Config) Address() string {
	return c.address
}

// Option configures how we set up the grpc.
type Option interface {
	apply(*Config)
}

func newFuncOption(f func(*Config)) Option {
	return &funcOption{
		f: f,
	}
}

type funcOption struct {
	f func(*Config)
}

func (fdo *funcOption) apply(do *Config) {
	fdo.f(do)
}
func WithTLS() Option {
	return newFuncOption(func(o *Config) {
		o.enableTLS = true
	})
}
