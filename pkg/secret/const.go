package secret

import (
	"bytes"
	"fmt"
	"strings"
)

// Enum value maps for Type.
var (
	CrendentialType_name = map[int32]string{
		0: "CRENDENTIAL_API_KEY",
		1: "CRENDENTIAL_PASSWORD",
	}
	CrendentialType_value = map[string]int32{
		"CRENDENTIAL_API_KEY":  0,
		"CRENDENTIAL_PASSWORD": 1,
	}
)

// ParseCrendentialTypeFromString Parse Type from string
func ParseCrendentialTypeFromString(str string) (CrendentialType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := CrendentialType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Type: %s", str)
	}

	return CrendentialType(v), nil
}

func (t CrendentialType) String() string {
	return CrendentialType_name[int32(t)]
}

// Equal type compare
func (t CrendentialType) Equal(target CrendentialType) bool {
	return t == target
}

// IsIn todo
func (t CrendentialType) IsIn(targets ...CrendentialType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t CrendentialType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *CrendentialType) UnmarshalJSON(b []byte) error {
	ins, err := ParseCrendentialTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
