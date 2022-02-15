package resource

import (
	"net/http"
	"strings"

	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/http/request"
)

const (
	AppName = "Resource"
)

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewSearchRequestFromHTTP(r *http.Request) *SearchRequest {
	qs := r.URL.Query()

	return &SearchRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

func NewDefaultResource() *Resource {
	return &Resource{
		Base:        &Base{},
		Information: &Information{},
	}
}

func (i *Information) PrivateIPToString() string {
	return strings.Join(i.PrivateIp, ",")
}

func (i *Information) PublicIPToString() string {
	return strings.Join(i.PublicIp, ",")
}

func (i *Information) LoadPrivateIPString(s string) {
	if s != "" {
		i.PrivateIp = strings.Split(s, ",")
	}
}

func (i *Information) LoadPublicIPString(s string) {
	if s != "" {
		i.PublicIp = strings.Split(s, ",")
	}
}

func (i *Information) Hash() string {
	return utils.Hash(i)
}

func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Items: []*Resource{},
	}
}

func (r *ResourceSet) Add(item *Resource) {
	r.Items = append(r.Items, item)
}

type AccountGetter struct {
	accountId string
}

func (ag *AccountGetter) WithAccountId(id string) {
	ag.accountId = id
}

func (ag *AccountGetter) GetAccountId() string {
	return ag.accountId
}
