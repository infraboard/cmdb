package provider

import (
	"context"
)

type CmsOperator interface {
	DescribeMetricLast(context.Context, *DescribeMetricLastRequeset) (*DatapointSet, error)
}

func NewDescribeMetricLastRequeset(namespace, metricname string) *DescribeMetricLastRequeset {
	return &DescribeMetricLastRequeset{
		Namespace:  namespace,
		MetricName: metricname,
	}
}

type DescribeMetricLastRequeset struct {
	Namespace  string
	MetricName string
}

type DatapointSet struct {
	Items []*Datapoint
}

type Datapoint struct {
	Timestamp  int64
	InstanceId string
	Minimum    float64
	Maximum    float64
	Average    float64
}
