package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	web "github.com/labi-le/control-panel/http"
	"github.com/labi-le/control-panel/internal"
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
	Config := internal.Config{}

	_, err := toml.DecodeFile(config, &Config)

	if err != nil {
		log.Fatal(err)
	}

	db := internal.NewDB(Config)
	Config.DB = db

	if db.Migrate() != nil {
		log.Fatal(err)
	}

	err = web.Start(&Config)
	if err != nil {
		log.Fatal(err)
	}
}
