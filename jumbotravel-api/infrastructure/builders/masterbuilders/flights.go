package masterbuilders

import (
	"fmt"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type MasterFlightsQueryBuilder struct {
	builders.MySQLQueryBuilder

	flightId      int
	routeId       int
	airplaneId    int
	agentId       int
	providerId    int
	status        string
	departureTime time.Time
	arrivalTime   time.Time
}

func (qb *MasterFlightsQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *MasterFlightsQueryBuilder) SetRouteId(routeId int) {
	qb.routeId = routeId
}

func (qb *MasterFlightsQueryBuilder) SetAirplaneId(airplaneId int) {
	qb.airplaneId = airplaneId
}

func (qb *MasterFlightsQueryBuilder) SetAgentId(agentId int) {

	qb.agentId = agentId
}

func (qb *MasterFlightsQueryBuilder) SetProviderId(providerId int) {
	qb.providerId = providerId
}

func (qb *MasterFlightsQueryBuilder) SetStatus(status string) {
	qb.status = status
}

func (qb *MasterFlightsQueryBuilder) SetDepartureTime(departureTime time.Time) {
	qb.departureTime = departureTime
}

func (qb *MasterFlightsQueryBuilder) SetArrivalTime(arrivalTime time.Time) {
	qb.arrivalTime = arrivalTime
}

func (qb *MasterFlightsQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "WHERE 1=1"
	args := []interface{}{}

	if qb.flightId != 0 {
		partialQuery = fmt.Sprintf("%s AND f.flight_id != ?", partialQuery)
		args = append(args, qb.flightId)
	}

	if qb.routeId != 0 {
		partialQuery = fmt.Sprintf("%s AND f.route_id = ?", partialQuery)
	}

	if qb.airplaneId != 0 {
		partialQuery = fmt.Sprintf("%s AND fr.airplanemapping_id = ?", partialQuery)
		args = append(args, qb.airplaneId)
	}

	if qb.agentId != 0 {
		partialQuery = fmt.Sprintf("%s AND fa.agentmapping_id = ?", partialQuery)
		args = append(args, qb.agentId)
	}

	if qb.providerId != 0 {
		partialQuery = fmt.Sprintf("%s AND ma3.agentmapping_id = ?", partialQuery)
		args = append(args, qb.providerId)
	}

	if qb.status != "" {
		partialQuery = fmt.Sprintf("%s AND f.status != ?", partialQuery)
		args = append(args, qb.status)
	}

	if !qb.departureTime.IsZero() {
		partialQuery = fmt.Sprintf("%s AND f.departure_time >= ?", partialQuery)
		args = append(args, qb.departureTime)
	}

	if !qb.arrivalTime.IsZero() {
		partialQuery = fmt.Sprintf("%s AND f.arrival_time <= ?", partialQuery)
		args = append(args, qb.arrivalTime)
	}

	return partialQuery, args, nil
}

func (qb *MasterFlightsQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}
	orderClause := "ORDER BY f.created_at"

	query := fmt.Sprintf(`
		SELECT
			f.flight_id,
			f.route_id,
			f.status,
			f.departure_time,
			f.arrival_time,
			f.created_at,
			f.updated_at
		FROM flights f 
		LEFT JOIN flight_routes fr ON fr.route_id = f.route_id
		LEFT JOIN flight_agents fa ON fa.flight_id = f.flight_id
		LEFT JOIN master_airports ma ON ma.country = fr.arrival_country AND ma.city = fr.arrival_city AND ma.airport = fr.arrival_airport
		LEFT JOIN master_agents ma2 ON ma2.airport_id = ma.airport_id
		LEFT JOIN master_agentmapping ma3 ON ma3.agent_id = ma2.agent_id
		%s
		%s
	`, whereClauses, orderClause)

	return query, args, nil
}
