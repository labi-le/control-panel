package structures

type Config struct {
	Dsn      string `toml:"dsn"`
	LogLevel string `toml:"log_level"`
	Addr     string `toml:"addr"`
	Port     string `toml:"port"`
}
