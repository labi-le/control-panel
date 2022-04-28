package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/http/api"
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

// @title          Control Panel API
// @version        1.0
// @description    Control Panel API
// @contact.name   Control Panel API Support
// @contact.url    https://github.com/labi-le/control-panel
// @contact.email  https://github.com/labi-le/control-panel/issues
// @license.name   GNU General Public License v3.0
// @license.url    https://www.gnu.org/licenses/gpl-3.0.html
// @host           localhost:8080
// @BasePath       /api
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

	apiResolver := api.NewMethods(conf)
	srv := api.NewServer(apiResolver.GetRoutes(), conf)

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
