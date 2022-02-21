package structures

import "github.com/labi-le/control-panel/internal"

type Config struct {
	Dsn      string `toml:"dsn"`
	LogLevel string `toml:"log_level"`
	Addr     string `toml:"addr"`
	Port     int    `toml:"port"`

	DB *internal.DB
}
