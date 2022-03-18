package application

import "github.com/pomaretta/jumbotravel/jumbotravel-api/config"

type Application struct {
	Environment string
	Version     string
	Config      *config.DBConfig

	MySQLFetcher
}
