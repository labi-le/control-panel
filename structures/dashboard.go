package structures

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type Dashboard struct {
	CPULoad *CPULoad        `json:"cpu_load"`
	IO      *disk.UsageStat `json:"io"`
	Mem     *Memory         `json:"mem"`
}

type DashboardParams struct {
	Path string `json:"path"`
	// ...
}
