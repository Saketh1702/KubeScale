# KubeScale

## 🚀 Overview

KubeScale is an intelligent, adaptive scaling service designed to optimize resource utilization for containerized applications. Built with Go and following microservices architecture principles, it provides predictive scaling based on workload analysis rather than simple threshold-based approaches.

## 🏗️ Architecture


https://github.com/user-attachments/assets/3f009ca6-d751-4166-bb7b-2e7a87c85a09




KubeScale consists of the following core components:

1. **Workload Monitor** - Collects and analyzes resource utilization metrics from Kubernetes clusters
2. **Prediction Engine** - Uses time series analysis to forecast future resource demands
3. **Scaling Controller** - Implements custom scaling logic based on predictions
4. **Cloud Provider Integration** - Connects with cloud providers (AWS, GCP) for infrastructure provisioning
5. **REST API** - Provides configuration and manual control capabilities
6. **Metrics Exporter** - Reports operational metrics for observability

## 🛠️ Technologies

- **Language**: Go
- **Architecture**: Microservices
- **Container Orchestration**: Kubernetes-inspired 
- **Cloud Integration**: AWS, GCP
- **Metrics & Monitoring**: Prometheus-compatible
- **API**: REST with JSON

## ✨ Key Features

- Advanced workload prediction using time-series analysis
- Custom scaling algorithms that minimize resource costs while maintaining performance
- Multi-cloud provider support
- Detailed metrics and observability
- RESTful API for configuration and manual control
- Containerized deployment with Docker

## 🔧 Installation & Deployment

### Prerequisites

- Docker
- Go 1.20 or higher

### Running with Docker

```bash
# Pull the image
docker pull saketh1702/kubescale:latest

# Run the container
docker run -p 8080:8080 -p 9090:9090 \
  -e AWS_ACCESS_KEY=your_key \
  -e AWS_SECRET_KEY=your_secret \
  saketh1702/kubescale:latest
```

## Building from source
```bash
git clone https://github.com/Saketh1702/KubeScale.git

# Change to project directory
cd KubeScale

# Build and run
go mod download
go run cmd/api/main.go
```

## API Endpoints
``` bash
GET/health - Health check endpoint
GET/metrics - Prometheus-compatible metrics endpoint
GET/scale - Trigger scaling operation
GET/config - Get current configuration
PUT/config  - Update configuration
```

## Project Structure
```bash
kubescale/
├── cmd/                # Application entry points
│   └── api/            # API server 
├── pkg/                # Core packages
│   ├── monitor/        # Workload monitoring
│   ├── cloud/          # Cloud provider integrations
│   │   ├── aws/        # AWS implementation
│   │   └── gcp/        # GCP implementation
│   ├── metrics/        # Metrics handling
│   └── config/         # Configuration
├── api/                # API definitions
├── deploy/             # Deployment configurations
├── Dockerfile          # Container definition
└── config.yaml         # Default configuration
```

## 👨‍💻 Development
KubeScale was developed to demonstrate advanced techniques in Go microservices, cloud-native applications, and intelligent resource management. It showcases knowledge of Kubernetes concepts, container orchestration, and cloud provider integrations
