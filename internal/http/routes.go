package http

import (
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/http/api"
	"github.com/labi-le/control-panel/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetRoutes(m *api.Methods) *echo.Echo {
	e := echo.New()

	if utils.IsDirExist(internal.ProductionStaticPath) {
		e.Static("/", internal.ProductionStaticPath)
	} else {
		e.Static("/", internal.DevelopStaticPath)
	}

	e.Router().Add(http.MethodGet, "/ws/dashboard", m.GetDashboardInfo)
	e.Router().Add(http.MethodGet, "/ws/package/update", m.UpdatePackage)
	e.Router().Add(http.MethodGet, "/ws/package/install/:package", m.InstallPackage)
	e.Router().Add(http.MethodGet, "/ws/package/remove/:package", m.DeletePackage)

	e.Router().Add(http.MethodGet, "/api/settings", m.GetSettings)
	e.Router().Add(http.MethodPut, "/api/settings", m.UpdateSettings)
	e.Router().Add(http.MethodGet, "/api/settings/reset", m.ResetSettings)
	e.Router().Add(http.MethodGet, "/api/disk_partitions", m.GetDiskPartitions)
	e.Router().Add(http.MethodGet, "/api/version", m.GetVersion)

	return e
}
