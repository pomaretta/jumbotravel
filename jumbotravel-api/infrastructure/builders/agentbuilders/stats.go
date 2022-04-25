package agentbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type BookingCountQueryBuilder struct {
	builders.MySQLQueryBuilder

	countType  string
	agentId    int
	agentType  string
	flightId   int
	airplaneId int
	days       int
}

func (qb *BookingCountQueryBuilder) SetCountType(countType string) {
	qb.countType = countType
}

func (qb *BookingCountQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *BookingCountQueryBuilder) SetAgentType(agentType string) {
	qb.agentType = agentType
}

func (qb *BookingCountQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *BookingCountQueryBuilder) SetAirplaneId(airplaneId int) {
	qb.airplaneId = airplaneId
}

func (qb *BookingCountQueryBuilder) SetDays(days int) {
	qb.days = days
}

func (qb *BookingCountQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "WHERE 1=1"
	args := []interface{}{}

	if qb.agentId == 0 {
		return "", args, fmt.Errorf("agent id is required")
	}
	agentColumn := "bd.agent_id"
	if qb.agentType == "PROVIDER" {
		agentColumn = "ma3.agentmapping_id"
	}
	partialQuery = fmt.Sprintf("%s AND %s = ?", partialQuery, agentColumn)
	args = append(args, qb.agentId)

	if qb.flightId != 0 {
		partialQuery = fmt.Sprintf("%s AND bd.flight_id = ?", partialQuery)
		args = append(args, qb.flightId)
	}

	if qb.airplaneId != 0 {
		partialQuery = fmt.Sprintf("%s AND bd.airplane_id = ?", partialQuery)
		args = append(args, qb.airplaneId)
	}

	if qb.days > 0 {
		partialQuery = fmt.Sprintf("%s AND DATE(bd.created_at) BETWEEN DATE_SUB(CURRENT_DATE, INTERVAL ? DAY) AND CURRENT_DATE", partialQuery)
		args = append(args, qb.days)
	}

	partialQuery = fmt.Sprintf("%s AND bl.lt = bd.created_at", partialQuery)

	return partialQuery, args, nil
}

func (qb *BookingCountQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}

	nameColumn := "DATE(bd.created_at) as name,"
	if qb.countType == "STATUS" {
		nameColumn = "bd.status as name,"
	}

	providerClauses := ""
	if qb.agentType == "PROVIDER" {
		providerClauses = `
			LEFT JOIN flights f 
				ON f.flight_id = bd.flight_id
			LEFT JOIN flight_routes fr 
				ON fr.route_id = f.route_id
			LEFT JOIN master_airports ma 
				ON ma.airport = fr.arrival_airport
			LEFT JOIN master_agents ma2 
				ON ma2.airport_id = ma.airport_id
			LEFT JOIN master_agentmapping ma3 
				ON ma3.agent_id = ma2.agent_id
		`
	}

	query := fmt.Sprintf(`
	with
		bookinglatest as (
			SELECT
				b.bookingreferenceid,
				max(b.created_at) as lt
			FROM bookings b
			GROUP BY 1
		),
		bookingsummary as (
			SELECT
				b.bookingreferenceid,
				SUM(b.items) as items,
				SUM(b.items * b.price) as total
			FROM bookings b 
			LEFT JOIN master_agents ma 
				ON ma.agent_id = b.agent_id
			LEFT JOIN master_products mp 
				ON mp.product_id = b.product_id
			GROUP BY 1
		),
		bookingdetail as (
			SELECT
				b.bookingreferenceid,
				b.productcode,
				b.status as status,
				b.flight_id,
				b.agentmapping_id as agent_id,
				ma.name as agent_name,
				ma.surname as agent_surname,
				b.providermapping_id as provider_id,
				ma2.name as provider_name,
				ma2.surname as provider_surname,
				b.created_at,
				bs.items,
				ROUND(bs.total, 2) as total,
				fr2.airplanemapping_id as airplane_id
			FROM bookings as b
			LEFT JOIN bookingsummary bs
				ON b.bookingreferenceid = bs.bookingreferenceid
			LEFT JOIN master_agents ma 
				ON ma.agent_id = b.agent_id
			LEFT JOIN master_agents ma2
				ON ma2.agent_id = b.provider_id
			LEFT JOIN flights f2 
				ON f2.flight_id = b.flight_id
			LEFT JOIN flight_routes fr2
				ON fr2.route_id = f2.route_id
		)
	SELECT DISTINCT
		%s
		COUNT(DISTINCT bd.bookingreferenceid) as value
	FROM bookingdetail bd
	LEFT JOIN bookinglatest bl
		ON bl.bookingreferenceid = bd.bookingreferenceid
	LEFT JOIN invoice_bookings ib
		ON ib.bookingreferenceid = bd.bookingreferenceid
	%s
	%s
	GROUP BY 1
	`, nameColumn, providerClauses, whereClauses)

	return query, args, nil
}

type BookingCompositeQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId    int
	agentType  string
	flightId   int
	airplaneId int
	days       int
}

func (qb *BookingCompositeQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *BookingCompositeQueryBuilder) SetAgentType(agentType string) {
	qb.agentType = agentType
}

func (qb *BookingCompositeQueryBuilder) SetFlightId(flightId int) {

	qb.flightId = flightId
}

func (qb *BookingCompositeQueryBuilder) SetAirplaneId(airplaneId int) {
	qb.airplaneId = airplaneId
}

func (qb *BookingCompositeQueryBuilder) SetDays(days int) {
	qb.days = days
}

func (qb *BookingCompositeQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "WHERE 1=1"
	args := []interface{}{}

	if qb.agentId == 0 {
		return "", args, fmt.Errorf("agent id is required")
	}
	agentColumn := "bd.agent_id"
	if qb.agentType == "PROVIDER" {
		agentColumn = "ma3.agentmapping_id"
	}
	partialQuery = fmt.Sprintf("%s AND %s = ?", partialQuery, agentColumn)
	args = append(args, qb.agentId)

	if qb.flightId != 0 {
		partialQuery = fmt.Sprintf("%s AND bd.flight_id = ?", partialQuery)
		args = append(args, qb.flightId)
	}

	if qb.airplaneId != 0 {
		partialQuery = fmt.Sprintf("%s AND bd.airplane_id = ?", partialQuery)
		args = append(args, qb.airplaneId)
	}

	if qb.days > 0 {
		partialQuery = fmt.Sprintf("%s AND DATE(c.created_at) BETWEEN DATE_SUB(CURRENT_DATE, INTERVAL ? DAY) AND CURRENT_DATE", partialQuery)
		args = append(args, qb.days)
	}

	return partialQuery, args, nil
}

func (qb *BookingCompositeQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}

	providerClauses := ""
	if qb.agentType == "PROVIDER" {
		providerClauses = `
			LEFT JOIN flight_routes fr 
				ON fr.route_id = f.route_id
			LEFT JOIN master_airports ma 
				ON ma.airport = fr.arrival_airport
			LEFT JOIN master_agents ma2 
				ON ma2.airport_id = ma.airport_id
			LEFT JOIN master_agentmapping ma3 
				ON ma3.agent_id = ma2.agent_id
		`
	}

	query := fmt.Sprintf(`
	with
		bookinglatest as (
			SELECT
				b.bookingreferenceid,
				max(b.created_at) as lt
			FROM bookings b
			GROUP BY 1
		),
		bookingsummary as (
			SELECT
				b.bookingreferenceid,
				SUM(b.items) as items,
				SUM(b.items * b.price) as total
			FROM bookings b 
			LEFT JOIN master_agents ma 
				ON ma.agent_id = b.agent_id
			LEFT JOIN master_products mp 
				ON mp.product_id = b.product_id
			GROUP BY 1
		),
		bookingdetail as (
			SELECT
				b.bookingreferenceid,
				b.productcode,
				b.status as status,
				b.flight_id,
				b.agentmapping_id as agent_id,
				ma.name as agent_name,
				ma.surname as agent_surname,
				b.providermapping_id as provider_id,
				ma2.name as provider_name,
				ma2.surname as provider_surname,
				b.created_at,
				bs.items,
				ROUND(bs.total, 2) as total,
				fr2.airplanemapping_id as airplane_id
			FROM bookings as b
			LEFT JOIN bookingsummary bs
				ON b.bookingreferenceid = bs.bookingreferenceid
			LEFT JOIN master_agents ma 
				ON ma.agent_id = b.agent_id
			LEFT JOIN master_agents ma2
				ON ma2.agent_id = b.provider_id
			LEFT JOIN flights f2 
				ON f2.flight_id = b.flight_id
			LEFT JOIN flight_routes fr2
				ON fr2.route_id = f2.route_id
		),
		composite as (
			SELECT DISTINCT 
				DATE(created_at) as created_at
			FROM bookingdetail bd
			GROUP BY 1
			UNION ALL
			SELECT DISTINCT
				DATE(created_at) as created_at
			FROM flights f3
			GROUP BY 1
		)
	SELECT DISTINCT 
		DATE(c.created_at) as name,
		COUNT(DISTINCT f.flight_id) as flights,
		COUNT(DISTINCT bd.bookingreferenceid) as bookings,
		SUM(DISTINCT bd.total) as total
	FROM composite c
	LEFT JOIN bookingdetail bd
		ON DATE(bd.created_at) = DATE(c.created_at)
	LEFT JOIN flights f
		ON DATE(f.created_at) = DATE(c.created_at)
	%s
	%s
	GROUP BY 1
	`, providerClauses, whereClauses)

	return query, args, nil
}
