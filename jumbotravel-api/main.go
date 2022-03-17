package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/api"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/config"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/mysql"
)

// @title JumboTravel API
// @version 0.1
// @description Get data from JumboTravel Database.
//
// @contact.name JumboTravel Admin
// @contact.email cpomares@nigul.cide.es
//
// @BasePath /
// @schemes http
//
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {

	var configFile string
	flag.StringVar(&configFile, "c", "", "Configuration file.")
	flag.Parse()

	var config config.Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("Error reading configuration file", configFile, err)
	}

	mysqlDB := &mysql.MySQL{
		Host:         config.Database.Host,
		Port:         config.Database.Port,
		DatabaseName: config.Database.DatabaseName,
		User:         config.Database.User,
		Password:     config.Database.Password,
	}
	err := mysqlDB.Connect()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	defer mysqlDB.Disconnect()

	app := &application.Application{
		MySQLFetcher: mysqlDB,
	}
	api := api.New(app)

	http.ListenAndServe(":3000", api)
}
