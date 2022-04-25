package application

import (
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
)

func (app *Application) GetAgentNotifications(agentId int, agentType, seen, active, expired, popup string) (s []entity.Notification, err error) {
	return app.MySQLFetcher.FetchAgentNotifications(agentId, agentType, seen, active, expired, popup)
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

func (app *Application) GetAgentBookingsAggregate(agentId int, agentType string, flightId, airplaneId int) ([]dto.BookingAggregate, error) {
	return app.MySQLFetcher.FetchAgentBookingsAggregate(agentId, agentType, flightId, airplaneId)
}

func (app *Application) GetAgentBookingDetails(agentId int, agenType, bookingReferenceId string) (*dto.BookingAggregate, error) {
	return app.MySQLFetcher.FetchAgentBookingDetails(agentId, agenType, bookingReferenceId)
}

func (app *Application) GetAgentBookingOperations(agentId int, bookingReferenceId string) ([]entity.Notification, error) {
	return app.MySQLFetcher.FetchAgentBookingOperations(agentId, bookingReferenceId)
}

func (app *Application) GetAgentBookingItems(agentId int, agentType, bookingReferenceId string) ([]dto.BookingItem, error) {
	return app.MySQLFetcher.FetchAgentBookingItems(agentId, agentType, bookingReferenceId)
}

func (app *Application) UpdateFlightStatus(flightId int, status string) (int64, error) {
	return app.MySQLFetcher.UpdateFlightStatus(flightId, status)
}

func (app *Application) PutBookingCreation(bookings []dto.BookingInput) (int64, error) {
	return app.MySQLFetcher.PutBooking(bookings)
}

func (app *Application) UpdateBookingStatus(bookingReferenceId string, status string, providerId, providerMappingId int) (int64, error) {
	return app.MySQLFetcher.UpdateBooking(bookingReferenceId, status, providerId, providerMappingId)
}

func (app *Application) UpdateStockStatus(stock []dto.StockInput) (int64, error) {
	return app.MySQLFetcher.UpdateStock(stock)
}

func (app *Application) GetAgentBookingCount(agentId int, agentType, countType string, flightId, airplaneId, days int) ([]dto.BookingCount, error) {
	return app.MySQLFetcher.FetchAgentBookingCount(agentId, agentType, countType, flightId, airplaneId, days)
}

func (app *Application) GetAgentBookingCompositeCount(agentId int, agentType string, flightId, airplaneId, days int) ([]dto.BookingCompositeCount, error) {
	return app.MySQLFetcher.FetchAgentBookingCompositeCount(agentId, agentType, flightId, airplaneId, days)
}

func (app *Application) GetAgentFlightCount(agentId int, agentType string, flightId, airplaneId, days int) ([]dto.BookingCount, error) {
	return app.MySQLFetcher.FetchAgentFlightsCount(agentId, agentType, flightId, airplaneId, days)
}
