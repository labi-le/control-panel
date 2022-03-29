package api

import (
	"github.com/labi-le/control-panel/internal"
	structures2 "github.com/labi-le/control-panel/internal/structures"
	"net/http"
	"time"
)

type Methods struct {
	resp structures2.Response
	w    http.ResponseWriter
	db   *internal.DB
}

func NewMethods(w http.ResponseWriter, db *internal.DB) *Methods {
	return &Methods{
		resp: structures2.Response{
			Version: "0.1",
			Time:    time.Now(),
		},
		w:  w,
		db: db,
	}
}

// GetDashboardInfo the method that will display statistics in the dashboard will call cpu_load, disk, mem, etc...
func (m *Methods) GetDashboardInfo(dashboardParams structures2.DashboardParams) *Methods {
	cpuLoad, err := internal.GetCPULoad()
	if err != nil {
		return m.BadRequest(err)
	}

	io, err := internal.GetDiskInfo(dashboardParams.Path)
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Dashboard has been retrieved", structures2.Dashboard{
		CPULoad: cpuLoad,
		Mem:     internal.GetVirtualMemory(),
		IO:      io,
	})
}

// GetCPUInfo returns cpu statistics.
func (m *Methods) GetCPUInfo() *Methods {
	CPUInfo, err := internal.GetCPUInfo()
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Cpu info has been retrieved", CPUInfo)
}

// GetDiskPartitions returns disk partitions.
func (m *Methods) GetDiskPartitions() *Methods {
	DiskPartitions, err := internal.GetDiskPartitions()
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Disk partitions has been retrieved", DiskPartitions)
}

// GetDiskInfo returns disk usage statistics.
func (m *Methods) GetDiskInfo(path string) *Methods {
	DiskUsage, err := internal.GetDiskInfo(path)
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Disk usage has been retrieved", DiskUsage)
}

func (m *Methods) GetSettings() *Methods {
	settings, err := m.db.GetSettings()
	if err != nil {
		return m.BadRequest(err)
	}

	m.resp.Success = true
	m.resp.Message = "Settings has been retrieved"
	m.resp.Data = settings

	return m
}

// UpdateSettings updates settings
func (m *Methods) UpdateSettings(settings structures2.PanelSettings) *Methods {
	err := m.db.UpdateSettings(settings)
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Settings has been updated", []string{})
}
