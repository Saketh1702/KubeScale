package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/Saketh1702/KubeScale/pkg/monitor"
	"github.com/Saketh1702/KubeScale/pkg/cloud"
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
		// Get metrics data
		metrics, err := collector.CollectMetrics(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		// Use the predictor to predict workload
		prediction, err := predictor.PredictWorkload(metrics)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		// Determine scaling action based on prediction
		var action cloud.ScalingAction
		if prediction.ExpectedCPU > 0.75 {
			action = cloud.ScalingAction{
				InstanceType: "t3.medium",
				TargetCount: 5,
				Region: "us-west-2",
			}
			
			// Use cloud provider to scale
			err = cloudProvider.ScaleCluster(c, action)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		
		c.JSON(http.StatusOK, gin.H{
			"message": "Scaling request processed",
			"prediction": prediction,
			"action": action,
		})
	})
    
    // Start server
    log.Println("Starting KubeScale API server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}