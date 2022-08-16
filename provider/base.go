package provider

type DescribeRequest struct {
	Id string `json:"id"`
}

func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Rate: 5,
	}
}

type QueryRequest struct {
	Rate float64 `json:"rate"`
}
