package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"io/ioutil"
	"net/http"
)

type Methods struct {
	Settings *internal.PanelSettings
}

func NewMethods(s *internal.PanelSettings) *Methods {
	return &Methods{Settings: s}
}

func (m *Methods) GetRoutes() *mux.Router {
	r := mux.NewRouter()

	// web interface
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/"))).Methods(http.MethodGet)
	// api put\post data
	r.HandleFunc("/api/settings", m.ApiSettingsResolver).Methods(http.MethodPut, http.MethodPost)
	// dashboard
	r.HandleFunc("/api/dashboard", m.GetDashboardInfo).Methods(http.MethodPost)
	// api get data
	r.HandleFunc("/api/disk_partitions", m.GetDiskPartitions).Methods(http.MethodPost)

	return r
}

func (m *Methods) ApiSettingsResolver(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		m.updateSettings(w, r)
	case http.MethodPost:
		m.getSettings(w, r)

	default:
		MethodNotFound(w)
	}
}

// GetDashboardInfo the method that will display statistics in the dashboard will call cpu_load, disk, mem, etc...
func (m *Methods) GetDashboardInfo(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		BadRequest(w, err)
	}

	var dashboard structures.DashboardParams
	if err := json.Unmarshal(data, &dashboard); err != nil {
		BadRequest(w, err)
	}

	cpuLoad, err := internal.GetCPULoad()
	if err != nil {
		BadRequest(w, err)
	}

	io, err := internal.GetDiskInfo(dashboard.Path)
	if err != nil {
		BadRequest(w, err)
	}

	SuccessResponse(w, "Dashboard has been retrieved", structures.Dashboard{
		CPULoad: cpuLoad,
		Mem:     internal.GetVirtualMemory(),
		IO:      io,
	})
}

// getCPUInfo returns cpu statistics.
func (m *Methods) getCPUInfo(w http.ResponseWriter, _ *http.Request) {
	CPUInfo, err := internal.GetCPUInfo()
	if err != nil {
		BadRequest(w, err)
	}

	SuccessResponse(w, "Cpu info has been retrieved", CPUInfo)
}

// GetDiskPartitions returns disk partitions.
func (m *Methods) GetDiskPartitions(w http.ResponseWriter, _ *http.Request) {
	DiskPartitions, err := internal.GetDiskPartitions()
	if err != nil {
		BadRequest(w, err)
	}

	SuccessResponse(w, "Disk partitions has been retrieved", DiskPartitions)
}

// getDiskInfo returns disk usage statistics.
func (m *Methods) getDiskInfo(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	path := param["path"]
	if path == "" {
		BadRequest(w, errors.New("param path is empty"))
	}

	DiskUsage, err := internal.GetDiskInfo(path)
	if err != nil {
		BadRequest(w, err)
	}

	SuccessResponse(w, "Disk usage has been retrieved", DiskUsage)
}

func (m *Methods) getSettings(w http.ResponseWriter, _ *http.Request) {
	settings, err := m.Settings.GetSettings()
	if err != nil {
		BadRequest(w, err)
	}

	SuccessResponse(w, "Settings has been retrieved", settings)
}

// updateSettings updates settings
func (m *Methods) updateSettings(w http.ResponseWriter, r *http.Request) {
	var settings internal.PanelSettings

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &settings)
	if err != nil {
		BadRequest(w, err)

	}

	err = m.Settings.UpdateSettings(settings)
	if err != nil {
		BadRequest(w, err)
	}

	SuccessResponse(w, "Settings has been updated", settings)
}
