package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/labi-le/control-panel/internal"
	router "github.com/labi-le/control-panel/internal/http"
	"github.com/labi-le/control-panel/internal/http/api"
	"github.com/labi-le/control-panel/pkg"
	"github.com/labi-le/control-panel/pkg/utils"
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
		utils.Log().Fatal(err)
	}

	utils.ConfigureLogger(conf.GetLogLevel())
	// Configure global logger

	resolver := api.NewMethods(conf)
	srv := pkg.NewServer(router.GetRoutes(resolver), conf)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Log().Fatal(err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (kill -2)
	<-stop
	utils.Log().Info("Gracefully shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Log().Fatal(err)
	}

}
