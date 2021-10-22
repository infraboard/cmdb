package resource

import (
	"bytes"
	"fmt"
	"strings"
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "Unsuport",
		1: "Host_Resource",
		2: "Rds_Resource",
	}
	Type_value = map[string]int32{
		"Unsuport":      0,
		"Host_Resource": 1,
		"Rds_Resource":  2,
	}
)

// ParseTypeFromString Parse Type from string
func ParseTypeFromString(str string) (Type, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Type_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Type: %s", str)
	}

	return Type(v), nil
}

func (t Type) String() string {
	return Type_name[int32(t)]
}

// Equal type compare
func (t Type) Equal(target Type) bool {
	return t == target
}

// IsIn todo
func (t Type) IsIn(targets ...Type) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t Type) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Type) UnmarshalJSON(b []byte) error {
	ins, err := ParseTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// Enum value maps for Type.
var (
	Vendor_name = map[int32]string{
		0: "VENDOR_ALIYUN",
		1: "VENDOR_TENCENT",
		2: "VENDOR_HUAWEI",
		3: "VENDOR_VSPHERE",
	}
	Vendor_value = map[string]int32{
		"VENDOR_ALIYUN":  0,
		"VENDOR_TENCENT": 1,
		"VENDOR_HUAWEI":  2,
		"VENDOR_VSPHERE": 3,
	}
)

// ParseVendorFromString Parse Type from string
func ParseVendorFromString(str string) (Vendor, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Vendor_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Vendor: %s", str)
	}

	return Vendor(v), nil
}

func (t Vendor) String() string {
	return Vendor_name[int32(t)]
}

// Equal type compare
func (t Vendor) Equal(target Vendor) bool {
	return t == target
}

// IsIn todo
func (t Vendor) IsIn(targets ...Vendor) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t Vendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Vendor) UnmarshalJSON(b []byte) error {
	ins, err := ParseVendorFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
