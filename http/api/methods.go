package api

import (
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/structures"
	"net/http"
	"time"
)

type Methods struct {
	resp structures.Response
	w    http.ResponseWriter
	db   *internal.DB
}

func NewMethods(w http.ResponseWriter, db *internal.DB) *Methods {
	return &Methods{
		resp: structures.Response{
			Version: "0.1",
			Time:    time.Now(),
		},
		w:  w,
		db: db,
	}
}

// GetVirtualMemory returns virtual memory statistics.
func (m *Methods) GetVirtualMemory() *Methods {
	return m.SuccessResponse("Mem has been retrieved", internal.GetVirtualMemory())
}

// GetCPUInfo returns cpu statistics.
func (m *Methods) GetCPUInfo() *Methods {
	CPUInfo, err := internal.GetCPUInfo()
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Cpu info has been retrieved", CPUInfo)
}

// GetCPUAvg returns cpu usage statistics.
func (m *Methods) GetCPUAvg() *Methods {
	CPUUsage, err := internal.GetAvg()
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Cpu average has been retrieved", CPUUsage)
}

// GetCPUTimes GetCpuTimes returns cpu usage statistics.
func (m *Methods) GetCPUTimes() *Methods {
	CPUTimes, err := internal.GetCPUTimes()
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Cpu times has been retrieved", CPUTimes)
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
func (m *Methods) UpdateSettings(settings structures.PanelSettings) *Methods {
	err := m.db.UpdateSettings(settings)
	if err != nil {
		return m.BadRequest(err)
	}

	return m.SuccessResponse("Settings has been updated", []string{})
}

func (m *Methods) SuccessResponse(msg string, data interface{}) *Methods {
	m.resp.Success = true
	m.resp.Message = msg
	m.resp.Data = data

	return m
}

func (m *Methods) BadRequest(err error) *Methods {
	m.resp.Success = false
	m.resp.Message = err.Error()

	return m
}

func (m *Methods) MethodNotFound() *Methods {
	m.resp.Success = false
	m.resp.Message = "Method not found"

	return m
}
