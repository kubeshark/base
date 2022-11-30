package models

type HealthResponse struct {
	TargettedPods         []*PodInfo      `json:"targettedPods"`
	ConnectedWorkersCount int             `json:"connectedWorkersCount"`
	WorkersStatus         []*WorkerStatus `json:"workersStatus"`
}

type VersionResponse struct {
	Ver string `json:"ver"`
}
