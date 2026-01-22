package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Prometheus metric
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Increment metric
	httpRequestsTotal.WithLabelValues(r.Method, "/").Inc()

	fmt.Fprintln(w, "Hello from Jenkins + Helm Real Project 🚀")
}

func main() {
	// Register metric
	prometheus.MustRegister(httpRequestsTotal)

	// App endpoints
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

