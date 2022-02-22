package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/labi-le/control-panel/http/api"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/structures"
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

	Config := structures.Config{}
	_, err := toml.DecodeFile(config, &Config)

	if err != nil {
		log.Fatal(err)
	}

	server := api.Server{
		Config: &Config,
		DB:     internal.NewDB(&Config),
	}

	if server.DB.Migrate() != nil {
		log.Fatal(err)
	}

	err = api.Start(&server)
	if err != nil {
		log.Fatal(err)
	}
}
