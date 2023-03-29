package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// var (
// 	MetricHttpRequestTotal = prometheus.NewCounterVec(
// 		prometheus.CounterOpts{
// 			Namespace: "promdemo",
// 			Subsystem: "demo",
// 			Name:      "http_request_total",
// 			Help:      "http request total",
// 		},
// 		[]string{"from"},
// 	)
// )

// func init() {
// 	prometheus.MustRegister(MetricHttpRequestTotal)
// }

// func main() {

// 	go func() {
// 		muxProm := http.NewServeMux()
// 		muxProm.Handle("/metrics", promhttp.Handler())

// 		http.ListenAndServe(":9527", muxProm)
// 	}()

// 	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
// 		values := req.URL.Query()
// 		from := values.Get("from")
// 		MetricHttpRequestTotal.WithLabelValues(from).Inc()
// 		w.Write([]byte("Hello,from " + from))
// 	})

// 	http.ListenAndServe(":8080", nil)

// }

var (
	successfulRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myapp_successful_requests_total",
		Help: "Total number of successful HTTP requests.",
	})
	failedRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myapp_failed_requests_total",
		Help: "Total number of failed HTTP requests.",
	})
	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "myapp_request_duration_seconds",
		Help:    "Duration of HTTP requests in seconds.",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 5), // Buckets for the histogram
	})
)

func main() {
	fmt.Println("start")
	// Register Prometheus metrics with the global Prometheus registry
	prometheus.MustRegister(successfulRequests)
	prometheus.MustRegister(failedRequests)
	prometheus.MustRegister(requestDuration)

	// Start the HTTP server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8090", nil)
	}()

	// Start the actual scraping
	for {
		start := time.Now()

		// Make the HTTP request and track the result
		if err := scrape(); err != nil {
			failedRequests.Inc()
		} else {
			successfulRequests.Inc()
		}

		//
		duration := time.Since(start).Seconds()
		requestDuration.Observe(duration)

		// Wait for a fixed interval before making the next request
		time.Sleep(30 * time.Second)
	}
}

func scrape() error {
	// Perform the actual scraping here
	// ...

	// Return an error if the scraping fails
	return nil
}
