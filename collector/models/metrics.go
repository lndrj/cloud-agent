package models

import "time"

type AgentMetrics struct {
	Hostname          string  `json:"hostname"`
	Uptime            uint64  `json:"uptime"`
	MemoryAvailableMb uint64  `json:"memory_available_mb"`
	MemoryUsedMb      uint64  `json:"memory_used_mb"`
	CpuPercent        float64 `json:"cpu_percent"`
	DiskUsedPercent   float64 `json:"disk_used_percent"`
}

type AgentState struct {
	Metrics  AgentMetrics `json:"metrics"`
	LastSeen time.Time    `json:"last_seen"`
	Online   bool         `json:"online"`
}

type MetricsSnapshot struct {
	AgentId     string       `json:"agent_id"`
	Metrics     AgentMetrics `json:"metrics"`
	CollectedAt time.Time    `json:"collected_at"`
}
