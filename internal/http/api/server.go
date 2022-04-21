package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labi-le/control-panel/internal"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger

	PanelSettings *internal.PanelSettings
}

func NewServer(router *mux.Router, config *internal.PanelSettings) *Server {
	return &Server{router: router, logger: logrus.New(), PanelSettings: config}
}

func (s *Server) Start() error {
	s.configureLogger()
	s.router.Use(s.logRequestMiddleware)

	s.logger.Log(logrus.InfoLevel, "Server configuration:\n", s.PanelSettings.String())
	s.logger.Log(logrus.InfoLevel, "Rest api started")

	server := &http.Server{
		Handler: s,
		Addr:    fmt.Sprintf("%s:%s", s.PanelSettings.Addr, s.PanelSettings.Port),
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	c.Handler(s)

	return server.ListenAndServe()
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

func (s *Server) logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"IP": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}
