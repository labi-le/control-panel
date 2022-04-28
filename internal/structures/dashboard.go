package structures

type Dashboard struct {
	CPULoad *CPULoad   `json:"cpu_load"`
	IO      *UsageStat `json:"io"`
	Mem     *Memory    `json:"mem"`
}

type DashboardParams struct {
	Path string `json:"path"`
	//...
}
