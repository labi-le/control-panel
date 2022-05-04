package pkg

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router *echo.Echo
	logger *logrus.Logger

	PanelSettings ConfigInterface
	*http.Server
}

func (s *Server) Logger() *logrus.Logger {
	return s.logger
}

func NewServer(router *echo.Echo, config ConfigInterface) *Server {
	srv := &Server{router: router, logger: logrus.New(), PanelSettings: config}
	srv.configureLogger()

	return srv
}

func (s *Server) ListenAndServe() error {
	s.router.Use(s.logMiddleware)

	// todo add https support
	//goland:noinspection HttpUrlsUsage
	s.Logger().Infof("Panel is available at http://%s:%s", s.PanelSettings.GetAddr(), s.PanelSettings.GetPort())
	s.Logger().Log(logrus.InfoLevel, "Rest api started")

	s.Server = &http.Server{
		Handler: s,
		Addr:    fmt.Sprintf("%s:%s", s.PanelSettings.GetAddr(), s.PanelSettings.GetPort()),
	}

	return s.Server.ListenAndServe()
}

func (s *Server) logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		s.Logger().Info(fmt.Sprintf("%s %s %d", c.Request().Method, c.Request().URL.Path, c.Response().Status))
		return next(c)
	}
}

// implement
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureLogger() {
	level, err := logrus.ParseLevel(s.PanelSettings.GetLogLevel())
	if err != nil {
		panic("invalid log level")
	}

	s.logger.SetLevel(level)
}
