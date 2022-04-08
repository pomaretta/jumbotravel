package application

import (
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
)

func (app *Application) GetAgentNotifications(agentId int, seen, active, expired, popup string) (s []entity.Notification, err error) {
	return app.MySQLFetcher.FetchAgentNotifications(agentId, seen, active, expired, popup)
}

func (app *Application) PostUpdateAgentNotifications(notificationIds []int) (int64, error) {
	return app.MySQLFetcher.UpdateAgentNotifications(notificationIds)
}

func (app *Application) GetAgentFlights(agentId int, routeId, flightId, airplaneId int, status string, departureTime, arrivalTime time.Time) (s []dto.AgentFlight, err error) {
	return app.MySQLFetcher.FetchAgentFlights(agentId, routeId, flightId, airplaneId, status, departureTime, arrivalTime)
}

func (app *Application) GetAgentFlightOperations(agentId, flightId int) ([]entity.Notification, error) {
	return app.MySQLFetcher.FetchAgentFlightOperations(agentId, flightId)
}

func (app *Application) GetAgentFlightAgents(agentId, flightId int) ([]dto.FlightAgent, error) {
	return app.MySQLFetcher.FetchAgentFlightAgents(agentId, flightId)
}

func (app *Application) GetAgentFlightProducts(agentId, flightId int) ([]dto.FlightProduct, error) {
	return app.MySQLFetcher.FetchAgentFlightProducts(agentId, flightId)
}
