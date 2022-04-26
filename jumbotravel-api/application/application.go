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
	FetchMasterProducts(productId, productCode []int) ([]entity.Product, error)
	FetchNotifications(notificationId, resourceId []int, resourceUuid, notificationType, scope []string, seen, active, expired, popup string) ([]entity.Notification, error)
	FetchMasterFlights(flightId, routeId, airplaneId, agentId, providerId int, status string, departureTime, arrivalTime time.Time) ([]entity.MasterFlight, error)

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
	FetchAgentNotifications(agentId int, agentType, seen, active, expired, popup string) ([]entity.Notification, error)
	UpdateAgentNotifications(notificationIds []int) (int64, error)

	// Agent Flights
	FetchAgentFlights(agentId int, routeId, flightId, airplaneId int, status string, departureTime, arrivalTime time.Time) (s []dto.AgentFlight, err error)
	FetchAgentFlightOperations(agentId, flightId int) ([]entity.Notification, error)
	FetchAgentFlightAgents(agentId, flightId int) ([]dto.FlightAgent, error)
	FetchAgentFlightProducts(agentId, flightId int) ([]dto.FlightProduct, error)
	FetchAgentBookingsAggregate(agentId int, agentType string, flightId, airplaneId int) ([]dto.BookingAggregate, error)
	FetchAgentBookingsAggregateWithDays(agentId int, agentType string, flightId, airplaneId int, from, to time.Time) ([]dto.BookingAggregate, error)

	// Agent Bookings
	FetchAgentBookingDetails(agentId int, agentType, bookingReferenceId string) (*dto.BookingAggregate, error)
	FetchAgentBookingOperations(agentId int, bookingReferenceId string) ([]entity.Notification, error)
	FetchAgentBookingItems(agentId int, agentType, bookingReferenceId string) ([]dto.BookingItem, error)
	FetchAgentBookingCount(agentId int, agentType, countType string, flightId, airplaneId, days int) ([]dto.BookingCount, error)
	FetchAgentFlightsCount(agentId int, agentType string, flightId, airplaneId, days int) ([]dto.BookingCount, error)
	FetchAgentBookingCompositeCount(agentId int, agentType string, flightId, airplaneId, days int) ([]dto.BookingCompositeCount, error)

	// Functionalities
	PutNotification(notifications []dto.NotificationInput) (int64, error)
	UpdateFlightStatus(flightId int, status string) (int64, error)
	PutBooking(bookings []dto.BookingInput) (int64, error)
	UpdateBooking(bookingReferenceId, status string, providerId, providerMappingId int) (int64, error)
	UpdateStock(stock []dto.StockInput) (int64, error)

	// Invoice
	FetchInvoices(invoiceId, agentId, providerId int, bookingReferenceId string, from, to time.Time) ([]dto.Invoice, error)
	PutInvoice(invoice dto.InvoiceInput) (int64, error)
	PutInvoiceBookings(invoiceBookings []dto.InvoiceBookingInput) (int64, error)
}
