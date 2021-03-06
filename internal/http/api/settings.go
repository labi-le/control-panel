package api

import (
	"encoding/json"
	"github.com/labi-le/control-panel/internal"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func (m *Methods) GetSettings(c echo.Context) error {
	settings, err := m.Settings.GetSettings()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, settings)
}

func (m *Methods) UpdateSettings(c echo.Context) error {
	s := internal.PanelSettings{}

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &s)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = m.Settings.UpdateSettings(s)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, s)
}

func (m *Methods) ResetSettings(c echo.Context) error {
	err := m.Settings.ResetSettings()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, internal.DefaultPanelSettings())
}
