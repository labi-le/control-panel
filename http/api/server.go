package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/structures"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger

	Config *structures.Config
	DB     *internal.DB
}

// implement
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(srv *Server) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),

		DB:     internal.NewDB(srv.Config),
		Config: srv.Config,
	}

	s.route()

	return s
}

func Start(s *Server) error {
	srv := newServer(s)
	srv.configureLogger()

	srv.logger.Log(logrus.InfoLevel, "Rest api started")

	server := &http.Server{
		Handler: srv,
		Addr:    s.Config.Addr,
	}

	return server.ListenAndServe()
}

func (s *Server) configureLogger() {
	level, err := logrus.ParseLevel(s.Config.LogLevel)
	if err != nil {
		panic("invalid log level")
	}

	s.logger.SetLevel(level)
}

func (s *Server) route() {
	s.router.Use(s.logRequestMiddleware)

	// api put data
	s.router.HandleFunc("/api/settings", s.apiSettingsResolver).Methods(http.MethodPut, http.MethodPost)
	// api info
	s.router.HandleFunc("/api/{hardware}/{method}", s.hardwareInfoResolver).Methods(http.MethodPost)
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

func (s *Server) apiSettingsResolver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	method := NewMethods(w, s.DB)

	if r.Method == http.MethodPost {
		ResponseMethod(method.GetSettings())
		return
	} else if r.Method == http.MethodPut {
		var settings structures.PanelSettings

		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &settings)
		if err != nil {
			ResponseMethod(method.BadRequest(err))
			return
		}
		ResponseMethod(method.UpdateSettings(settings))
		return
	}

	ResponseMethod(method.MethodNotFound())
}

func (s *Server) hardwareInfoResolver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	hardware := params["hardware"]
	methodName := params["method"]

	method := NewMethods(w, s.DB)

	switch hardware {
	case "cpu":
		switch methodName {
		case "info":
			ResponseMethod(method.GetCPUInfo())

		case "load":
			ResponseMethod(method.GetCPUAvg())

		case "times":
			ResponseMethod(method.GetCPUTimes())
		}

	case "mem":
		ResponseMethod(method.GetVirtualMemory())

	default:
		ResponseMethod(method.MethodNotFound())
	}
}

func Response(response structures.Response, w http.ResponseWriter) {
	switch response.Success {
	case false:
		w.WriteHeader(http.StatusBadRequest)

	case true:
		w.WriteHeader(http.StatusOK)
	}

	_ = json.NewEncoder(w).Encode(response)
}

func ResponseMethod(m *Methods) {
	Response(m.resp, m.w)
}
