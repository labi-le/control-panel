package internal

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var (
	DefaultConfigPath   = os.Getenv("HOME") + "/.config/control-panel/"
	DefaultSettingsFile = "settings"
)

const (
	PanelVersion = "0.0.1"

	DefaultLanguage = "en"
	DefaultLogLevel = "debug"
	DefaultAddr     = "0.0.0.0"
	DefaultPort     = "7777"
)

// PanelSettings
// Port - the port on which the panel will work
// LogLevel - the debug level of the panel
// Language - panel language
// Theme - panel theme
type PanelSettings struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	LogLevel string `json:"debug_level"`
	Language string `json:"language"`

	dbConn *gorm.DB
}

func NewPanelSettings(settingsPath string) (*PanelSettings, error) {
	if settingsPath == "" {
		settingsPath = DefaultConfigPath
	}

	if _, err := os.Stat(settingsPath + DefaultSettingsFile); os.IsNotExist(err) {
		fmt.Printf("Settings does not exist, creating in %s...\n", settingsPath)
		err := os.MkdirAll(settingsPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	conn, err := gorm.Open(sqlite.Open(settingsPath+DefaultSettingsFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = conn.AutoMigrate(DefaultPanelSettings()); err != nil {
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
		Addr:     DefaultAddr,
		Port:     DefaultPort,
		LogLevel: DefaultLogLevel,
		Language: DefaultLanguage,
	}
}

// GetSettings returns the settings
// if the settings are not found, it will save and return default settings
func (p *PanelSettings) GetSettings() (*PanelSettings, error) {
	var settings *PanelSettings
	if err := p.dbConn.FirstOrCreate(settings).Error; err != nil {
		settings = DefaultPanelSettings()
	}

	return settings, nil
}

// UpdateSettings updates the settings
func (p *PanelSettings) UpdateSettings(settings PanelSettings) error {
	return p.dbConn.Model(&settings).Where("_rowid_ = ?", 1).Updates(&settings).Error
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
