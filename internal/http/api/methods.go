package api

import (
	"encoding/json"
	"errors"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"io/ioutil"
	"net/http"

	_ "github.com/labi-le/control-panel/docs"
)

type Methods struct {
	Settings *internal.PanelSettings
}

func NewMethods(s *internal.PanelSettings) *Methods {
	return &Methods{Settings: s}
}

func (m *Methods) GetRoutes() *echo.Echo {
	e := echo.New()

	e.Static("/", "./frontend/")
	e.Router().Add(http.MethodGet, "/", func(c echo.Context) error {
		return c.File("./frontend/index.html")
	})

	e.Router().Add(http.MethodGet, "/swagger/*", echoSwagger.WrapHandler)

	e.Router().Add(http.MethodGet, "/api/settings", m.getSettings)
	e.Router().Add(http.MethodPut, "/api/settings", m.updateSettings)
	e.Router().Add(http.MethodPost, "/api/dashboard", m.GetDashboardInfo)
	e.Router().Add(http.MethodPost, "/api/disk_partitions", m.GetDiskPartitions)

	// web interface
	// api put\post data
	//// dashboard
	//// api get data
	//r.HandleFunc("/api/disk_partitions", m.GetDiskPartitions).Methods(http.MethodPost)

	return e
}

// GetDashboardInfo the method that will display statistics in the dashboard will call cpu_load, disk, mem, etc...
// @Summary      Dashboard info
// @Description  Get system information and state
// @Tags         info
// @Accept       json
// @Produce      json
// @Param        json  body      structures.DashboardParams  true  "Dashboard info"
// @Success      200   {object}  structures.Dashboard        "Dashboard info"
// @Failure      500  {object}  structures.Response
// @Router       /dashboard [post]
func (m *Methods) GetDashboardInfo(c echo.Context) error {
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		BadRequest(c.Response(), err)
	}

	var dashboard structures.DashboardParams
	if err := json.Unmarshal(data, &dashboard); err != nil {
		BadRequest(c.Response(), err)
	}

	cpuLoad, err := internal.GetCPULoad()
	if err != nil {
		BadRequest(c.Response(), err)
	}

	io, err := internal.GetDiskInfo(dashboard.Path)
	if err != nil {
		BadRequest(c.Response(), err)
	}

	SuccessResponse(c.Response(), "Dashboard has been retrieved", structures.Dashboard{
		CPULoad: cpuLoad,
		Mem:     internal.GetVirtualMemory(),
		IO:      io,
	})

	return nil
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
// @Summary      Disk partitions
// @Description  Get disk partitions
// @Tags         info
// @Accept       json
// @Produce      json
// @Success      200  {object}  []structures.PartitionStat  "Disk partitions"
// @Failure      500  {object}  structures.Response
// @Router       /disk_partitions [post]
func (m *Methods) GetDiskPartitions(c echo.Context) error {
	DiskPartitions, err := internal.GetDiskPartitions()
	if err != nil {
		BadRequest(c.Response(), err)
	}

	SuccessResponse(c.Response(), "Disk partitions has been retrieved", DiskPartitions)
	return nil
}

// getDiskInfo returns disk usage statistics.
func (m *Methods) getDiskInfo(c echo.Context) error {
	// get var from request
	path := c.QueryParam("path")
	if path == "" {
		BadRequest(c.Response(), errors.New("param path is empty"))
	}

	DiskUsage, err := internal.GetDiskInfo(path)
	if err != nil {
		BadRequest(c.Response(), err)
	}

	SuccessResponse(c.Response(), "Disk usage has been retrieved", DiskUsage)
	return nil
}

// getSettings
// @Summary      Get settings
// @Description  Get settings
// @Tags         settings
// @Accept       json
// @Produce      json
// @Success      200  {object}  internal.PanelSettings  "Settings"
// @Failure      500   {object}  structures.Response
// @Router       /settings [post]
func (m *Methods) getSettings(c echo.Context) error {
	settings, err := m.Settings.GetSettings()
	if err != nil {
		BadRequest(c.Response(), err)
	}

	SuccessResponse(c.Response(), "Settings has been retrieved", settings)
	return nil
}

// updateSettings
// @Summary      Update settings
// @Description  Update settings
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        json  body      internal.PanelSettings  true  "Settings"
// @Success      200   {object}  internal.PanelSettings  "Settings"
// @Failure      400   {object}  structures.Response
// @Failure      500   {object}  structures.Response
// @Router       /settings [put]
func (m *Methods) updateSettings(c echo.Context) error {
	var settings internal.PanelSettings

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &settings)
	if err != nil {
		BadRequest(c.Response(), err)

	}

	err = m.Settings.UpdateSettings(settings)
	if err != nil {
		BadRequest(c.Response(), err)
	}

	SuccessResponse(c.Response(), "Settings has been updated", settings)
	return nil
}
