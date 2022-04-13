package agentbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type BookingsAggrQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId  int
	flightId int
}

func (qb *BookingsAggrQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *BookingsAggrQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *BookingsAggrQueryBuilder) buildWhereClauses() (string, []interface{}, error) {
	partialQuery := "WHERE 1=1"
	args := []interface{}{}

	if qb.agentId == 0 {
		return "", args, fmt.Errorf("agent id is required")
	}
	partialQuery = fmt.Sprintf("%s AND bd.agent_id = ?", partialQuery)
	args = append(args, qb.agentId)

	if qb.flightId != 0 {
		partialQuery = fmt.Sprintf("%s AND bd.flight_id = ?", partialQuery)
		args = append(args, qb.flightId)
	}

	partialQuery = fmt.Sprintf("%s AND bl.lt = bd.created_at", partialQuery)

	return partialQuery, args, nil
}

func (qb *BookingsAggrQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}

	orderClause := "ORDER BY bd.created_at DESC"

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
				ROUND(bs.total, 2) as total
			FROM bookings as b
			LEFT JOIN bookingsummary bs
				ON b.bookingreferenceid = bs.bookingreferenceid
			LEFT JOIN master_agents ma 
				ON ma.agent_id = b.agent_id
			LEFT JOIN master_agents ma2
				ON ma2.agent_id = b.provider_id
		)
	SELECT 
		bd.bookingreferenceid,
		bd.status,
		bd.flight_id,
		bd.agent_id,
		bd.agent_name,
		bd.agent_surname,
		bd.provider_id,
		bd.provider_name,
		bd.provider_surname,
		bd.items,
		bd.total,
		bd.created_at
	FROM bookingdetail bd
	LEFT JOIN bookinglatest bl
		ON bl.bookingreferenceid = bd.bookingreferenceid
	%s
	%s
	`, whereClauses, orderClause)

	return query, args, nil
}
