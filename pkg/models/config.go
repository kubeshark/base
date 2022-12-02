package models

import "github.com/op/go-logging"

type Resources struct {
	CpuLimit       string `yaml:"cpu-limit" default:"750m"`
	MemoryLimit    string `yaml:"memory-limit" default:"1Gi"`
	CpuRequests    string `yaml:"cpu-requests" default:"50m"`
	MemoryRequests string `yaml:"memory-requests" default:"50Mi"`
}

type Config struct {
	MaxDBSizeBytes     int64         `json:"maxDBSizeBytes"`
	InsertionFilter    string        `json:"insertionFilter"`
	PullPolicy         string        `json:"pullPolicy"`
	LogLevel           logging.Level `json:"logLevel"`
	WorkerResources    Resources     `json:"workerResources"`
	ResourcesNamespace string        `json:"resourceNamespace"`
	DatabasePath       string        `json:"databasePath"`
}
