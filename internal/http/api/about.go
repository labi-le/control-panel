package api

import (
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Methods) GetVersion(c echo.Context) error {
	return c.JSON(http.StatusOK, structures.Version{V: internal.PanelVersion})
}
