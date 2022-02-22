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
	Mem, err := internal.GetVirtualMemory()
	if err != nil {
		m.resp.Success = false
		m.resp.Message = err.Error()

		return m
	}
	m.resp.Success = true
	m.resp.Message = "Mem has been retrieved"
	m.resp.Data = Mem

	return m
}

// GetCpuInfo returns cpu statistics.
func (m *Methods) GetCpuInfo() *Methods {
	Cpu, err := internal.GetCpuInfo()
	if err != nil {
		m.resp.Success = false
		m.resp.Message = err.Error()

		return m
	}
	m.resp.Success = true
	m.resp.Message = "Cpu has been retrieved"
	m.resp.Data = Cpu

	return m
}

// GetCpuAvg returns cpu usage statistics.
func (m *Methods) GetCpuAvg() *Methods {
	CpuUsage, err := internal.GetAvg()
	if err != nil {
		m.resp.Success = false
		m.resp.Message = err.Error()

		return m
	}
	m.resp.Success = true
	m.resp.Message = "Cpu average has been retrieved"
	m.resp.Data = CpuUsage

	return m
}

// GetCpuTimes returns cpu usage statistics.
func (m *Methods) GetCpuTimes() *Methods {
	CpuUsage, err := internal.GetCpuTimes()
	if err != nil {
		m.resp.Success = false
		m.resp.Message = err.Error()

		return m
	}
	m.resp.Success = true
	m.resp.Message = "Cpu average has been retrieved"
	m.resp.Data = CpuUsage

	return m
}

func (m *Methods) GetSettings() *Methods {
	settings, err := m.db.GetSettings()
	if err != nil {
		m.resp.Success = false
		m.resp.Message = err.Error()

		return m
	}
	m.resp.Success = true
	m.resp.Message = "Settings has been retrieved"
	m.resp.Data = settings

	return m
}

//UpdateSettings updates settings
func (m *Methods) UpdateSettings(settings structures.PanelSettings) *Methods {
	err := m.db.UpdateSettings(settings)
	if err != nil {
		m.resp.Success = false
		m.resp.Message = err.Error()

		return m
	}
	m.resp.Success = true
	m.resp.Message = "Settings has been updated"

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
