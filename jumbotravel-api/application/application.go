package application

import (
	"log"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/config"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
)

type Application struct {
	Environment string
	Version     string
	Config      *config.DBConfig
	Logger      *log.Logger

	MySQLFetcher
}

type MySQLFetcher interface {

	// Master
	FetchMasterAirports(airpotId int, country, city, airport string) ([]entity.MasterAirport, error)
	FetchMasterAgents(agentId int, dni, agentType, email string, active bool) ([]entity.Agent, error)
	FetchMasterAirplanes(airplaneId, flightNumber int, carrier string) ([]entity.Airplane, error)
	FetchMasterProducts(productId, productCode int) ([]entity.Product, error)
	FetchNotifications(notificationId, resourceId []int, notificationType, scope []string, seen, active, expired, popup string) ([]entity.Notification, error)

	// Stock
	FetchStock(stockId, airplaneId, productId, productCode int) ([]dto.Stock, error)

	// Flights
	FetchRoute(routeId, airplaneId, flightId int, departureCountry, arrivalCountry, departureCity, arrivalCity, departureAirport, arrivalAirport, status, carrier string) ([]dto.Route, error)

	// Auth
	FetchAgentAuth(dni string) (dto.AgentAuth, error)
	FetchAuthToken(agentId int, jti string, active string, expired string, single bool) ([]entity.Token, error)
	PutToken(tokenId, subject, token string, agentId int, issuedAt, expiresAt time.Time) error

	// Access Logging
	PutAccessLogging(requestId, tokenId, tokenName, ip, method, path, query, errorMessage string, status int) error

	// Agent
	FetchAgentNotifications(agentId int, seen, active, expired, popup string) ([]entity.Notification, error)
	UpdateAgentNotifications(notificationIds []int) (int64, error)

	// Agent Flights
	FetchAgentFlights(agentId int, routeId, flightId, airplaneId int, status string, departureTime, arrivalTime time.Time) (s []dto.AgentFlight, err error)
}
