package agentbuilders

import (
	"fmt"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type FlightsQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId       int
	routeId       int
	flightId      int
	airplaneId    int
	status        string
	departureTime time.Time
	arrivalTime   time.Time
}

func (qb *FlightsQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *FlightsQueryBuilder) SetRouteId(routeId int) {
	qb.routeId = routeId
}

func (qb *FlightsQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *FlightsQueryBuilder) SetAirplaneId(airplaneId int) {
	qb.airplaneId = airplaneId
}

func (qb *FlightsQueryBuilder) SetStatus(status string) {
	qb.status = status
}

func (qb *FlightsQueryBuilder) SetDepartureTime(departureTime time.Time) {
	qb.departureTime = departureTime
}

func (qb *FlightsQueryBuilder) SetArrivalTime(arrivalTime time.Time) {
	qb.arrivalTime = arrivalTime
}

func (qb *FlightsQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Agent ID is required
	if qb.agentId <= 0 {
		return "", nil, fmt.Errorf("agent id is required")
	}
	partialQuery = fmt.Sprintf("%s and ma2.agentmapping_id = ?", partialQuery)
	args = append(args, qb.agentId)

	if qb.routeId > 0 {
		partialQuery = fmt.Sprintf("%s and fr.route_id = ?", partialQuery)
		args = append(args, qb.routeId)
	}

	if qb.flightId > 0 {
		partialQuery = fmt.Sprintf("%s and f.flight_id = ?", partialQuery)
		args = append(args, qb.flightId)
	}

	if qb.airplaneId > 0 {
		partialQuery = fmt.Sprintf("%s and ma.airplane_id = ?", partialQuery)
		args = append(args, qb.airplaneId)
	}

	if qb.status != "" {
		partialQuery = fmt.Sprintf("%s and f.status = ?", partialQuery)
		args = append(args, qb.status)
	}

	if !qb.departureTime.IsZero() {
		partialQuery = fmt.Sprintf("%s and f.departure_time >= ?", partialQuery)
		args = append(args, qb.departureTime)
	}

	if !qb.arrivalTime.IsZero() {
		partialQuery = fmt.Sprintf("%s and f.arrival_time <= ?", partialQuery)
		args = append(args, qb.arrivalTime)
	}

	return partialQuery, args, nil
}

func (qb *FlightsQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClause, args, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}

	orderClause := "order by f.departure_time asc"

	query := fmt.Sprintf(`
		select
			fr.route_id,
			ma.airplane_id,
			ma.carrier,
			ma.flight_number,
			ma.seats,
			f.flight_id,
			f.status,
			f.departure_time,
			f.arrival_time,
			f.created_at as flight_creation,
			f.updated_at as flight_lastmodified,
			fr.departure_country,
			fr.arrival_country,
			fr.departure_city,
			fr.arrival_city,
			fr.departure_airport,
			fr.arrival_airport,
			depair.common_name as departure_commonname,
			arrair.common_name as arrival_commoname,
			fr.updated_at,
			fr.created_at
		from
			flight_routes fr
		right join
			flights f 
			on f.route_id = fr.route_id 
		left join master_airplanesmapping mam
			on mam.airplanemapping_id = fr.airplanemapping_id
		left join master_airplanes ma 
			on ma.airplane_id = mam.airplanemapping_id
		inner join master_airports depair
			on depair.country = fr.departure_country and depair.city = fr.departure_city and depair.airport = fr.departure_airport
		inner join master_airports arrair
			on arrair.country = fr.arrival_country and arrair.city = fr.arrival_city and arrair.airport = fr.arrival_airport
		left join flight_agents fa 
			on fa.flight_id = f.flight_id
		left join master_agentmapping ma2 
			on fa.agentmapping_id = ma2.agentmapping_id
		%s %s
	`, whereClause, orderClause)

	return query, args, nil
}

type FlightOperationsQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId  int
	flightId int
}

func (qb *FlightOperationsQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *FlightOperationsQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *FlightOperationsQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Set SCOPE to FLIGHT
	partialQuery = fmt.Sprintf("%s and n.scope = ?", partialQuery)
	args = append(args, "FLIGHT")

	// Agent ID is required
	if qb.agentId <= 0 {
		return "", nil, fmt.Errorf("agent id is required")
	}
	partialQuery = fmt.Sprintf("%s and fa.agentmapping_id = ?", partialQuery)
	args = append(args, qb.agentId)

	if qb.flightId <= 0 {
		return "", nil, fmt.Errorf("flight id is required")
	}
	partialQuery = fmt.Sprintf("%s and fa.flight_id = ?", partialQuery)
	args = append(args, qb.flightId)

	return partialQuery, args, nil
}

func (qb *FlightOperationsQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}
	orderClause := "order by created_at asc"

	query := fmt.Sprintf(`
	SELECT 
		n.notification_id,
		n.scope,
		n.resource_id,
		n.resource_uuid,
		n.title,
		n.message,
		n.link,
		n.extra,
		n.type,
		n.popup,
		n.expires_at,
		n.created_at,
		n.seen,
		n.active
	FROM notifications n
	LEFT JOIN flight_agents fa
		ON fa.flight_id = n.resource_id
	%s %s
	`, whereClauses, orderClause)

	return query, args, nil
}

type FlightAgentsQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId  int
	flightId int
}

func (qb *FlightAgentsQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *FlightAgentsQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *FlightAgentsQueryBuilder) buildWithClauses() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Agent ID is required
	if qb.agentId <= 0 {
		return "", nil, fmt.Errorf("agent id is required")
	}
	partialQuery = fmt.Sprintf("%s and fa2.agentmapping_id = ?", partialQuery)
	args = append(args, qb.agentId)

	return partialQuery, args, nil
}

func (qb *FlightAgentsQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Agent ID is required
	if qb.flightId <= 0 {
		return "", nil, fmt.Errorf("flight id is required")
	}
	partialQuery = fmt.Sprintf("%s and fa.flight_id = ?", partialQuery)
	args = append(args, qb.flightId)

	return partialQuery, args, nil
}

func (qb *FlightAgentsQueryBuilder) BuildQuery() (string, []interface{}, error) {

	var args []interface{}

	withClauses, withArgs, err := qb.buildWithClauses()
	if err != nil {
		return "", nil, err
	}
	args = append(args, withArgs...)

	whereClauses, whereArgs, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}
	args = append(args, whereArgs...)

	orderClause := "order by email desc"

	query := fmt.Sprintf(`
		with agent_flights as (
			SELECT
				f2.flight_id
			FROM flights f2 
			LEFT JOIN flight_agents fa2 
				ON f2.flight_id = fa2.flight_id 
			%s
		)
		SELECT 
			fa.agentmapping_id as agent_id,
			ma.name,
			ma.surname,
			ma.email
		FROM agent_flights f
		LEFT JOIN flight_agents fa
			ON fa.flight_id = f.flight_id
		LEFT JOIN master_agents ma 
			ON ma.agent_id = fa.agent_id
		%s
		%s
	`, withClauses, whereClauses, orderClause)

	return query, args, nil
}

// =======================
// Functionality BUILDERS
// =======================

type UpdateFlightStatusQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId  int
	flightId int
	status   string
}

func (qb *UpdateFlightStatusQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *UpdateFlightStatusQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *UpdateFlightStatusQueryBuilder) SetStatus(status string) {
	qb.status = status
}

func (qb *UpdateFlightStatusQueryBuilder) buildSetClauses() (string, []interface{}, error) {

	partialQuery := "SET "
	args := []interface{}{}

	// Flight ID is required
	if qb.flightId <= 0 {
		return "", nil, fmt.Errorf("flight id is required")
	}
	partialQuery = fmt.Sprintf("%s status = ?", partialQuery)
	args = append(args, qb.status)

	return partialQuery, args, nil
}

func (qb *UpdateFlightStatusQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Flight ID is required
	if qb.flightId <= 0 {
		return "", nil, fmt.Errorf("flight id is required")
	}
	partialQuery = fmt.Sprintf("%s AND flight_id = ?", partialQuery)
	args = append(args, qb.flightId)

	return partialQuery, args, nil
}

func (qb *UpdateFlightStatusQueryBuilder) BuildQuery() (string, []interface{}, error) {

	var args []interface{}

	setClauses, setArgs, err := qb.buildSetClauses()
	if err != nil {
		return "", nil, err
	}
	args = append(args, setArgs...)

	whereClauses, whereArgs, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}
	args = append(args, whereArgs...)

	query := fmt.Sprintf(`
		UPDATE flights
		%s
		%s
	`, setClauses, whereClauses)

	return query, args, nil
}
