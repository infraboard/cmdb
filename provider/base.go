package provider

import "fmt"

func NewDescribeRequest(id string) *DescribeRequest {
	return &DescribeRequest{
		Id: id,
	}
}

type DescribeRequest struct {
	Id string `json:"id"`
}

func (req *DescribeRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("resource id required")
	}

	return nil
}

func NewQueryRequestWithRate(rate int32) *QueryRequest {
	return &QueryRequest{
		Rate: float64(rate),
	}
}

func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Rate: 5,
	}
}

type QueryRequest struct {
	Rate float64 `json:"rate"`
}
