package internal

import (
	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
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

// GetAvg returns cpu load
func GetAvg() (*load.AvgStat, error) {
	return load.Avg()
}

// GetCPUTimes GetCpuTimes returns cpu times
func GetCPUTimes() ([]cpu.TimesStat, error) {
	return cpu.Times(true)
}
