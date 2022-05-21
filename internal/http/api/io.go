package api

import (
	"github.com/labi-le/control-panel/internal"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Methods) GetDiskPartitions(c echo.Context) error {
	dp, err := internal.GetDiskPartitions()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, dp)
}
