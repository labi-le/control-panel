package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/structures"
	"github.com/sirupsen/logrus"
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

	// api info
	s.router.HandleFunc("/api/{method}", s.apiInfoResolver).Methods(http.MethodPost)

	// api put data
	s.router.HandleFunc("/api/{method}", s.apiChangeResolver).Methods(http.MethodPut)

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
			time.Now().Sub(start),
		)
	})
}

func (s *Server) apiInfoResolver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//var Item http.Response
	//body, _ := ioutil.ReadAll(r.Body)

	params := mux.Vars(r)
	method := params["method"]

	resp := structures.Response{
		Version: "0.1",
		Time:    time.Now(),
	}

	switch method {
	case "settings":
		Settings, err := s.DB.GetSettings()
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()

			response(resp, w)
			break
		}

		resp.Success = true
		resp.Message = "Settings has been retrieved"
		resp.Data = Settings

		response(resp, w)
		break

	case "mem":
		Mem, err := internal.GetVirtualMemory()
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()

			response(resp, w)
			break
		}

		resp.Success = true
		resp.Message = "Mem has been retrieved"
		resp.Data = Mem

		response(resp, w)
		break

	default:
		resp.Success = false
		resp.Message = "Method not found"

		response(resp, w)
		break
	}

}

func (s *Server) apiChangeResolver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//var Item http.Response
	//body, _ := ioutil.ReadAll(r.Body)

	params := mux.Vars(r)
	method := params["method"]

	resp := structures.Response{
		Version: "0.1",
		Time:    time.Now(),
	}

	switch method {
	case "settings":
		//todo
		break
	default:
		resp.Success = false
		resp.Message = "Method not found"

		response(resp, w)
		break
	}

}

func response(response structures.Response, w http.ResponseWriter) {
	switch response.Success {
	case false:
		w.WriteHeader(http.StatusBadRequest)
		break

	case true:
		w.WriteHeader(http.StatusOK)
		break
	}

	_ = json.NewEncoder(w).Encode(response)
	return
}
