package resource

import "fmt"

const (
	PROMETHEUS_SCRAPE = "prometheus.io/%/enabled"
	PROMETHEUS_PORT   = "prometheus.io/%s/port"
	PROMETHEUS_PATH   = "prometheus.io/%s/path"
)

func NewPrometheusScrapeTag() *TagSelector {
	return &TagSelector{
		Key:    PROMETHEUS_SCRAPE,
		Values: []string{},
	}
}

func (r *Resource) PrometheusEndpont() (string, error) {
	if len(r.Information.PrivateIp) == 0 {
		return "", fmt.Errorf("instance no private ip")
	}

	ip := r.Information.PrivateIp[0]

	port := r.GetTagValueOne(PROMETHEUS_PORT)
	if port == "" {
		switch r.Base.ResourceType {
		case Type_RDS:
			port = "6221"
		default:
			port = "9100"
		}
	}

	return fmt.Sprintf("%s:%s", ip, port), nil
}

func (r *Resource) PrometheusTarget() (*PrometheusTarget, error) {
	ep, err := r.PrometheusEndpont()
	if err != nil {
		return nil, err
	}

	t := &PrometheusTarget{
		Targets: []string{ep},
		Labels: map[string]string{
			"domain":      r.Base.Domain,
			"namespace":   r.Base.Namespace,
			"env":         r.Base.Env,
			"accout":      r.Information.SyncAccount,
			"vendor":      r.Base.Vendor.String(),
			"region":      r.Base.Region,
			"instance_id": r.Base.Id,
		},
	}

	// 重写metric path
	path := r.GetTagValueOne(PROMETHEUS_PATH)
	if path != "" {
		t.Labels["__metrics_path__"] = path
	}

	return t, nil
}

type PrometheusTarget struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}
