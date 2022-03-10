package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/api"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"honnef.co/go/tools/config"
)

func main() {

	var configFile string
	flag.StringVar(&configFile, "c", "", "Configuration file.")
	flag.Parse()

	var config config.Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("Error reading configuration file", configFile, err)
	}

	app := &application.Application{}
	api := api.New(app)

	http.ListenAndServe(":3000", api)
}
