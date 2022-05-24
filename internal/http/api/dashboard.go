package api

import (
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"time"
)

func (m *Methods) GetDashboardInfo(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		m.Logger().Infof("Client connected %s", ws.Request().RemoteAddr)

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
				m.badResponseWS(ws, err)
				break
			}

			if m.successResponseWS(ws, resp) == false {
				break
			}

			time.Sleep(m.Settings.DashboardUpdateTimeout)
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
