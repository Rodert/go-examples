package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	MetricHttpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "promdemo",
			Subsystem: "demo",
			Name:      "http_request_total",
			Help:      "http request total",
		},
		[]string{"from"},
	)
)

func init() {
	prometheus.MustRegister(MetricHttpRequestTotal)
}

func main() {

	go func() {
		muxProm := http.NewServeMux()
		muxProm.Handle("/metrics", promhttp.Handler())

		http.ListenAndServe(":9527", muxProm)
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		values := req.URL.Query()
		from := values.Get("from")
		MetricHttpRequestTotal.WithLabelValues(from).Inc()
		w.Write([]byte("Hello,from " + from))
	})

	http.ListenAndServe(":8080", nil)

}
