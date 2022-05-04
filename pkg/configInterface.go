package pkg

type ConfigInterface interface {
	GetAddr() string
	GetPort() string
	GetLogLevel() string
	GetLanguage() string
}
