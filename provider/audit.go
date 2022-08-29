package provider

import "time"

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
