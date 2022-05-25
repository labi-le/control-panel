package internal

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

var (
	DefaultConfigPath   = os.Getenv("HOME") + "/.config/control-panel/"
	DefaultSettingsFile = "settings"
)

const (
	PanelVersion = "0.0.1"

	DefaultLanguage               = "en"
	DefaultLogLevel               = "debug"
	DefaultAddr                   = "0.0.0.0"
	DefaultPort                   = "7777"
	DefaultDashboardUpdateTimeout = time.Second * 1
)

// PanelSettings
// port - the port on which the panel will work
// logLevel - the debug level of the panel
// language - panel language
// Theme - panel theme
type PanelSettings struct {
	Addr                   string        `json:"addr"`
	Port                   string        `json:"port"`
	LogLevel               string        `json:"log_level"`
	Language               string        `json:"language"`
	DashboardUpdateTimeout time.Duration `json:"dashboard_update_timeout"`

	dbConn *gorm.DB
}

func (p *PanelSettings) GetAddr() string {
	return p.Addr
}

func (p *PanelSettings) GetPort() string {
	return p.Port
}

func (p *PanelSettings) GetLogLevel() string {
	return p.LogLevel
}

func (p *PanelSettings) GetLanguage() string {
	return p.Language
}

func (p *PanelSettings) GetDashboardUpdateTimeout() time.Duration {
	return p.DashboardUpdateTimeout
}

func NewPanelSettings(settingsPath string) (*PanelSettings, error) {
	if settingsPath == "" {
		settingsPath = DefaultConfigPath
	}

	if _, err := os.Stat(settingsPath + DefaultSettingsFile); os.IsNotExist(err) {
		err := os.MkdirAll(settingsPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	conn, err := gorm.Open(
		sqlite.Open(settingsPath+DefaultSettingsFile),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
	)
	if err != nil {
		return nil, err
	}
	if err = conn.AutoMigrate(&PanelSettings{}); err != nil {
		return nil, err
	}

	return (&PanelSettings{dbConn: conn}).GetSettings()
}

// DefaultPanelSettings
// ErrorLevels:
// panic
// fatal
// error
// warning
// info
// debug
// trace
func DefaultPanelSettings() *PanelSettings {
	return &PanelSettings{
		Addr:                   DefaultAddr,
		Port:                   DefaultPort,
		LogLevel:               DefaultLogLevel,
		Language:               DefaultLanguage,
		DashboardUpdateTimeout: DefaultDashboardUpdateTimeout,
	}
}

// GetSettings returns the settings
// if the settings are not found, it will save and return default settings
func (p *PanelSettings) GetSettings() (*PanelSettings, error) {
	if err := p.dbConn.Where("_rowid_ = ?", 1).First(p).Error; err != nil {
		if err = p.dbConn.Create(DefaultPanelSettings()).Error; err != nil {
			return nil, err
		} else {
			pp := DefaultPanelSettings()
			pp.dbConn = p.dbConn

			return pp, nil
		}
	}

	return p, nil
}

// UpdateSettings updates the settings
func (p *PanelSettings) UpdateSettings(settings PanelSettings) error {
	return p.dbConn.Model(&settings).Where("_rowid_ = ?", 1).Updates(&settings).Error
}

func (p *PanelSettings) ResetSettings() error {
	return p.UpdateSettings(*DefaultPanelSettings())
}

// String returns config as string
func (p *PanelSettings) String() string {
	return fmt.Sprintf(
		"lang: %s\nlog level: %s\naddr: %s\nport: %s",
		p.Language,
		p.LogLevel,
		p.Addr,
		p.Port,
	)
}
