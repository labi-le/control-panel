package internal

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

// GetVirtualMemory returns virtual memory info
func GetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
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
