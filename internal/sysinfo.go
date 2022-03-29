package internal

import (
	structures2 "github.com/labi-le/control-panel/internal/structures"
	io "github.com/mackerelio/go-osstat/disk"
	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
)

// GetVirtualMemory returns virtual memory info
func GetVirtualMemory() *structures2.Memory {
	return &structures2.Memory{
		Total: memory.TotalMemory(),
		Free:  memory.FreeMemory(),
	}
}

// GetCPUInfo returns cpu info
func GetCPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// GetCPULoad returns cpu load
func GetCPULoad() (*structures2.CPULoad, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	return &structures2.CPULoad{Load: percent[0]}, nil
}

// GetDiskIO returns disk usage
func GetDiskIO() ([]io.Stats, error) {
	return io.Get()
}

// GetDiskPartitions returns disk partitions
func GetDiskPartitions() ([]disk.PartitionStat, error) {
	return disk.Partitions(true)
}

// GetDiskInfo returns disk info
func GetDiskInfo(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}

// GetCPUTimes GetCpuTimes returns cpu times
func GetCPUTimes() ([]cpu.TimesStat, error) {
	return cpu.Times(true)
}
