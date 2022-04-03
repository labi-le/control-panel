package structures

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

const (
	DefaultConfigPath = "~/.config/control-panel/"
	DefaultConfigName = "config.toml"
	DefaultConfig     = DefaultConfigPath + DefaultConfigName
	PanelVersion      = "0.0.1"
)

type Config struct {
	Dsn      string `toml:"dsn"`
	LogLevel string `toml:"log_level"`
	Addr     string `toml:"addr"`
	Port     string `toml:"port"`

	ConfigPath string
}

func NewConfig() *Config {
	return &Config{}
}

// GetDefaultConfig returns default config
// ErrorLevels:
// panic
// fatal
// error
// warning
// info
// debug
// trace
func GetDefaultConfig() *Config {
	return &Config{
		Dsn:      "settings",
		LogLevel: "debug",
		Addr:     "0.0.0.0",
		Port:     "7777",
	}
}

// SaveConfig saves config to file
func (c *Config) SaveConfig(config *Config, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(path + DefaultConfigName)
	if err != nil {
		return err
	}

	err = toml.NewEncoder(f).Encode(config)
	if err != nil {
		return err
	}

	return nil
}

// LoadConfig loads config from file
func (c *Config) LoadConfig(configPath string) (*Config, error) {
	c.ConfigPath = configPath
	// Check if config file exists from default path and if not, generate default config and load it
	if configPath == DefaultConfig && c.ConfigExists(configPath) == false {
		fmt.Println("Config file not found. Generate default config file...")
		err := c.SaveConfig(GetDefaultConfig(), DefaultConfigPath)
		if err != nil {
			return c.LoadConfig(configPath)
		}
	}

	if c.ConfigExists(configPath) == false {
		fmt.Printf("Config file not found in %s", configPath)
		os.Exit(1)
	}

	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		return c, err
	}
	return c, nil
}

// ConfigExists checks if config file exists
func (c *Config) ConfigExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// String returns config as string
func (c *Config) String() string {
	return fmt.Sprintf(
		"dsn: %s\nlog level: %s\naddr: %s\nport: %s\nconfig path: %s",
		c.Dsn,
		c.LogLevel,
		c.Addr,
		c.Port,
		c.ConfigPath,
	)
}
