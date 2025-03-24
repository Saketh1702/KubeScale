package metrics

import (
    "net/http"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    scaleOperationsCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "kubescale_scaling_operations_total",
            Help: "Total number of scaling operations",
        },
        []string{"direction", "provider"},
    )
    
    predictionAccuracyGauge = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "kubescale_prediction_accuracy",
            Help: "Accuracy of workload predictions",
        },
        []string{"resource_type"},
    )
)

// InitMetrics initializes the metrics server
func InitMetrics() {
    // Register metrics
    prometheus.MustRegister(scaleOperationsCounter)
    prometheus.MustRegister(predictionAccuracyGauge)
    
    // Expose metrics endpoint
    http.Handle("/metrics", promhttp.Handler())
    go func() {
        http.ListenAndServe(":9090", nil)
    }()
}

// RecordScaleOperation records a scaling operation
func RecordScaleOperation(direction, provider string) {
    scaleOperationsCounter.WithLabelValues(direction, provider).Inc()
}

// SetPredictionAccuracy sets the prediction accuracy
func SetPredictionAccuracy(resourceType string, accuracy float64) {
    predictionAccuracyGauge.WithLabelValues(resourceType).Set(accuracy)
}