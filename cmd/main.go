package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/http/api"
	"github.com/labi-le/control-panel/internal/structures"
	"log"
)

var (
	config string
)

func init() {
	flag.StringVar(&config, "config", "config.toml", "path to config file")
}

func main() {
	flag.Parse()

	var conf *structures.Config
	_, err := toml.DecodeFile(config, &conf)

	if err != nil {
		log.Fatal(err)
	}

	db := internal.NewDB(conf)
	if db.Migrate() != nil {
		log.Fatal(err)
	}

	apiResolver := api.NewMethods(db)
	srv := api.NewServer(apiResolver.GetRoutes(), conf)

	if srv.Start() != nil {
		log.Fatal(err)
	}
}
