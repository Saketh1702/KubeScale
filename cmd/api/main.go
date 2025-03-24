package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/Saketh1702/KubeScale/pkg/monitor"
    "github.com/Saketh1702/KubeScale/pkg/cloud/aws"
)

func main() {
    // Initialize components
    collector := monitor.NewKubernetesCollector("default", 30*time.Second)
    predictor := monitor.NewTimeSeriesPredictor(24*time.Hour, 1*time.Hour)
    cloudProvider := aws.NewAWSProvider("us-west-2", "dummy-key", "dummy-secret")
    
    // Setup API server
    r := gin.Default()
    
    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "healthy",
            "version": "0.1.0",
        })
    })
    
    // API endpoints
    r.GET("/metrics", func(c *gin.Context) {
        metrics, err := collector.CollectMetrics(c)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, metrics)
    })
    
    r.POST("/scale", func(c *gin.Context) {
        // TODO: Implement scaling endpoint
        c.JSON(http.StatusOK, gin.H{"message": "Scaling request received"})
    })
    
    // Start server
    log.Println("Starting KubeScale API server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}