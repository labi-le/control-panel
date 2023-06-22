package internal

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/labi-le/control-panel/internal/types"
	"github.com/labi-le/control-panel/pkg/log"
	"github.com/labi-le/control-panel/pkg/response"
	"github.com/labi-le/control-panel/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"syscall"
	"time"
)

type API struct {
	reply   *response.Reply
	service *PanelSettings
}

func RegisterHandlers(
	r fiber.Router,
	service *PanelSettings,
) {

	if utils.IsDirExist(ProductionStaticPath) {
		r.Static("/", ProductionStaticPath)
	} else {
		r.Static("/", DevelopStaticPath)
	}

	api := &API{
		reply:   response.New(log.GlobalLog),
		service: service,
	}

	r.Add(http.MethodGet, "/ws/dashboard", websocket.New(api.GetDashboardInfo))
	r.Add(http.MethodGet, "/ws/package/update", api.UpdatePackage)
	r.Add(http.MethodGet, "/ws/package/install/:package", api.InstallPackage)
	r.Add(http.MethodGet, "/ws/package/remove/:package", api.DeletePackage)

	r.Add(http.MethodGet, "/api/settings", api.GetSettings)
	r.Add(http.MethodPut, "/api/settings", api.UpdateSettings)
	r.Add(http.MethodGet, "/api/settings/reset", api.ResetSettings)
	r.Add(http.MethodGet, "/api/disk_partitions", api.GetDiskPartitions)
	r.Add(http.MethodGet, "/api/version", api.GetVersion)
}

func (a *API) successResponseWS(ws *websocket.Conn, d ...any) bool {
	if err := ws.WriteJSON(d); err != nil {
		if errors.Is(err, syscall.EPIPE) {
			log.Infof("Client disconnected %s", ws.RemoteAddr)
			return false
		}
		log.Error(err)

		return false
	}

	return true
}

func (a *API) badResponseWS(ws *websocket.Conn, err error) bool {
	r := types.Response{
		Message: err.Error(),
		Data:    []string{},
	}
	log.Error(err)

	return a.successResponseWS(ws, r)
}

func (a *API) GetVersion(c *fiber.Ctx) error {
	return a.reply.OK(c, types.Version{V: PanelVersion})
}

func (a *API) GetDashboardInfo(ws *websocket.Conn) {
	defer ws.Close()

	log.Infof("Client connected %s", ws.RemoteAddr)

	var dashboard types.DashboardParams
	if err := ws.ReadJSON(&dashboard); err != nil {
		log.Error(err)
		return
	}

	for {
		cpuLoad, err := GetCPULoad()
		io, err := GetDiskInfo(dashboard.Path)
		mem, err := GetVirtualMemory()

		resp := types.Dashboard{
			CPULoad: cpuLoad,
			Mem:     mem,
			IO:      io,
		}

		if err != nil {
			a.badResponseWS(ws, err)
			break
		}

		if a.successResponseWS(ws, resp) == false {
			break
		}

		time.Sleep(a.service.DashboardUpdateTimeout)
	}
}

func (a *API) GetDiskPartitions(c *fiber.Ctx) error {
	dp, err := GetDiskPartitions()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return a.reply.OK(c, dp)
}

func (a *API) GetSettings(c *fiber.Ctx) error {
	settings, err := a.service.GetSettings()
	if err != nil {
		return a.reply.InternalServerError(c, err)
	}

	return a.reply.OK(c, settings)
}

func (a *API) UpdateSettings(c *fiber.Ctx) error {
	s := PanelSettings{}

	err := json.Unmarshal(c.Body(), &s)
	if err != nil {
		return a.reply.BadRequest(c, err)
	}

	err = a.service.UpdateSettings(s)
	if err != nil {
		return a.reply.InternalServerError(c, err)
	}

	return a.reply.OK(c, s)
}

func (a *API) ResetSettings(c *fiber.Ctx) error {
	err := a.service.ResetSettings()
	if err != nil {
		return a.reply.InternalServerError(c, err)
	}

	return a.reply.OK(c, DefaultPanelSettings())
}

func (a *API) UpdatePackage(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (a *API) InstallPackage(ctx *fiber.Ctx) error {
	panic("implement me")

}

func (a *API) DeletePackage(ctx *fiber.Ctx) error {
	panic("implement me")
}
