package cms

import (
	"context"
	"encoding/json"
	"fmt"

	cms "github.com/alibabacloud-go/cms-20190101/v7/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/provider"
)

// 查询指定监控项的最新监控数据
// 参考: https://next.api.aliyun.com/api/Cms/2019-01-01/DescribeMetricLast
func (o *CmsOperator) DescribeMetricLast(ctx context.Context, req *provider.DescribeMetricLastRequeset) (
	*provider.DatapointSet, error) {
	describeMetricLastRequest := &cms.DescribeMetricLastRequest{
		Namespace:  tea.String(req.Namespace),
		MetricName: tea.String(req.MetricName),
		Period:     tea.String(req.PeriodToString()),
	}

	resp, err := o.client.DescribeMetricLast(describeMetricLastRequest)
	if err != nil {
		return nil, err
	}

	if !*resp.Body.Success {
		o.log.Debug(resp.String())
		return nil, fmt.Errorf(*resp.Body.Message)
	}

	datapoints := []*Datapoint{}
	json.Unmarshal([]byte(tea.StringValue(resp.Body.Datapoints)), &datapoints)

	set := TransferToDatapointSet(datapoints)
	return set, nil
}

func TransferToDatapointSet(items []*Datapoint) *provider.DatapointSet {
	set := provider.NewDatapointSet()
	for i := range items {
		set.Add(TransferOne(items[i]))
	}
	return set
}

func TransferOne(ins *Datapoint) *provider.Datapoint {
	return &provider.Datapoint{
		Timestamp:  ins.Timestamp,
		InstanceId: ins.InstanceId,
		Minimum:    ins.Minimum,
		Maximum:    ins.Maximum,
		Average:    ins.Average,
	}
}

type Datapoint struct {
	Timestamp  int64   `json:"timestamp"`
	InstanceId string  `json:"instanceId"`
	Minimum    float64 `json:"Minimum"`
	Maximum    float64 `json:"Maximum"`
	Average    float64 `json:"Average"`
}
