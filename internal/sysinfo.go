package internal

import "github.com/shirou/gopsutil/v3/mem"
import "github.com/shirou/gopsutil/v3/cpu"

//GetVirtualMemory returns virtual memory info
func GetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

//GetProcLoad returns load info
func GetProcLoad() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

//etc
