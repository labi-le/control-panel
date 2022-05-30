package pkg

import (
	"fmt"
	"github.com/labi-le/control-panel/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router        *echo.Echo
	PanelSettings ConfigInterface
	*http.Server
}

func NewServer(router *echo.Echo, config ConfigInterface) *Server {
	srv := &Server{router: router, PanelSettings: config}

	return srv
}

func (s *Server) ListenAndServe() error {
	s.router.Use(s.logMiddleware)

	// todo add https support
	//goland:noinspection HttpUrlsUsage
	utils.Log().Infof("Panel is available at http://%s:%s", s.PanelSettings.GetAddr(), s.PanelSettings.GetPort())
	utils.Log().Log(logrus.InfoLevel, "Rest api started")

	s.Server = &http.Server{
		Handler: s,
		Addr:    fmt.Sprintf("%s:%s", s.PanelSettings.GetAddr(), s.PanelSettings.GetPort()),
	}

	return s.Server.ListenAndServe()
}

func (s *Server) logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		utils.Log().Info(fmt.Sprintf("%s %s %d", c.Request().Method, c.Request().URL.Path, c.Response().Status))
		return next(c)
	}
}

// implement
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
