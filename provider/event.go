package provider

import (
	"time"

	"github.com/infraboard/mcube/pager"
)

type EventOperator interface {
	PageQueryEvent(req *QueryEventRequest) pager.Pager
}

func NewQueryEventRequest() *QueryEventRequest {
	now := time.Now()
	return &QueryEventRequest{
		Rate:      5,
		StartTime: now.Add(-1 * time.Hour),
		EndTime:   now,
	}
}

type QueryEventRequest struct {
	Rate      float64
	StartTime time.Time
	EndTime   time.Time
}
