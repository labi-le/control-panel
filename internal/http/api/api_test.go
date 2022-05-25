package api

import (
	"bytes"
	"encoding/json"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var m *Methods

func init() {
	settings, err := internal.NewPanelSettings("./")
	if err != nil {
		panic(err)
	}

	m = NewMethods(settings, logrus.New())
}

func TestMethods_GetVersion(t *testing.T) {
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/version", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(r, rec)

	if err := m.GetVersion(c); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var v structures.Version
	if err := json.Unmarshal(rec.Body.Bytes(), &v); err != nil {
		t.Error(err)
	}

	assert.Equal(t, internal.PanelVersion, v.V)
}

func TestMethods_GetDiskPartitions(t *testing.T) {
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/disk_partitions", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(r, rec)

	if err := m.GetDiskPartitions(c); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var p []structures.PartitionStat
	if err := json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Error(err)
	}

	for _, stat := range p {
		assert.NotEmptyf(t, stat.Device, "Device is empty")
		assert.NotEmptyf(t, stat.Mountpoint, "Mountpoint is empty")
		assert.NotEmptyf(t, stat.Fstype, "Fstype is empty")
		assert.NotEmptyf(t, stat.Opts, "Opts is empty")
	}
}

func TestMethods_GetSettings(t *testing.T) {
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/settings", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(r, rec)

	if err := m.GetSettings(c); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var p internal.PanelSettings
	if err := json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, p.GetDashboardUpdateTimeout())
	assert.NotEmpty(t, p.GetLogLevel())
	assert.NotEmpty(t, p.GetLanguage())
	assert.NotEmpty(t, p.GetPort())
	assert.NotEmpty(t, p.GetAddr())

}

func TestMethods_UpdateSettings(t *testing.T) {
	e := echo.New()

	body := &internal.PanelSettings{
		Addr:                   "0.0.0.0",
		Port:                   "9999",
		LogLevel:               "info",
		Language:               "jp",
		DashboardUpdateTimeout: 10 * time.Second,
	}

	b, _ := json.Marshal(body)

	r := httptest.NewRequest(http.MethodPost, "/api/settings", bytes.NewReader(b))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(r, rec)

	if err := m.UpdateSettings(c); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	resp := &internal.PanelSettings{}
	if err := json.Unmarshal(rec.Body.Bytes(), resp); err != nil {
		t.Error(err)
	}

	assert.Equal(t, body.Addr, resp.GetAddr())
	assert.Equal(t, body.Port, resp.GetPort())
	assert.Equal(t, body.LogLevel, resp.GetLogLevel())
	assert.Equal(t, body.Language, resp.GetLanguage())
	assert.Equal(t, body.DashboardUpdateTimeout, resp.GetDashboardUpdateTimeout())
}

func TestMethods_ResetSettings(t *testing.T) {
	defer os.Remove(internal.DefaultSettingsFile)

	e := echo.New()

	r := httptest.NewRequest(http.MethodGet, "/api/settings/reset", nil)
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(r, rec)

	if err := m.ResetSettings(c); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	resp := &internal.PanelSettings{}
	if err := json.Unmarshal(rec.Body.Bytes(), resp); err != nil {
		t.Error(err)
	}

	defaultSettings := internal.DefaultPanelSettings()

	assert.Equal(t, defaultSettings.GetAddr(), resp.GetAddr())
	assert.Equal(t, defaultSettings.GetPort(), resp.GetPort())
	assert.Equal(t, defaultSettings.GetLogLevel(), resp.GetLogLevel())
	assert.Equal(t, defaultSettings.GetLanguage(), resp.GetLanguage())
	assert.Equal(t, defaultSettings.GetDashboardUpdateTimeout(), resp.GetDashboardUpdateTimeout())
}
