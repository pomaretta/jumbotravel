package application

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
)

func (app *Application) GetMasterAirports(airpotId int, country, city, airport string) ([]entity.MasterAirport, error) {
	return app.MySQLFetcher.FetchMasterAirports(airpotId, country, city, airport)
}

func (app *Application) GetMasterAgents(agentId int, dni, agentType, email string, active bool) ([]entity.Agent, error) {
	return app.MySQLFetcher.FetchMasterAgents(agentId, dni, agentType, email, active)
}

func (app *Application) GetMasterAirplanes(airplaneId, flightNumber int, carrier string) ([]entity.Airplane, error) {
	return app.MySQLFetcher.FetchMasterAirplanes(airplaneId, flightNumber, carrier)
}

func (app *Application) GetMasterProducts(productId, productCode []int) ([]entity.Product, error) {
	return app.MySQLFetcher.FetchMasterProducts(productId, productCode)
}

func (app *Application) GetNotifications(notificationId, resourceId []int, resourceUuid, notificationType, scope []string, seen, active, expired, popup string) ([]entity.Notification, error) {
	return app.MySQLFetcher.FetchNotifications(notificationId, resourceId, resourceUuid, notificationType, scope, seen, active, expired, popup)
}

func (app *Application) PutNotifications(notifications []dto.NotificationInput) (int64, error) {
	return app.MySQLFetcher.PutNotification(notifications)
}

func (app *Application) GetInvoices(invoiceId, agentId, providerId int, bookingReferenceId string) ([]dto.Invoice, error) {
	return app.MySQLFetcher.FetchInvoices(invoiceId, agentId, providerId, bookingReferenceId)
}

func (app *Application) RegisterInvoice(invoice dto.InvoiceInput) (int64, error) {
	return app.MySQLFetcher.PutInvoice(invoice)
}

func (app *Application) RegisterInvoiceBookings(invoiceBookings []dto.InvoiceBookingInput) (int64, error) {
	return app.MySQLFetcher.PutInvoiceBookings(invoiceBookings)
}
