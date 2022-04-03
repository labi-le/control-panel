package main

import (
	"flag"
	"fmt"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/http/api"
	"github.com/labi-le/control-panel/internal/structures"
	"log"
)

var (
	config      string
	versionFlag bool
)

func init() {
	flag.StringVar(&config, "config", structures.DefaultConfig, "path to config file")
	flag.BoolVar(&versionFlag, "version", false, "print version and exit")
}

func main() {
	flag.Parse()

	if versionFlag == true {
		fmt.Println(structures.PanelVersion)
		return
	}

	conf := structures.NewConfig()
	_, err := conf.LoadConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	//
	db := internal.NewDB(conf)
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	apiResolver := api.NewMethods(db)
	srv := api.NewServer(apiResolver.GetRoutes(), conf)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
