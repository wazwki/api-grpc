package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ObserveRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Duration of requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func init() {
	prometheus.MustRegister(ObserveRequestDuration)
}
