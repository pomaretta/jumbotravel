package mysql

import (
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/masterbuilders"
)

func (db *MySQL) FetchMasterAirports(airpotId int, country, city, airport string) (s []entity.MasterAirport, err error) {

	qb := &masterbuilders.MasterAirportQueryBuilder{}
	qb.SetAirportID(airpotId)
	qb.SetCountry(country)
	qb.SetCity(city)
	qb.SetAirport(airport)

	ent, err := db.Fetch(&entity.MasterAirport{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.MasterAirport))
	}

	return
}

func (db *MySQL) FetchMasterAgents(agentId int, dni, agentType, email string, active bool) (s []entity.Agent, err error) {

	qb := &masterbuilders.MasterAgentQueryBuilder{}
	qb.SetAgentID(agentId)
	qb.SetDNI(dni)
	qb.SetType(agentType)
	qb.SetActive(active)
	qb.SetEmail(email)

	ent, err := db.Fetch(&entity.Agent{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.Agent))
	}

	return
}

func (db *MySQL) FetchMasterAirplanes(airplaneId, flightNumber int, carrier string) (s []entity.Airplane, err error) {

	qb := &masterbuilders.AirplaneQueryBuilder{}
	qb.SetAirplaneID(airplaneId)
	qb.SetCarrier(carrier)
	qb.SetFlightNumber(flightNumber)

	ent, err := db.Fetch(&entity.Airplane{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.Airplane))
	}

	return
}

func (db *MySQL) FetchMasterProducts(productId, productCode []int) (s []entity.Product, err error) {

	qb := &masterbuilders.MasterProductQueryBuilder{}
	qb.SetProductID(productId)
	qb.SetProductCode(productCode)

	ent, err := db.Fetch(&entity.Product{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.Product))
	}

	return
}

func (db *MySQL) FetchNotifications(notificationId, resourceId []int, resourceUuid, notificationType, scope []string, seen, active, expired, popup string) (s []entity.Notification, err error) {

	qb := &masterbuilders.NotificationQueryBuilder{}
	qb.SetNotificationId(notificationId)
	qb.SetResourceId(resourceId)
	qb.SetResourceUuid(resourceUuid)
	qb.SetNotificationType(notificationType)
	qb.SetScope(scope)
	qb.SetSeen(seen)
	qb.SetActive(active)
	qb.SetExpired(expired)
	qb.SetPopup(popup)

	ent, err := db.Fetch(&entity.Notification{}, qb)
	if err != nil {
		return
	}

	for _, e := range ent {
		s = append(s, e.(entity.Notification))
	}

	return
}

func (db *MySQL) PutNotification(notifications []dto.NotificationInput) (int64, error) {

	qb := &masterbuilders.PutNotificationQueryBuilder{}
	qb.SetNotifications(notifications)

	return db.Put(qb)
}

func (db *MySQL) PutInvoice(invoice dto.InvoiceInput) (int64, error) {

	qb := &masterbuilders.PutInvoiceQueryBuilder{}
	qb.SetInvoice(invoice)

	return db.Put(qb)
}

func (db *MySQL) PutInvoiceBookings(invoiceBookings []dto.InvoiceBookingInput) (int64, error) {

	qb := &masterbuilders.PutInvoiceBookingsQueryBuilder{}
	qb.SetBookings(invoiceBookings)

	return db.Put(qb)
}

func (db *MySQL) FetchMasterFlights(flightId, routeId, airplaneId, agentId, providerId int, status string, departureTime, arrivalTime time.Time) (s []entity.MasterFlight, err error) {
	qb := &masterbuilders.MasterFlightsQueryBuilder{}
	qb.SetFlightId(flightId)
	qb.SetRouteId(routeId)
	qb.SetAirplaneId(airplaneId)
	qb.SetAgentId(agentId)
	qb.SetProviderId(providerId)
	qb.SetStatus(status)
	qb.SetDepartureTime(departureTime)
	qb.SetArrivalTime(arrivalTime)

	ent, err := db.Fetch(&entity.MasterFlight{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.MasterFlight))
	}

	return
}
