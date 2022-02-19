package utils

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcube/logger/zap"
)

func Hash(x interface{}) string {
	hash := sha1.New()
	b, err := json.Marshal(x)
	if err != nil {
		zap.L().Errorf("hash %v error, %s", x, err)
		return ""
	}
	hash.Write(b)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
