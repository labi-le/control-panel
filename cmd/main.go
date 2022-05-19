package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/http/api"
	"github.com/labi-le/control-panel/pkg"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	config      string
	versionFlag bool
)

func init() {
	flag.StringVar(&config, "config", internal.DefaultConfigPath, "path to config file")
	flag.BoolVar(&versionFlag, "version", false, "print version and exit")
}

func main() {
	flag.Parse()

	if versionFlag == true {
		fmt.Println(internal.PanelVersion)
		return
	}

	conf, err := internal.NewPanelSettings(config)
	if err != nil {
		log.Fatal(err)
	}

	level, err := logrus.ParseLevel(conf.GetLogLevel())
	if err != nil {
		panic("invalid log level")
	}

	logger := logrus.New()
	logger.SetLevel(level)

	apiResolver := api.NewMethods(conf, logger)
	srv := pkg.NewServer(apiResolver.GetRoutes(), conf, logger)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.Logger().Fatal(err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (kill -2)
	<-stop
	srv.Logger().Info("Gracefully shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger().Fatal(err)
	}

}
