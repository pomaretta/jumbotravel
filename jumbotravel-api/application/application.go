package application

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/config"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
)

type Application struct {
	Environment string
	Version     string
	Config      *config.DBConfig

	MySQLFetcher
}

type MySQLFetcher interface {

	// Master
	FetchMasterAirports(airpotId int, country, city, airport string) ([]entity.MasterAirport, error)
	FetchMasterAgents(agentId int, dni, agentType, email string, active bool) ([]entity.Agent, error)
	FetchMasterAirplanes(airplaneId, flightNumber int, carrier string) ([]entity.Airplane, error)
	FetchMasterProducts(productId, productCode int) ([]entity.Product, error)

	// Stock
	FetchStock(stockId, airplaneId, productId, productCode int) ([]dto.Stock, error)
}
