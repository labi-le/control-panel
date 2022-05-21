package http

import (
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/http/api"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetRoutes(m *api.Methods) *echo.Echo {
	e := echo.New()

	e.Static("/", "./frontend/")
	e.Router().Add(http.MethodGet, "/", func(c echo.Context) error {
		c.Response().Header().Set("Version", internal.PanelVersion)
		c.Response().Header().Set("Date", time.Now().Format(time.RFC3339Nano))

		return c.File("./frontend/index.html")
	})

	e.Router().Add(http.MethodGet, "/ws/dashboard", m.GetDashboardInfo)

	e.Router().Add(http.MethodGet, "/api/settings", m.GetSettings)
	e.Router().Add(http.MethodPut, "/api/settings", m.UpdateSettings)
	e.Router().Add(http.MethodGet, "/api/disk_partitions", m.GetDiskPartitions)
	e.Router().Add(http.MethodGet, "/api/version", m.GetVersion)

	return e
}
