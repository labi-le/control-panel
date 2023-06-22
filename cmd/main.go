package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/pkg/log"
	"go.uber.org/zap"
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
	flag.StringVar(&config, "config", internal.DefaultConfigPath, "path to config file")
	flag.BoolVar(&versionFlag, "version", false, "print version and exit")
}

func main() {
	flag.Parse()

	if versionFlag == true {
		log.Info(internal.PanelVersion)
		return
	}

	conf, err := internal.NewPanelSettings(config)
	if err != nil {
		log.Fatal(err)
	}

	checkPermissions()

	logger := MustLogger(log.New())

	// Setting up signal capturing
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srv := fiber.New(fiber.Config{})
	internal.RegisterHandlers(srv, conf)

	go func() {
		err := srv.Listen(conf.GetAddr() + ":" + conf.GetPort())
		if err != nil {
			logger.Fatal(err)
		}
	}()

	<-ctx.Done()
	// Waiting for SIGINT (kill -2)
	logger.Info("Gracefully shutdown server...")

	if err := srv.Shutdown(); err != nil {
		logger.Fatal(err)
	}

}

func checkPermissions() {
	if os.Geteuid() != 0 {
		log.Fatal("You must run this program as root")
	}
}

func MustLogger(l log.Logger) log.Logger {
	if debugMode {
		devlogger, err := zap.NewDevelopment(zap.IncreaseLevel(zap.DebugLevel))
		if err != nil {
			panic(err.Error())
		}

		l = log.NewWithZap(devlogger)
		l.Info("debug mode enabled")
	}

	log.SetGlobalLog(l)

	return l
}
