package application

import (
	"time"

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

	// Flights
	FetchRoute(routeId, airplaneId, flightId int, departureCountry, arrivalCountry, departureCity, arrivalCity, departureAirport, arrivalAirport, status, carrier string) ([]dto.Route, error)

	// Auth
	FetchAgentAuth(dni string) (dto.AgentAuth, error)
	FetchAuthToken(agentId int) (string, error)
	PutToken(tokenId, subject, token string, agentId int, issuedAt, expiresAt time.Time) error

	// Access Logging
	PutAccessLogging(requestId, tokenId, tokenName, ip, method, path, query, errorMessage string, status int) error
}
