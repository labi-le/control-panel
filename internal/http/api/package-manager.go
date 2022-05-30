package api

import (
	"errors"
	"github.com/ShinyTrinkets/overseer"
	utils2 "github.com/labi-le/control-panel/pkg/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"syscall"
)

func (m *Methods) UpdatePackage(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		utils2.Log().Infof("Client connected %s", ws.Request().RemoteAddr)

		opt := overseer.Options{
			Buffered:  false,
			Streaming: true,
		}

		over := overseer.NewOverseer()
		cmd := over.Add("pacman update", "pacman", []string{
			"-Syu",
			"--noconfirm",
		}, opt)

		err := utils2.ManageProc(cmd, over, ws)
		if err != nil {
			utils2.Log().Info(cmd.Stop().Error())
			if errors.Is(err, syscall.EPIPE) {
				utils2.Log().Infof("Client disconnected %s", ws.Request().RemoteAddr)
				return
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
