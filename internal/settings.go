package internal

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"time"
)

const (
	ProductionStaticPath = "/opt/control-panel/static/"
	DevelopStaticPath    = "./frontend/"
)

type Config interface {
	GetAddr() string
	GetPort() int
	GetLogLevel() string
	GetLanguage() string
	GetDashboardUpdateTimeout() time.Duration
}

// PanelSettings
// port - the port on which the panel will work
// logLevel - the debug level of the panel
// language - panel language
// DashboardDelay - the timeout of updates dashboard
type PanelSettings struct {
	Addr           string        `toml:"addr"`
	Port           int           `toml:"port"`
	LogLevel       string        `toml:"log_level"`
	Language       string        `toml:"language"`
	DashboardDelay time.Duration `toml:"dashboard_delay"`
}

func (p *PanelSettings) GetAddr() string {
	return p.Addr
}

func (p *PanelSettings) GetPort() int {
	return p.Port
}

func (p *PanelSettings) GetLogLevel() string {
	return p.LogLevel
}

func (p *PanelSettings) GetLanguage() string {
	return p.Language
}

func (p *PanelSettings) GetDashboardUpdateTimeout() time.Duration {
	return p.DashboardDelay
}

func NewPanelSettings(settingsPath string) (*PanelSettings, error) {
	if settingsPath == "" {
		settingsPath = DefaultConfigPath()
	}

	var (
		file     io.Reader
		settings PanelSettings
	)

	log.Debug().Msgf("Settings path: %s", settingsPath)
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		file, err = createConfigFile(settingsPath, DefaultPanelSettings())
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Open(settingsPath)
	if err != nil {
		return nil, err
	}

	if _, err := toml.NewDecoder(file).Decode(&settings); err != nil {
		return nil, err
	}
	return &settings, err
}

func createConfigFile(path string, settings *PanelSettings) (io.Reader, error) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return file, toml.NewEncoder(file).Encode(settings)

}

func DefaultPanelSettings() *PanelSettings {
	return &PanelSettings{
		Addr:           "0.0.0.0",
		Port:           7777,
		LogLevel:       "info",
		Language:       "en",
		DashboardDelay: time.Second,
	}
}

// String returns config as string
func (p *PanelSettings) String() string {
	return fmt.Sprintf(
		"lang: %s\nlog level: %s\naddr: %s\nport: %d",
		p.Language,
		p.LogLevel,
		p.Addr,
		p.Port,
	)
}

func DefaultConfigPath() string {
	return os.Getenv("HOME") + "/.config/control-panel/config.toml"
}
