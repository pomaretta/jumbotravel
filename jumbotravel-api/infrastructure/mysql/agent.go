package mysql

import (
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/agentbuilders"
)

func (db *MySQL) FetchAgentNotifications(agentId int, seen, active, expired, popup string) (s []entity.Notification, err error) {

	qb := &agentbuilders.NotificationsQueryBuilder{}
	qb.SetAgentId(agentId)
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
