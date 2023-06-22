package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/labi-le/control-panel/internal"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

var (
	config      string
	versionFlag bool
	// Режим отладки всего приложения, sql запросы
	debugMode bool
)

func init() {
	flag.BoolVar(&debugMode, "debug", false, "debug mode")
	flag.StringVar(&config, "config", internal.DefaultConfigPath(), "path to config file")
	flag.BoolVar(&versionFlag, "version", false, "print version and exit")
}

func main() {
	flag.Parse()

	if versionFlag == true {
		log.Info().Msgf("Version: %s", internal.PanelVersion)
		return
	}

	conf, err := internal.NewPanelSettings(config)
	if err != nil {
		log.Fatal().Err(err)
	}

	checkPermissions()

	// Setting up signal capturing
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srv := fiber.New(fiber.Config{})
	internal.RegisterHandlers(srv, conf)

	go func() {
		err := srv.Listen(fmt.Sprintf("%s:%d", conf.GetAddr(), conf.GetPort()))
		if err != nil {
			log.Fatal().Err(err)
		}
	}()

	<-ctx.Done()
	log.Info().Msgf("Gracefully shutdown server...")

	if err := srv.Shutdown(); err != nil {
		log.Fatal().Err(err)
	}

}

var ErrPermissionDenied = errors.New("permission denied")

func checkPermissions() {
	if os.Geteuid() != 0 {
		log.Fatal().Err(ErrPermissionDenied)
	}
}
