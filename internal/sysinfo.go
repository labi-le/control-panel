package internal

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

//GetVirtualMemory returns virtual memory info
func GetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

//GetCpuInfo returns cpu info
func GetCpuInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

//GetAvg returns cpu load
func GetAvg() (*load.AvgStat, error) {
	return load.Avg()
}

//GetCpuTimes returns cpu times
func GetCpuTimes() ([]cpu.TimesStat, error) {
	return cpu.Times(true)
}
