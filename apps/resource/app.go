package resource

import (
	"fmt"
	"net/http"
	"sort"
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

func NewSearchRequestFromHTTP(r *http.Request) (*SearchRequest, error) {
	qs := r.URL.Query()
	req := &SearchRequest{
		Page:        request.NewPageRequestFromHTTP(r),
		Keywords:    qs.Get("keywords"),
		ExactMatch:  qs.Get("exact_match") == "true",
		Domain:      qs.Get("domain"),
		Namespace:   qs.Get("namespace"),
		Env:         qs.Get("env"),
		Status:      qs.Get("status"),
		SyncAccount: qs.Get("sync_account"),
		WithTags:    qs.Get("with_tags") == "true",
	}

	umStr := qs.Get("usage_mode")
	if umStr != "" {
		mode, err := ParseUsageModeFromString(umStr)
		if err != nil {
			return nil, err
		}
		req.UsageMode = &mode
	}

	rtStr := qs.Get("resource_type")
	if rtStr != "" {
		rt, err := ParseTypeFromString(rtStr)
		if err != nil {
			return nil, err
		}
		req.Type = &rt
	}

	tgStr := qs.Get("tag")
	if tgStr != "" {
		tg, err := NewTagsFromString(tgStr)
		if err != nil {
			return nil, err
		}
		req.Tags = tg
	}

	return req, nil
}

func (req *SearchRequest) GroupByKey() map[string][]*Tag {
	gt := map[string][]*Tag{}
	for i := range req.Tags {
		t := req.Tags[i]
		if _, ok := gt[t.Key]; !ok {
			gt[t.Key] = []*Tag{}
		}

		gt[t.Key] = append(gt[t.Key], t)
	}

	return gt
}

func (req *SearchRequest) HasTag() bool {
	if req.Tags == nil {
		return false
	}
	return len(req.Tags) > 0
}

func NewDefaultResource() *Resource {
	return &Resource{
		Base:        &Base{},
		Information: &Information{},
	}
}

func (r *Information) AddTag(t *Tag) {
	r.Tags = append(r.Tags, t)
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

func (i *Information) SortTag() {
	sort.Slice(i.Tags, func(m, n int) bool {
		return i.Tags[m].Weight < i.Tags[n].Weight
	})
}

func (i *Information) Hash() string {
	return utils.Hash(i)
}

func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Items: []*Resource{},
	}
}

func (s *ResourceSet) Add(item *Resource) {
	s.Items = append(s.Items, item)
}

func (s *ResourceSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Base.Id)
	}

	return
}

func (s *ResourceSet) UpdateTag(tags []*Tag) {
	for i := range tags {
		for j := range s.Items {
			if s.Items[j].Base.Id == tags[i].ResourceId {
				s.Items[j].Information.AddTag(tags[i])
			}
		}
	}
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

func NewDefaultTag() *Tag {
	return &Tag{
		Type:   TagType_USER,
		Weight: 1,
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

// key=value,key=value
func NewTagsFromString(tagStr string) (tags []*Tag, err error) {
	if tagStr == "" {
		return
	}

	items := strings.Split(tagStr, ",")
	for _, v := range items {
		kv := strings.Split(v, "=")
		if len(kv) != 2 {
			err = fmt.Errorf("key,value format error, requred key=value")
			return
		}
		tags = append(tags, &Tag{
			Key:   kv[0],
			Value: kv[1],
		})
	}

	return
}

func NewTagSet() *TagSet {
	return &TagSet{
		Items: []*Tag{},
	}
}
