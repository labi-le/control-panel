package api

import (
	"github.com/labi-le/control-panel/internal"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Methods) GetVersion(c echo.Context) error {
	return c.String(http.StatusOK, internal.PanelVersion)
}
