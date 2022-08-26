package host

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/infraboard/mcube/types/ftime"
)

const (
	AppName = "host"
)

func NewDefaultHost() *Host {
	return &Host{
		Resource: resource.NewDefaultResource(resource.TYPE_HOST),
		Describe: &Describe{},
	}
}

func (h *Host) Put(req *UpdateHostData) {
	oldRH, oldDH := h.Resource.Base.ResourceHash, h.Resource.Base.DescribeHash

	h.Resource.Information = req.Information
	h.Describe = req.Describe
	h.Resource.Information.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()

	if h.Resource.Base.ResourceHash != oldRH {
		h.Resource.Base.ResourceHashChanged = true
	}
	if h.Resource.Base.DescribeHash != oldDH {
		h.Resource.Base.DescribeHashChanged = true
	}
}

func (h *Host) ShortDesc() string {
	return fmt.Sprintf("%s %s", h.Resource.Information.Name, h.Resource.Information.PrivateIp)
}

func NewUpdateHostDataByIns(ins *Host) *UpdateHostData {
	return &UpdateHostData{
		Information: ins.Resource.Information,
		Describe:    ins.Describe,
	}
}

func (h *Host) Patch(req *UpdateHostData) error {
	oldRH, oldDH := h.Resource.Base.ResourceHash, h.Resource.Base.DescribeHash

	err := ObjectPatch(h.Resource.Information, req.Information)
	if err != nil {
		return err
	}

	err = ObjectPatch(h.Describe, req.Describe)
	if err != nil {
		return err
	}

	h.Resource.Information.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()

	if h.Resource.Base.ResourceHash != oldRH {
		h.Resource.Base.ResourceHashChanged = true
	}
	if h.Resource.Base.DescribeHash != oldDH {
		h.Resource.Base.DescribeHashChanged = true
	}

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
	h.Resource.Base.ResourceHash = h.Resource.Information.Hash()
	h.Resource.Base.DescribeHash = utils.Hash(h.Describe)
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

func (s *HostSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Host))
	}
}

func (s *HostSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Resource.Base.Id)
	}
	return
}

func (s *HostSet) Length() int64 {
	return int64(len(s.Items))
}

func (s *HostSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *HostSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s *HostSet) UpdateTag(tags []*resource.Tag) {
	for i := range tags {
		for j := range s.Items {
			if s.Items[j].Resource.Base.Id == tags[i].ResourceId {
				s.Items[j].Resource.AddTag(tags[i])
			}
		}
	}
}

func (h *Host) Status() STATUS {
	s, err := ParseSTATUSFromString(h.Resource.Information.Status)
	if err != nil {
		return STATUS_UNKNOW
	}

	return s
}

func (req *DescribeHostRequest) Where() (string, interface{}) {
	switch req.DescribeBy {
	default:
		return "r.id = ?", req.Value
	}
}

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewQueryHostRequestFromHTTP(r *http.Request) *QueryHostRequest {
	qs := r.URL.Query()
	page := request.NewPageRequestFromHTTP(r)
	kw := qs.Get("keywords")

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

func NewDeleteHostRequestWithID(id string) *ReleaseHostRequest {
	return &ReleaseHostRequest{Id: id}
}

func NewUpdateHostRequest(id string) *UpdateHostRequest {
	return &UpdateHostRequest{
		Id:             id,
		UpdateMode:     pb_request.UpdateMode_PUT,
		UpdateHostData: &UpdateHostData{},
	}
}

func (req *UpdateHostRequest) Validate() error {
	return validate.Struct(req)
}
