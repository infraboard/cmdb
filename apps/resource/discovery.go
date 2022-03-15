package resource

func (r *Resource) PrometheusTarget() *PrometheusTarget {
	return &PrometheusTarget{
		Targets: []string{},
		Labels:  map[string]string{},
	}
}

type PrometheusTarget struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}
