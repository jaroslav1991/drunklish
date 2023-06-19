package httputils

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"time"
)

type responseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (rw *responseWriter) Write(out []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}
	return rw.ResponseWriter.Write(out)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

type CustomMetrics struct {
	httpRequestDuration *prometheus.HistogramVec
	httpRequestTotal    *prometheus.CounterVec
	requestSize         *prometheus.CounterVec
	responseSize        *prometheus.CounterVec
	goroutines          prometheus.Counter
	goroutinesCount     prometheus.Counter
}

func NewMetrics(reg prometheus.Registerer) *CustomMetrics {
	m := &CustomMetrics{
		httpRequestDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Subsystem: "http",
			Name:      "drunklish_request_duration",
			Help:      "request duration",
			Buckets:   prometheus.LinearBuckets(0.02, 0.02, 10),
		}, []string{"path", "status"}),
		httpRequestTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Subsystem: "http",
			Name:      "drunklish_request_total",
			Help:      "http request total",
		}, []string{"path", "status"}),
		requestSize: prometheus.NewCounterVec(prometheus.CounterOpts{
			Subsystem: "http",
			Name:      "request_size",
			Help:      "request size for handler",
		}, []string{"request_url"}),
		responseSize: prometheus.NewCounterVec(prometheus.CounterOpts{
			Subsystem: "http",
			Name:      "response_size",
			Help:      "response size from handler",
		}, []string{"response_url"}),
		goroutines: prometheus.NewCounter(prometheus.CounterOpts{
			Subsystem: "go",
			Name:      "goroutines",
			Help:      "check goroutines",
		}),
		goroutinesCount: prometheus.NewCounter(prometheus.CounterOpts{
			Subsystem: "go",
			Name:      "goroutines_count",
			Help:      "check goroutines count",
		}),
	}

	reg.MustRegister(m.httpRequestDuration)
	reg.MustRegister(m.httpRequestTotal)
	reg.MustRegister(m.requestSize)
	reg.MustRegister(m.responseSize)
	reg.MustRegister(m.goroutines)
	reg.MustRegister(m.goroutinesCount)

	return m
}

func (m *CustomMetrics) Middleware(mtr *CustomMetrics, wrapper http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ww := &responseWriter{ResponseWriter: w}

		start := time.Now()

		wrapper.ServeHTTP(ww, r)

		mtr.httpRequestDuration.With(prometheus.Labels{"path": r.URL.Path, "status": strconv.Itoa(ww.statusCode)}).Observe(float64(time.Since(start).Milliseconds()))
		mtr.httpRequestTotal.With(prometheus.Labels{"path": r.URL.Path, "status": strconv.Itoa(ww.statusCode)}).Inc()
	}
}
