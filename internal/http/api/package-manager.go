package api

import (
	"errors"
	"github.com/ShinyTrinkets/overseer"
	"github.com/labi-le/control-panel/pkg"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"syscall"
)

func (m *Methods) UpdatePackage(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		m.Logger().Infof("Client connected %s", ws.Request().RemoteAddr)

		opt := overseer.Options{
			Buffered:  false,
			Streaming: true,
		}

		over := overseer.NewOverseer()
		cmd := over.Add("pacman update", "pacman", []string{
			"-Syu",
			"--noconfirm",
		}, opt)

		proc := pkg.NewProcess(m.Logger())
		err := proc.ManageProc(cmd, over, ws)
		if err != nil {
			m.Logger().Info(cmd.Stop())
			if errors.Is(err, syscall.EPIPE) {
				m.Logger().Infof("Client disconnected %s", ws.Request().RemoteAddr)
				return
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
