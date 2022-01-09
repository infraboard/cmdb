package host

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
)

const (
	AppName = "Host"
)

func NewDefaultHost() *Host {
	return &Host{
		Base: &resource.Base{
			ResourceType: resource.Type_HOST,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

func (h *Host) Put(req *UpdateHostData) {
	h.Information = req.Information
	h.Describe = req.Describe
	h.Information.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()
}

func (h *Host) ShortDesc() string {
	return fmt.Sprintf("%s %s", h.Information.Name, h.Information.PrivateIp)
}

func (h *Host) Patch(req *UpdateHostData) error {
	err := ObjectPatch(h.Information, req.Information)
	if err != nil {
		return err
	}

	err = ObjectPatch(h.Describe, req.Describe)
	if err != nil {
		return err
	}

	h.Information.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()
	return nil
}

func ObjectPatch(old, new interface{}) error {
	newByte, err := json.Marshal(new)
	if err != nil {
		return err
	}
	return json.Unmarshal(newByte, old)
}

func (h *Host) GenHash() error {
	hash := sha1.New()

	b, err := json.Marshal(h.Information)
	if err != nil {
		return err
	}
	hash.Write(b)
	h.Base.ResourceHash = fmt.Sprintf("%x", hash.Sum(nil))

	b, err = json.Marshal(h.Describe)
	if err != nil {
		return err
	}
	hash.Reset()
	hash.Write(b)
	h.Base.DescribeHash = fmt.Sprintf("%x", hash.Sum(nil))
	return nil
}

func (d *Describe) KeyPairNameToString() string {
	return strings.Join(d.KeyPairName, ",")
}

func (d *Describe) SecurityGroupsToString() string {
	return strings.Join(d.SecurityGroups, ",")
}

func (d *Describe) LoadKeyPairNameString(s string) {
	if s != "" {
		d.KeyPairName = strings.Split(s, ",")
	}
}

func (d *Describe) LoadSecurityGroupsString(s string) {
	if s != "" {
		d.SecurityGroups = strings.Split(s, ",")
	}
}

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
}

func (s *HostSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (req *DescribeHostRequest) Where() (string, interface{}) {
	switch req.DescribeBy {
	case DescribeBy_HOST_ID:
		return "id = ?", req.Value
	case DescribeBy_INSTANCE_ID:
		return "instance_id = ?", req.Value
	default:
		return "", nil
	}
}

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewQueryHostRequestFromHTTP(r *http.Request) *QueryHostRequest {
	qs := r.URL.Query()
	page := request.NewPageRequestFromHTTP(r)

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
	return &QueryHostRequest{
		Page:     page,
		Keywords: kw,
	}
}

func (req *QueryHostRequest) OffSet() int64 {
	return int64(req.Page.PageSize) * int64(req.Page.PageNumber-1)
}

func NewDescribeHostRequestWithID(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		DescribeBy: DescribeBy_HOST_ID,
		Value:      id,
	}
}

func NewDescribeHostRequestInstanceID(instanceId string) *DescribeHostRequest {
	return &DescribeHostRequest{
		DescribeBy: DescribeBy_INSTANCE_ID,
		Value:      instanceId,
	}
}

func NewDeleteHostRequestWithID(id string) *DeleteHostRequest {
	return &DeleteHostRequest{Id: id}
}

func NewUpdateHostRequest(id string) *UpdateHostRequest {
	return &UpdateHostRequest{
		Id:             id,
		UpdateMode:     UpdateMode_PUT,
		UpdateHostData: &UpdateHostData{},
	}
}

func (req *UpdateHostRequest) Validate() error {
	return validate.Struct(req)
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewHostSet(),
	}
}

type PagerResult struct {
	Data    *HostSet
	Err     error
	HasNext bool
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}
