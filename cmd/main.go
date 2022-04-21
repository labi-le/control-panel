package main

import (
	"flag"
	"fmt"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/http/api"
	"log"
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

	apiResolver := api.NewMethods(conf)
	srv := api.NewServer(apiResolver.GetRoutes(), conf)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
