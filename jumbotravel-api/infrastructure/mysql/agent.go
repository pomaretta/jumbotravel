package mysql

import (
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/agentbuilders"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/masterbuilders"
)

func (db *MySQL) FetchAgentNotifications(agentId int, agentType, seen, active, expired, popup string) (s []entity.Notification, err error) {

	qb := &agentbuilders.NotificationsQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetAgenType(agentType)
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

func (db *MySQL) UpdateAgentNotifications(notificationIds []int) (int64, error) {

	qb := &agentbuilders.ReadNotificationsQueryBuilder{}
	qb.SetNotificationIds(notificationIds)

	return db.Update(qb)
}

func (db *MySQL) FetchAgentFlights(agentId int, routeId, flightId, airplaneId int, status string, departureTime, arrivalTime time.Time) (s []dto.AgentFlight, err error) {

	qb := &agentbuilders.FlightsQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetRouteId(routeId)
	qb.SetFlightId(flightId)
	qb.SetAirplaneId(airplaneId)
	qb.SetStatus(status)
	qb.SetDepartureTime(departureTime)
	qb.SetArrivalTime(arrivalTime)

	ent, err := db.Fetch(&dto.AgentFlight{}, qb)
	if err != nil {
		return
	}

	for _, e := range ent {
		s = append(s, e.(dto.AgentFlight))
	}

	return
}

func (db *MySQL) FetchAgentFlightOperations(agentId, flightId int) (s []entity.Notification, err error) {

	qb := &agentbuilders.FlightOperationsQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetFlightId(flightId)

	ent, err := db.Fetch(&entity.Notification{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.Notification))
	}

	return
}

func (db *MySQL) FetchAgentFlightAgents(agentId, flightId int) (s []dto.FlightAgent, err error) {

	qb := &agentbuilders.FlightAgentsQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetFlightId(flightId)

	ent, err := db.Fetch(&dto.FlightAgent{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.FlightAgent))
	}

	return
}

func (db *MySQL) FetchAgentFlightProducts(agentId, flightId int) (s []dto.FlightProduct, err error) {

	qb := &agentbuilders.FlightProductsQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetFlightId(flightId)

	ent, err := db.Fetch(&dto.FlightProduct{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.FlightProduct))
	}

	return
}

func (db *MySQL) FetchAgentBookingsAggregate(agentId int, agentType string, flightId, airplaneId int) (s []dto.BookingAggregate, err error) {

	qb := &agentbuilders.BookingsAggrQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetFlightId(flightId)
	qb.SetAgentType(agentType)
	qb.SetAirplaneId(airplaneId)

	ent, err := db.Fetch(&dto.BookingAggregate{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.BookingAggregate))
	}

	return
}

func (db *MySQL) FetchAgentBookingsAggregateWithDays(agentId int, agentType string, flightId, airplaneId int, from, to time.Time) (s []dto.BookingAggregate, err error) {

	qb := &agentbuilders.BookingsAggrQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetFlightId(flightId)
	qb.SetAgentType(agentType)
	qb.SetAirplaneId(airplaneId)
	qb.SetFrom(from)
	qb.SetTo(to)

	ent, err := db.Fetch(&dto.BookingAggregate{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.BookingAggregate))
	}

	return
}

func (db *MySQL) FetchAgentBookingDetails(agentId int, agentType, bookingReferenceId string) (*dto.BookingAggregate, error) {

	qb := &agentbuilders.BookingsAggrQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetAgentType(agentType)
	qb.SetBookingReferenceId(bookingReferenceId)

	ent, err := db.Fetch(&dto.BookingAggregate{}, qb)
	if err != nil {
		return nil, err
	}
	if len(ent) == 0 {
		return nil, nil
	}
	bookingDetails := ent[0].(dto.BookingAggregate)

	return &bookingDetails, nil
}

func (db *MySQL) FetchAgentBookingOperations(agentId int, bookingReferenceId string) (s []entity.Notification, err error) {

	qb := &masterbuilders.NotificationQueryBuilder{}
	qb.SetScope([]string{"BOOKING"})
	qb.SetResourceUuid([]string{bookingReferenceId})
	qb.SetActive("0")
	qb.SetExpired("0")
	qb.SetPopup("0")
	qb.SetSeen("0")

	ent, err := db.Fetch(&entity.Notification{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.Notification))
	}

	return
}

func (db *MySQL) FetchAgentBookingItems(agentId int, agentType, bookingReferenceId string) (s []dto.BookingItem, err error) {

	qb := &agentbuilders.BookingItemsQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetAgentType(agentType)
	qb.SetBookingReferenceId(bookingReferenceId)

	ent, err := db.Fetch(&dto.BookingItem{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.BookingItem))
	}

	return
}

func (db *MySQL) UpdateFlightStatus(flightId int, status string) (int64, error) {

	qb := &agentbuilders.UpdateFlightStatusQueryBuilder{}
	qb.SetFlightId(flightId)
	qb.SetStatus(status)

	return db.Update(qb)
}

func (db *MySQL) PutBooking(bookings []dto.BookingInput) (int64, error) {

	qb := &agentbuilders.PutBookingQueryBuilder{}
	qb.SetBookings(bookings)

	return db.Update(qb)
}

func (db *MySQL) FetchInvoices(invoiceId, agentId, providerId int, bookingReferenceId string, from, to time.Time) (s []dto.Invoice, err error) {

	qb := &masterbuilders.InvoiceQueryBuilder{}
	qb.SetInvoiceId(invoiceId)
	qb.SetAgentId(agentId)
	qb.SetProviderId(providerId)
	qb.SetBookingReferenceId(bookingReferenceId)
	qb.SetFrom(from)
	qb.SetTo(to)

	ent, err := db.Fetch(&dto.Invoice{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.Invoice))
	}

	return
}

func (db *MySQL) UpdateBooking(bookingReferenceId, status string, providerId, providerMappingId int) (int64, error) {

	qb := &agentbuilders.UpdateBookingQueryBuilder{}
	qb.SetBookingReferenceId(bookingReferenceId)
	qb.SetStatus(status)
	qb.SetProviderId(providerId)
	qb.SetProviderMappingId(providerMappingId)

	return db.Update(qb)
}

func (db *MySQL) UpdateStock(stock []dto.StockInput) (int64, error) {

	qb := &agentbuilders.UpdateProductStockQueryBuilder{}
	qb.SetProducts(stock)

	return db.Put(qb)
}

func (db *MySQL) FetchAgentBookingCount(agentId int, agentType, countType string, flightId, airplaneId, days int) (s []dto.BookingCount, err error) {

	qb := &agentbuilders.BookingCountQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetAgentType(agentType)
	qb.SetFlightId(flightId)
	qb.SetAirplaneId(airplaneId)
	qb.SetDays(days)
	qb.SetCountType(countType)

	ent, err := db.Fetch(&dto.BookingCount{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.BookingCount))
	}

	return
}

func (db *MySQL) FetchAgentBookingCompositeCount(agentId int, agentType string, flightId, airplaneId, days int) (s []dto.BookingCompositeCount, err error) {

	qb := &agentbuilders.BookingCompositeQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetAgentType(agentType)
	qb.SetFlightId(flightId)
	qb.SetAirplaneId(airplaneId)
	qb.SetDays(days)

	ent, err := db.Fetch(&dto.BookingCompositeCount{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.BookingCompositeCount))
	}

	return
}

func (db *MySQL) FetchAgentFlightsCount(agentId int, agentType string, flightId, airplaneId, days int) (s []dto.BookingCount, err error) {

	qb := &agentbuilders.FlightsCountQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetAgentType(agentType)
	qb.SetFlightId(flightId)
	qb.SetAirplaneId(airplaneId)
	qb.SetDays(days)

	ent, err := db.Fetch(&dto.BookingCount{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.BookingCount))
	}

	return
}
