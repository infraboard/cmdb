package resource

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/http/request"
)

var (
	validate = validator.New()
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

func (i *Information) LoadTags(keys, values, describes, weights, types string) error {
	if keys == "" {
		return nil
	}

	kl := strings.Split(keys, ",")
	vl := strings.Split(values, ",")
	dl := strings.Split(describes, ",")
	wl := strings.Split(weights, ",")
	tl := strings.Split(types, ",")

	if len(kl) != len(vl) || len(kl) != len(dl) {
		return fmt.Errorf("len is not equal")
	}

	for idx := range kl {
		t := &Tag{
			Key:      kl[idx],
			Value:    vl[idx],
			Describe: dl[idx],
		}
		t.Weight, _ = strconv.ParseInt(wl[idx], 10, 64)
		tti, _ := strconv.ParseInt(tl[idx], 10, 64)
		t.Type = TagType(int32(tti))
		i.Tags = append(i.Tags, t)
	}

	return nil
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

func NewUpdateTagRequest(resourceId string, action UpdateAction) *UpdateTagRequest {
	return &UpdateTagRequest{
		Id:     resourceId,
		Action: action,
	}
}

func NewThirdTag(key, value string) *Tag {
	return &Tag{
		Type:   TagType_THIRD,
		Key:    key,
		Value:  value,
		Weight: 1,
	}
}

func (req *UpdateTagRequest) Validate() error {
	if len(req.Tags) == 0 {
		return fmt.Errorf("no tags")
	}

	return validate.Struct(req)
}
