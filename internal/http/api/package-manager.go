package api

import (
	"errors"
	"github.com/ShinyTrinkets/overseer"
	"github.com/labi-le/control-panel/internal/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"syscall"
)

func (m *Methods) UpdatePackage(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		utils.Log().Infof("Client connected %s", ws.Request().RemoteAddr)

		opt := overseer.Options{
			Buffered:  false,
			Streaming: true,
		}

		over := overseer.NewOverseer()
		cmd := over.Add("pacman update", "pacman", []string{
			"-Syu",
			"--noconfirm",
		}, opt)

		err := utils.ManageProc(cmd, over, ws)
		if err != nil {
			utils.Log().Info(cmd.Stop().Error())
			if errors.Is(err, syscall.EPIPE) {
				utils.Log().Infof("Client disconnected %s", ws.Request().RemoteAddr)
				return
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
