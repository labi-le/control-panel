package api

import (
	"errors"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/labi-le/control-panel/pkg/utils"
	"golang.org/x/net/websocket"
	"syscall"
)

type Methods struct {
	Settings *internal.PanelSettings
}

func NewMethods(s *internal.PanelSettings) *Methods {
	return &Methods{Settings: s}
}

func (m *Methods) successResponseWS(ws *websocket.Conn, d ...any) bool {
	err := websocket.JSON.Send(ws, d)
	if err != nil {
		if errors.Is(err, syscall.EPIPE) {
			utils.Log().Infof("Client disconnected %s", ws.Request().RemoteAddr)
			return false
		}
		utils.Log().Error(err.Error())

		return false

	}

	return true
}

func (m *Methods) badResponseWS(ws *websocket.Conn, err error) bool {
	r := structures.Response{
		Message: err.Error(),
		Data:    []string{},
	}
	utils.Log().Error(err.Error())

	return m.successResponseWS(ws, r)
}
