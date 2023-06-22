package internal

import (
	"github.com/labi-le/control-panel/internal/types"
	io "github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
)

// GetVirtualMemory returns virtual memory info
func GetVirtualMemory() (*types.Memory, error) {
	mem, err := memory.Get()
	if err != nil {
		return &types.Memory{}, err
	}

	return &types.Memory{
		Total:  mem.Total,
		Free:   mem.Free,
		Used:   mem.Used,
		Cached: mem.Cached,
	}, nil
}

// GetCPUInfo returns cpu info
func GetCPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// GetCPULoad returns cpu load
func GetCPULoad() (*types.CPULoad, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	return &types.CPULoad{Load: percent[0]}, nil
}

// GetDiskIO returns disk usage
func GetDiskIO() ([]io.Stats, error) {
	return io.Get()
}

// GetDiskPartitions returns disk partitions
func GetDiskPartitions() ([]types.PartitionStat, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	var result []types.PartitionStat
	for _, partition := range partitions {
		result = append(result, types.PartitionStat{
			Device:     partition.Device,
			Mountpoint: partition.Mountpoint,
			Fstype:     partition.Fstype,
			Opts:       partition.Opts,
		})
	}

	return result, nil
}

// GetDiskInfo returns disk info
func GetDiskInfo(path string) (*types.UsageStat, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}

	return &types.UsageStat{
		Path:              usage.Path,
		Fstype:            usage.Fstype,
		Total:             usage.Total,
		Free:              usage.Free,
		Used:              usage.Used,
		UsedPercent:       usage.UsedPercent,
		InodesTotal:       usage.InodesTotal,
		InodesUsed:        usage.InodesUsed,
		InodesFree:        usage.InodesFree,
		InodesUsedPercent: usage.InodesUsedPercent,
	}, err
}

// GetCPUTimes GetCpuTimes returns cpu times
func GetCPUTimes() ([]cpu.TimesStat, error) {
	return cpu.Times(true)
}
