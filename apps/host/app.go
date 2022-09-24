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
	h.Resource.Spec = req.Spec
	h.Describe = req.Describe
	h.Resource.Spec.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()
}

func (h *Host) ShortDesc() string {
	return fmt.Sprintf("%s %s", h.Resource.Spec.Name, h.Resource.Status.PrivateIp)
}

func NewUpdateHostDataByIns(ins *Host) *UpdateHostData {
	return &UpdateHostData{
		Spec:     ins.Resource.Spec,
		Describe: ins.Describe,
	}
}

func (h *Host) Patch(req *UpdateHostData) error {
	err := ObjectPatch(h.Resource.Spec, req.Spec)
	if err != nil {
		return err
	}

	err = ObjectPatch(h.Describe, req.Describe)
	if err != nil {
		return err
	}

	h.Resource.Spec.UpdateAt = ftime.Now().Timestamp()
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
	h.Resource.Meta.ResourceHash = h.Resource.Spec.Hash()
	h.Resource.Meta.DescribeHash = utils.Hash(h.Describe)
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
		ids = append(ids, s.Items[i].Resource.Meta.Id)
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

// func (s *HostSet) UpdateTag(tags []*resource.Tag) {
// 	for i := range tags {
// 		for j := range s.Items {
// 			if s.Items[j].Resource.Meta.Id == tags[i].ResourceId {
// 				s.Items[j].Resource.AddTag(tags[i])
// 			}
// 		}
// 	}
// }

func (h *Host) Status() STATUS {
	s, err := ParseSTATUSFromString(h.Resource.Status.Phase)
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
