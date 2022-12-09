package models

import v1 "k8s.io/api/core/v1"

type PodInfo struct {
	Name      string `json:"name"`
	NodeName  string `json:"nodeName"`
	Namespace string `json:"namespace"`
}

type TargettedPodStatus struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	IsTargetted bool   `json:"isTargetted"`
}

type WorkerStatus struct {
	Name     string `json:"name"`
	NodeName string `json:"nodeName"`
	Status   string `json:"status"`
}

type NodeToPodsMap map[string][]v1.Pod

func (np NodeToPodsMap) Summary() map[string][]string {
	summary := make(map[string][]string)
	for node, pods := range np {
		for _, pod := range pods {
			summary[node] = append(summary[node], pod.Namespace+"/"+pod.Name)
		}
	}

	return summary
}
