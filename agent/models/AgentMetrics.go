package models

type AgentMetrics struct {
	Hostname          string  `json:"hostname"`
	Uptime            uint64  `json:"uptime"`
	MemoryAvailableMb uint64  `json:"memory_available_mb"`
	MemoryUsedMb      uint64  `json:"memory_used_mb"`
	CpuPercent        float64 `json:"cpu_percent"`
	DiskUsedPercent   float64 `json:"disk_used_percent"`
}
