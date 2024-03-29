package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/pkg/pm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

var (
	config string
)

func init() {
	flag.StringVar(&config, "config", internal.DefaultConfigPath(), "path to config file")
}

func main() {
	flag.Parse()

	conf, err := internal.NewPanelSettings(config)
	if err != nil {
		log.Fatal().Err(err)
	}

	configureLogger(conf)

	log.Info().Msg(internal.BuildVersion())

	checkPermissions()

	// Setting up signal capturing
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srv := fiber.New(fiber.Config{})
	internal.RegisterHandlers(srv, conf, pm.MustManager())

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

func configureLogger(conf *internal.PanelSettings) {
	level, err := zerolog.ParseLevel(conf.LogLevel)
	if err != nil {
		log.Warn().Msgf("invalid log level %s, fallback to info", conf.LogLevel)
		level = zerolog.InfoLevel
	}

	log.Info().Msgf("log level set to %s", level)
	zerolog.SetGlobalLevel(level)
}

var ErrPermissionDenied = errors.New("permission denied")

func checkPermissions() {
	if os.Geteuid() != 0 {
		log.Fatal().Err(ErrPermissionDenied)
	}
}
