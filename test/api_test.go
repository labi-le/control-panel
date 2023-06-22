package test

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/types"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app = fiber.New()

func init() {
	internal.RegisterHandlers(app, internal.DefaultPanelSettings())
}

func TestMethods_GetVersion(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/version", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var v types.Version
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		t.Error(err)
	}

	assert.Equal(t, internal.Version, v.V)
}

func TestMethods_GetDiskPartitions(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/disk_partitions", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var p []types.PartitionStat
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		t.Error(err)
	}

	for _, stat := range p {
		assert.NotEmptyf(t, stat.Device, "Device is empty")
		assert.NotEmptyf(t, stat.Mountpoint, "Mountpoint is empty")
		assert.NotEmptyf(t, stat.Fstype, "Fstype is empty")
		assert.NotEmptyf(t, stat.Opts, "Opts is empty")
	}
}
