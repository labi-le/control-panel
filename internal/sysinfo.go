package internal

import (
	io "github.com/mackerelio/go-osstat/disk"
	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
)

// Memory in kibibyte
type Memory struct {
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
}

// GetVirtualMemory returns virtual memory info
func GetVirtualMemory() *Memory {
	return &Memory{
		Total: memory.TotalMemory(),
		Free:  memory.FreeMemory(),
	}
}

// GetCPUInfo returns cpu info
func GetCPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

type CPULoad struct {
	Load float64 `json:"load"`
}

// GetCPULoad returns cpu load
func GetCPULoad() (*CPULoad, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	return &CPULoad{Load: percent[0]}, nil
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
