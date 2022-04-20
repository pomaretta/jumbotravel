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

func (app *Application) GetAgentBookingsAggregate(agentId, flightId int) ([]dto.BookingAggregate, error) {
	return app.MySQLFetcher.FetchAgentBookingsAggregate(agentId, flightId)
}

func (app *Application) GetAgentBookingDetails(agentId int, bookingReferenceId string) (*dto.BookingAggregate, error) {
	return app.MySQLFetcher.FetchAgentBookingDetails(agentId, bookingReferenceId)
}

func (app *Application) GetAgentBookingOperations(agentId int, bookingReferenceId string) ([]entity.Notification, error) {
	return app.MySQLFetcher.FetchAgentBookingOperations(agentId, bookingReferenceId)
}

func (app *Application) GetAgentBookingItems(agentId int, bookingReferenceId string) ([]dto.BookingItem, error) {
	return app.MySQLFetcher.FetchAgentBookingItems(agentId, bookingReferenceId)
}

func (app *Application) UpdateFlightStatus(flightId int, status string) (int64, error) {
	return app.MySQLFetcher.UpdateFlightStatus(flightId, status)
}

func (app *Application) PutBookingCreation(bookings []dto.BookingInput) (int64, error) {
	return app.MySQLFetcher.PutBooking(bookings)
}
