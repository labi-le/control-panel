package pkg

import "time"

type ConfigInterface interface {
	GetAddr() string
	GetPort() string
	GetLogLevel() string
	GetLanguage() string
	GetDashboardUpdateTimeout() time.Duration
}
