package config

import (
    "io/ioutil"
    
    "gopkg.in/yaml.v2"
)

// Config holds application configuration
type Config struct {
    Cloud struct {
        Provider  string `yaml:"provider"`
        Region    string `yaml:"region"`
        AccessKey string `yaml:"accessKey"`
        SecretKey string `yaml:"secretKey"`
    } `yaml:"cloud"`
    
    Scaling struct {
        MinNodes     int     `yaml:"minNodes"`
        MaxNodes     int     `yaml:"maxNodes"`
        ScaleUpThreshold float64 `yaml:"scaleUpThreshold"`
        ScaleDownThreshold float64 `yaml:"scaleDownThreshold"`
    } `yaml:"scaling"`
    
    Monitoring struct {
        Namespace string        `yaml:"namespace"`
        Interval  string        `yaml:"interval"`
    } `yaml:"monitoring"`
}

// LoadConfig loads configuration from a file
func LoadConfig(path string) (*Config, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    
    var config Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return nil, err
    }
    
    return &config, nil
}