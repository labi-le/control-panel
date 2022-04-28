package api

import (
	"fmt"
	"github.com/labi-le/control-panel/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router *echo.Echo
	logger *logrus.Logger

	PanelSettings *internal.PanelSettings
	*http.Server
}

func (s *Server) Logger() *logrus.Logger {
	return s.logger
}

func NewServer(router *echo.Echo, config *internal.PanelSettings) *Server {
	return &Server{router: router, logger: logrus.New(), PanelSettings: config}
}

func (s *Server) ListenAndServe() error {
	s.configureLogger()
	s.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status}\n",
	}))

	s.logger.Log(logrus.InfoLevel, "Server configuration:\n", s.PanelSettings.String())
	s.logger.Log(logrus.InfoLevel, "Rest api started")

	s.Server = &http.Server{
		Handler: s,
		Addr:    fmt.Sprintf("%s:%s", s.PanelSettings.Addr, s.PanelSettings.Port),
	}

	return s.Server.ListenAndServe()
}

// implement
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureLogger() {
	level, err := logrus.ParseLevel(s.PanelSettings.LogLevel)
	if err != nil {
		panic("invalid log level")
	}

	s.logger.SetLevel(level)
}
