package cloud

import "context"

// ScalingAction defines what scaling action to take
type ScalingAction struct {
    InstanceType  string
    TargetCount   int
    Region        string
}

// Provider interface for cloud operations
type Provider interface {
    ScaleCluster(ctx context.Context, action ScalingAction) error
    GetCurrentCapacity(ctx context.Context, clusterID string) (int, error)
}