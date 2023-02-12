package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/alibabacloud-go/tea/tea"
)

type CmsOperator interface {
	DescribeMetricLast(context.Context, *DescribeMetricLastRequeset) (*DatapointSet, error)
}

func NewDescribeMetricLastRequeset(namespace, metricname string) *DescribeMetricLastRequeset {
	return &DescribeMetricLastRequeset{
		Namespace:  namespace,
		MetricName: metricname,
		Period:     60,
	}
}

type DescribeMetricLastRequeset struct {
	// 云服务的命名空间
	Namespace string
	// 云服务的监控项名称
	MetricName string
	// 监控项的统计周期, 单位：秒
	Period int64
	// 查询监控项的开始时间
	StartTime time.Time
	// 查询监控项的结束时间
	EndTime time.Time
	// 维度Map
	Dimensions map[string]string
}

func (req *DescribeMetricLastRequeset) PeriodToString() string {
	return fmt.Sprintf("%d", req.Period)
}

func NewDatapointSet() *DatapointSet {
	return &DatapointSet{
		Items: []*Datapoint{},
	}
}

type DatapointSet struct {
	Items []*Datapoint
}

func (d *DatapointSet) String() string {
	return tea.Prettify(d)
}

func (s *DatapointSet) Add(item *Datapoint) {
	s.Items = append(s.Items, item)
}

type Datapoint struct {
	Timestamp  int64   `json:"timestamp"`
	InstanceId string  `json:"instance_id"`
	Minimum    float64 `json:"minimum"`
	Maximum    float64 `json:"maximum"`
	Average    float64 `json:"average"`
}

func (d *Datapoint) String() string {
	return tea.Prettify(d)
}
