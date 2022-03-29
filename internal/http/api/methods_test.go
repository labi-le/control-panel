package api

import (
	"bytes"
	"encoding/json"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/shirou/gopsutil/v3/disk"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var db *internal.DB
var config *structures.Config
var m *Methods

//goland:noinspection GoUnhandledErrorResult
func init() {
	_, err := gorm.Open(sqlite.Open("test"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	config = &structures.Config{
		Dsn: "test",
	}

	db = internal.NewDB(config)

	db.Migrate()
	m = NewMethods(db)

	defer os.Remove("test")
}

func TestMethods_GetRoutes(t *testing.T) {
	if (m.GetRoutes()) == nil {
		t.Fatal("GetRoutes() returned nil")
	}
}

func TestMethods_GetSettings(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/api/settings", nil)
	w := httptest.NewRecorder()

	m.ApiSettingsResolver(w, r)

	if w.Code != http.StatusOK {
		t.Fatal("ApiSettingsResolver() returned wrong status code")
	}

	type testResp struct {
		structures.Response
		Data structures.PanelSettings `json:"data"`
	}

	response := testResp{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("ApiSettingsResolver() returned wrong body")
	}

	if len(response.Data.Port) == 0 || len(response.Data.Language) == 0 || len(response.Data.Theme) == 0 {
		t.Fatal("ApiSettingsResolver() returned wrong body")
	}
}
func TestMethods_UpdateSettings(t *testing.T) {
	ps := structures.PanelSettings{
		Port:     "7777",
		Language: "jp",
		Theme:    "dark",
	}

	body, _ := json.Marshal(ps)

	r := httptest.NewRequest(http.MethodPut, "/api/settings", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	m.ApiSettingsResolver(w, r)

	if w.Code != http.StatusOK {
		t.Fatal("ApiSettingsResolver() returned wrong status code")
	}

	type testResp struct {
		structures.Response
		Data structures.PanelSettings `json:"data"`
	}

	response := testResp{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("ApiSettingsResolver() returned wrong body")
	}

	if response.Data.Theme != ps.Theme || response.Data.Language != ps.Language || response.Data.Port != ps.Port {
		t.Fatal("ApiSettingsResolver() returned wrong body")
	}
}

func TestMethods_GetDashboardInfo(t *testing.T) {
	param := structures.DashboardParams{
		Path: "/",
	}

	body, _ := json.Marshal(param)

	r := httptest.NewRequest(http.MethodPost, "/api/dashboard", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	m.GetDashboardInfo(w, r)

	if w.Code != http.StatusOK {
		t.Fatal("GetDashboardInfo() returned wrong status code")
	}

	type testResp struct {
		structures.Response
		Data structures.Dashboard `json:"data"`
	}

	response := testResp{}
	if json.Unmarshal(w.Body.Bytes(), &response) != nil {
		t.Fatal("GetDashboardInfo() returned wrong body")
	}

	if response.Data.Mem == nil || response.Data.CPULoad == nil || response.Data.IO == nil {
		t.Fatal("GetDashboardInfo() returned wrong body")
	}
}

func TestMethods_GetDiskPartitions(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/api/disk_partitions", nil)
	w := httptest.NewRecorder()

	m.GetDiskPartitions(w, r)

	if w.Code != http.StatusOK {
		t.Fatal("GetDiskPartitions() returned wrong status code")
	}

	type testResp struct {
		structures.Response
		Data []disk.PartitionStat `json:"data"`
	}

	response := testResp{}
	if json.Unmarshal(w.Body.Bytes(), &response) != nil {
		t.Fatal("GetDiskPartitions() returned wrong body")
	}

	if response.Data == nil {
		t.Fatal("GetDiskPartitions() returned wrong body")
	}
}
