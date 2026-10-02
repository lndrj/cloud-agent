package metrics

import (
	"agent/models"
	"os"
	"time"

	cpuLib "github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func CollectMetrics() (models.AgentMetrics, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return models.AgentMetrics{}, err
	}

	memory, err := mem.VirtualMemory()
	if err != nil {
		return models.AgentMetrics{}, err
	}

	uptime, err := host.Uptime()
	if err != nil {
		return models.AgentMetrics{}, err
	}

	cpu, err := cpuLib.Percent(time.Second, false)
	if err != nil {
		return models.AgentMetrics{}, err
	}
	cpuPercent := 0.0
	if len(cpu) > 0 {
		cpuPercent = cpu[0]
	}

	diskUse, err := disk.Usage("/")
	if err != nil {
		return models.AgentMetrics{}, err
	}

	currentMetrics := models.AgentMetrics{
		Hostname:          hostname,
		Uptime:            uptime,
		MemoryAvailableMb: memory.Available / 1024 / 1024,
		MemoryUsedMb:      memory.Used / 1024 / 1024,
		CpuPercent:        cpuPercent,
		DiskUsedPercent:   diskUse.UsedPercent,
	}

	return currentMetrics, nil
}
