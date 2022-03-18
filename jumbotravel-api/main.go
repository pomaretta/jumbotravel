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

var (
	environment = "DEV"
	version     = "0.2"
)

func establishDatabaseConnection(config *config.DBConfig) *mysql.MySQL {
	mysqlDB := &mysql.MySQL{
		Host:         config.Host,
		Port:         config.Port,
		DatabaseName: config.DatabaseName,
		User:         config.User,
		Password:     config.Password,
	}
	return mysqlDB
}

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

	var configuration config.Config
	if _, err := toml.DecodeFile(configFile, &configuration); err != nil {
		log.Fatal("Error reading configuration file", configFile, err)
	}

	dbConfig := configuration.Database[0]
	if environment == "DEV" {
		dbConfig = configuration.Database[1]
	}

	mysqlDB := establishDatabaseConnection(&dbConfig)
	err := mysqlDB.Connect()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	defer mysqlDB.Disconnect()

	app := &application.Application{
		Config:      &dbConfig,
		Version:     version,
		Environment: environment,

		MySQLFetcher: mysqlDB,
	}
	api := api.New(app)

	http.ListenAndServe(":3000", api)
}
