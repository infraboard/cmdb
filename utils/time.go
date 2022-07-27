package utils

import (
	"strings"
	"time"

	"github.com/infraboard/mcube/logger/zap"
)

const (
	DEFAULT_TIME_SECOND_FORMAT = "2006-01-02T15:04:05Z"
	DEFAULT_TIME_MINITE_FORMAT = "2006-01-02T15:04Z"
)

func ParseDefaultSecondTime(t string) int64 {
	return ParseTime(DEFAULT_TIME_SECOND_FORMAT, t)
}

func ParseDefaultMiniteTime(t string) int64 {
	return ParseTime(DEFAULT_TIME_MINITE_FORMAT, t)
}

func ParseTime(format, t string) int64 {
	t = strings.TrimSpace(t)
	if t == "" {
		return 0
	}

	ts, err := time.Parse(format, t)
	if err != nil {
		zap.L().Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.Unix()
}
