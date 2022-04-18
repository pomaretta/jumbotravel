package agentbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type BookingsAggrQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId            int
	bookingReferenceId string
	flightId           int
}

func (qb *BookingsAggrQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *BookingsAggrQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *BookingsAggrQueryBuilder) SetBookingReferenceId(bookingReferenceId string) {
	qb.bookingReferenceId = bookingReferenceId
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

	if qb.bookingReferenceId != "" {
		partialQuery = fmt.Sprintf("%s AND bd.bookingreferenceid = ?", partialQuery)
		args = append(args, qb.bookingReferenceId)
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

type BookingItemsQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId            int
	bookingReferenceId string
}

func (qb *BookingItemsQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *BookingItemsQueryBuilder) SetBookingReferenceId(bookingReferenceId string) {
	qb.bookingReferenceId = bookingReferenceId
}

func (qb *BookingItemsQueryBuilder) buildWhereClauses() (string, []interface{}, error) {
	partialQuery := "WHERE 1=1"
	args := []interface{}{}

	if qb.agentId == 0 {
		return "", nil, fmt.Errorf("agent id is required")
	}
	partialQuery = fmt.Sprintf("%s AND fa.agentmapping_id = ?", partialQuery)
	args = append(args, qb.agentId)

	if qb.bookingReferenceId == "" {
		return "", nil, fmt.Errorf("booking reference id is required")
	}
	partialQuery = fmt.Sprintf("%s AND b.bookingreferenceid = ?", partialQuery)
	args = append(args, qb.bookingReferenceId)

	return partialQuery, args, nil
}

func (qb *BookingItemsQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}

	orderClause := "ORDER BY b.created_at DESC"

	query := fmt.Sprintf(`
	SELECT
		b.bookingreferenceid,
		b.productcode,
		b.status,
		b.items,
		b.price,
		mp.name,
		mp.brand,
		mp.saleprice,
		b.created_at,
		b.updated_at
	FROM
		bookings b
	LEFT JOIN flights f
		ON f.flight_id = b.flight_id
	RIGHT JOIN flight_agents fa 
		ON fa.flight_id = f.flight_id
	LEFT JOIN master_products mp 
		ON mp.product_id = b.product_id AND mp.product_code = b.productcode
	%s
	%s
	`, whereClauses, orderClause)

	return query, args, nil
}
