package utils_test

import (
	"testing"

	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger/zap"
)

func TestParseTime(t *testing.T) {
	utils.ParseDefaultSecondTime("")
}

func init() {
	zap.DevelopmentSetup()
}
