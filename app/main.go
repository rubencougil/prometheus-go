package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")

var (
	c = promauto.NewCounter(prometheus.CounterOpts{
		Name: "geekshubs_app_sample_metric",
		Help: "Sample metric for GeeksHubs devops bootcamp",
	})

	h = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "geekshubs_app_sample_histogram",
		Help: "Sample metric for GeeksHubs devops bootcamp",
	})
)

func main() {

	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			h.Observe(float64(rand.Intn(100 - 0 + 1) + 0))
			c.Inc()
			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
