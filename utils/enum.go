package utils

type EnumDescribe struct {
	Name     string          `json:"name"`
	Value    string          `json:"value"`
	Describe string          `json:"describe"`
	SubItems []*EnumDescribe `json:"sub_items"`
}
