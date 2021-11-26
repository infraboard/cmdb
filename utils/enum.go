package utils

import "fmt"

func SkipRegion(val bool) map[string]string {
	return map[string]string{"skip_region": fmt.Sprintf("%t", val)}
}

type EnumDescribe struct {
	Name     string            `json:"name"`
	Value    string            `json:"value"`
	Describe string            `json:"describe"`
	Meta     map[string]string `json:"meta"`
	SubItems []*EnumDescribe   `json:"sub_items"`
}
