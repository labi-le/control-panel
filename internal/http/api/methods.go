package api

import (
	"encoding/json"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
	"time"
)

type Methods struct {
	Settings *internal.PanelSettings
	logger   *logrus.Logger
}

func (m *Methods) Logger() *logrus.Logger {
	return m.logger
}

func (m *Methods) successResponse(ws *websocket.Conn, d ...any) {
	ws.Request().Header.Set("Content-Type", "application/json")
	ws.Request().Header.Set("Version", internal.PanelVersion)
	ws.Request().Header.Set("Date", time.Now().Format(time.RFC3339Nano))

	if err := websocket.JSON.Send(ws, d); err != nil {
		m.Logger().Error(err)
	}

	return
}

func (m *Methods) badResponse(ws *websocket.Conn, err error) {
	r := structures.Response{
		Message: err.Error(),
		Data:    []string{},
	}
	m.Logger().Error(err)

	m.successResponse(ws, r)
	return
}

func NewMethods(s *internal.PanelSettings, l *logrus.Logger) *Methods {
	return &Methods{Settings: s, logger: l}
}

func (m *Methods) GetRoutes() *echo.Echo {
	e := echo.New()

	e.Static("/", "./frontend/")
	e.Router().Add(http.MethodGet, "/", func(c echo.Context) error {
		return c.File("./frontend/index.html")
	})

	e.Router().Add(http.MethodGet, "/ws/settings", m.getSettings)
	e.Router().Add(http.MethodPut, "/ws/settings", m.updateSettings)
	e.Router().Add(http.MethodGet, "/ws/dashboard", m.GetDashboardInfo)
	e.Router().Add(http.MethodGet, "/ws/disk_partitions", m.GetDiskPartitions)

	return e
}

func (m *Methods) GetDashboardInfo(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		var dashboard structures.DashboardParams
		err := websocket.JSON.Receive(ws, &dashboard)
		if err != nil {
			if _, err := ws.Write([]byte("Invalid request")); err != nil {
				m.Logger().Error(err)
			}
		}

		for {
			cpuLoad, err := internal.GetCPULoad()
			io, err := internal.GetDiskInfo(dashboard.Path)
			mem, err := internal.GetVirtualMemory()

			resp := structures.Dashboard{
				CPULoad: cpuLoad,
				Mem:     mem,
				IO:      io,
			}

			if err != nil {
				m.badResponse(ws, err)
				m.Logger().Error(err)
				return
			}

			m.successResponse(ws, resp)

			time.Sleep(time.Second)
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func (m *Methods) GetDiskPartitions(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		DiskPartitions, err := internal.GetDiskPartitions()
		if err != nil {
			m.badResponse(ws, err)
			m.Logger().Error(err)
			return
		}

		m.successResponse(ws, DiskPartitions)

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func (m *Methods) getSettings(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		settings, err := m.Settings.GetSettings()
		if err != nil {
			m.badResponse(ws, err)
			m.Logger().Error(err)
			return
		}

		m.successResponse(ws, settings)
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func (m *Methods) updateSettings(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		var settings internal.PanelSettings

		body, _ := ioutil.ReadAll(c.Request().Body)
		err := json.Unmarshal(body, &settings)
		if err != nil {
			m.badResponse(ws, err)
			m.Logger().Error(err)
			return
		}

		err = m.Settings.UpdateSettings(settings)
		if err != nil {
			m.badResponse(ws, err)
			m.Logger().Error(err)
			return
		}

		m.successResponse(ws, settings)

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
