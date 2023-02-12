package resource

import (
	context "context"
	"fmt"
	"hash/fnv"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

var (
	validate = validator.New()
)

const (
	AppName = "resource"
)

type Service interface {
	Put(context.Context, *Resource) (*Resource, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	RPCServer
}

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewSearchRequestFromHTTP(r *http.Request) (*SearchRequest, error) {
	qs := r.URL.Query()
	req := &SearchRequest{
		Page:      request.NewPageRequestFromHTTP(r),
		Keywords:  qs.Get("keywords"),
		Domain:    qs.Get("domain"),
		Namespace: qs.Get("namespace"),
		Env:       qs.Get("env"),
		Status:    qs.Get("status"),
		Owner:     qs.Get("owner"),
		WithTags:  qs.Get("with_tags") == "true",
		Tags:      []*TagSelector{},
	}

	umStr := qs.Get("usage_mode")
	if umStr != "" {
		mode, err := ParseUSAGE_MODEFromString(umStr)
		if err != nil {
			return nil, err
		}
		req.UsageMode = &mode
	}

	rtStr := qs.Get("resource_type")
	if rtStr != "" {
		rt, err := ParseTYPEFromString(rtStr)
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
		req.AddTag(tg...)
	}

	return req, nil
}

func (req *SearchRequest) HasTag() bool {
	if req.Tags == nil {
		return false
	}
	return len(req.Tags) > 0
}

func (req *SearchRequest) AddTag(t ...*TagSelector) {
	req.Tags = append(req.Tags, t...)
}

func NewDefaultResource(rt TYPE) *Resource {
	return &Resource{
		Meta: &Meta{},
		Spec: &Spec{
			ResourceType: rt,
			Tags:         []*Tag{},
		},
		Cost:   &Cost{},
		Status: &Status{},
	}
}

func (r *Resource) AddTag(t *Tag) {
	r.Spec.Tags = append(r.Spec.Tags, t)
}

func (i *Status) PrivateIPToString() string {
	return strings.Join(i.PrivateAddress, ",")
}

func (i *Status) PublicIPToString() string {
	return strings.Join(i.PublicAddress, ",")
}

func (i *Status) LoadPrivateIPString(s string) {
	if s != "" {
		i.PrivateAddress = strings.Split(s, ",")
	}
}

func (i *Status) LoadPublicIPString(s string) {
	if s != "" {
		i.PublicAddress = strings.Split(s, ",")
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

func NewDefaultTag() *Tag {
	return &Tag{
		Purpose: TAG_PURPOSE_GROUP,
		Weight:  1,
	}
}

func NewGroupTag(key, value string) *Tag {
	return &Tag{
		Purpose: TAG_PURPOSE_GROUP,
		Key:     key,
		Value:   value,
		Weight:  1,
	}
}

type Operator string

const (
	Operator_EQUAL          = "="
	Operator_NOT_EQUAL      = "!="
	Operator_LIKE_EQUAL     = "=~"
	Operator_NOT_LIKE_EQUAL = "!~"
)

func ParExpr(str string) (*TagSelector, error) {
	var (
		op = ""
		kv = []string{}
	)

	if strings.Contains(str, Operator_LIKE_EQUAL) {
		op = "LIKE"
		kv = strings.Split(str, Operator_LIKE_EQUAL)
	} else if strings.Contains(str, Operator_NOT_LIKE_EQUAL) {
		op = "NOT LIKE"
		kv = strings.Split(str, Operator_NOT_LIKE_EQUAL)
	} else if strings.Contains(str, Operator_NOT_EQUAL) {
		op = "!="
		kv = strings.Split(str, Operator_NOT_EQUAL)
	} else if strings.Contains(str, Operator_EQUAL) {
		op = "="
		kv = strings.Split(str, Operator_EQUAL)
	} else {
		return nil, fmt.Errorf("no support operator [=, =~, !=, !~]")
	}

	if len(kv) != 2 {
		return nil, fmt.Errorf("key,value format error, requred key=value")
	}

	selector := &TagSelector{
		Key:     kv[0],
		Opertor: op,
		Values:  []string{},
	}

	// 如果Value等于*表示只匹配key
	if kv[1] != "*" {
		selector.Values = strings.Split(kv[1], ",")
	}

	return selector, nil
}

// key1=v1,v2,v3&key2=~v1,v2,v3
func NewTagsFromString(tagStr string) (tags []*TagSelector, err error) {
	if tagStr == "" {
		return
	}

	items := strings.Split(tagStr, "&")
	for _, v := range items {
		t, err := ParExpr(v)
		if err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	return
}

func (s *TagSelector) RelationShip() string {
	switch s.Opertor {
	case Operator_EQUAL, Operator_LIKE_EQUAL:
		return " OR "
	case Operator_NOT_EQUAL, Operator_NOT_LIKE_EQUAL:
		return " AND "
	default:
		return " OR "
	}
}

func NewTagSet() *TagSet {
	return &TagSet{
		Items: []*Tag{},
	}
}

func (r *Resource) Validate() error {
	return validate.Struct(r)
}

func (r *DeleteRequest) Validate() error {
	return nil
}

// GenResourceHashId 生成资源的短hashId
func GenResourceHashId(t TYPE, unitkeys ...string) string {
	hash := fnv.New64a()
	for _, key := range unitkeys {
		hash.Write([]byte(key))
	}
	return fmt.Sprintf("%s-%x", strings.ToLower(t.String()), hash.Sum64())
}
