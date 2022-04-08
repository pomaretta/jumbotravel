package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/api"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/config"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/mysql"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/utils"
)

var (
	environment = "DEV"
	version     = "0.3"
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

func initializeLogging() (*os.File, *log.Logger) {

	if !utils.IsWorker() {
		return nil, nil
	}

	// If logs directory doesn't exist, create it
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	workerName := strings.ToLower(os.Getenv("PWS_WORKER"))
	if workerName != "" {
		workerName = fmt.Sprintf("-%s", workerName)
	}

	// Create file with timestamp
	logFile, _ := os.Create(fmt.Sprintf(
		"logs/jumbotravel-api%s-%s-%s.json",
		workerName,
		strings.ToLower(environment),
		time.Now().UTC().Format("20060102150405"),
	))
	return logFile, log.New(logFile, "", 0)
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

	// Initialize the logger
	logFile, logger := initializeLogging()
	defer logFile.Close()

	app := &application.Application{
		Config:      &dbConfig,
		Version:     version,
		Environment: environment,
		Logger:      logger,

		MySQLFetcher: mysqlDB,
	}
	api := api.New(app)

	http.ListenAndServe(":3000", api)
}
