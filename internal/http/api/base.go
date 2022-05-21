package api

import (
	"errors"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"syscall"
)

type Methods struct {
	Settings *internal.PanelSettings
	logger   *logrus.Logger
}

func NewMethods(s *internal.PanelSettings, l *logrus.Logger) *Methods {
	return &Methods{Settings: s, logger: l}
}

func (m *Methods) Logger() *logrus.Logger {
	return m.logger
}

func (m *Methods) successResponseWS(ws *websocket.Conn, d ...any) bool {
	err := websocket.JSON.Send(ws, d)
	if err != nil {
		if errors.Is(err, syscall.EPIPE) {
			m.Logger().Infof("Client disconnected %s", ws.Request().RemoteAddr)
			return false
		}
		m.Logger().Error(err)

		return false

	}

	return true
}

func (m *Methods) badResponseWS(ws *websocket.Conn, err error) bool {
	r := structures.Response{
		Message: err.Error(),
		Data:    []string{},
	}
	m.Logger().Error(err)

	return m.successResponseWS(ws, r)
}
