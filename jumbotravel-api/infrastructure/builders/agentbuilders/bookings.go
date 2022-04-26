package agentbuilders

import (
	"fmt"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type BookingsAggrQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId            int
	agentType          string
	bookingReferenceId string
	flightId           int
	airplaneId         int
	from, to           time.Time
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

func (qb *BookingsAggrQueryBuilder) SetAgentType(agentType string) {
	qb.agentType = agentType
}

func (qb *BookingsAggrQueryBuilder) SetAirplaneId(airplaneId int) {
	qb.airplaneId = airplaneId
}

func (qb *BookingsAggrQueryBuilder) SetFrom(from time.Time) {
	qb.from = from
}

func (qb *BookingsAggrQueryBuilder) SetTo(to time.Time) {
	qb.to = to
}

func (qb *BookingsAggrQueryBuilder) buildWhereClauses() (string, []interface{}, error) {
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

	if qb.bookingReferenceId != "" {
		partialQuery = fmt.Sprintf("%s AND bd.bookingreferenceid = ?", partialQuery)
		args = append(args, qb.bookingReferenceId)
	}

	partialQuery = fmt.Sprintf("%s AND bl.lt = bd.created_at", partialQuery)

	if !qb.from.IsZero() {
		partialQuery = fmt.Sprintf("%s AND DATE(bd.created_at) >= DATE(?)", partialQuery)
		args = append(args, qb.from.Format("2006-01-02"))
	}

	if !qb.to.IsZero() {
		partialQuery = fmt.Sprintf("%s AND DATE(bd.created_at) <= DATE(?)", partialQuery)
		args = append(args, qb.to.Format("2006-01-02"))
	}

	return partialQuery, args, nil
}

func (qb *BookingsAggrQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
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
		bd.created_at,
		NOT(ISNULL(ib.invoice_id)) as has_invoice
	FROM bookingdetail bd
	LEFT JOIN bookinglatest bl
		ON bl.bookingreferenceid = bd.bookingreferenceid
	LEFT JOIN invoice_bookings ib
		ON ib.bookingreferenceid = bd.bookingreferenceid
	%s
	%s
	%s
	`, providerClauses, whereClauses, orderClause)

	return query, args, nil
}

type BookingItemsQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId            int
	agentType          string
	bookingReferenceId string
}

func (qb *BookingItemsQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *BookingItemsQueryBuilder) SetAgentType(agentType string) {
	qb.agentType = agentType
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
	agentColumn := "fa.agentmapping_id"
	if qb.agentType == "PROVIDER" {
		agentColumn = "ma3.agentmapping_id"
	}
	partialQuery = fmt.Sprintf("%s AND %s = ?", partialQuery, agentColumn)
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

	joinClause := `
		RIGHT JOIN flight_agents fa 
			ON fa.flight_id = f.flight_id
	`
	if qb.agentType == "PROVIDER" {
		joinClause = `
		LEFT JOIN flight_routes fr ON fr.route_id = f.route_id
		LEFT JOIN master_airports ma ON ma.airport = fr.arrival_airport
		LEFT JOIN master_agents ma2 ON ma2.airport_id = ma.airport_id
		LEFT JOIN master_agentmapping ma3 ON ma3.agent_id = ma2.agent_id
		`
	}

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
	%s
	LEFT JOIN master_products mp 
		ON mp.product_id = b.product_id AND mp.product_code = b.productcode
	%s
	%s
	`, joinClause, whereClauses, orderClause)

	return query, args, nil
}

type PutBookingQueryBuilder struct {
	builders.MySQLQueryBuilder

	bookings []dto.BookingInput
}

func (qb *PutBookingQueryBuilder) SetBookings(bookings []dto.BookingInput) {
	qb.bookings = bookings
}

func (qb *PutBookingQueryBuilder) BuildQuery() (string, []interface{}, error) {

	var values string
	var args []interface{}

	for idx, booking := range qb.bookings {

		valueQuery, valueArgs, err := booking.Build()
		if err != nil {
			return "", nil, err
		}

		if idx == 0 {
			values = valueQuery
		} else {
			values = fmt.Sprintf("%s, %s", values, valueQuery)
		}

		args = append(args, valueArgs...)
	}

	query := fmt.Sprintf(`
		INSERT INTO bookings (
			bookingreferenceid,
			productcode,
			status,
			agent_id,
			agentmapping_id,
			product_id,
			productmapping_id,
			flight_id,
			items,
			price,
			provider_id,
			providermapping_id,
			hash64
		) VALUES
		%s;
	`, values)

	return query, args, nil
}

type UpdateBookingQueryBuilder struct {
	builders.MySQLQueryBuilder

	bookingReferenceId string
	status             string
	providerId         int
	providerMappingId  int
}

func (qb *UpdateBookingQueryBuilder) SetBookingReferenceId(bookingReferenceId string) {
	qb.bookingReferenceId = bookingReferenceId
}

func (qb *UpdateBookingQueryBuilder) SetStatus(status string) {
	qb.status = status
}

func (qb *UpdateBookingQueryBuilder) SetProviderId(providerId int) {
	qb.providerId = providerId
}

func (qb *UpdateBookingQueryBuilder) SetProviderMappingId(providerMappingId int) {
	qb.providerMappingId = providerMappingId
}

func (qb *UpdateBookingQueryBuilder) buildSetClauses() (string, []interface{}, error) {

	partialQuery := "SET"
	args := []interface{}{}

	if qb.status == "" {
		return "", nil, fmt.Errorf("status is required")
	}

	partialQuery = fmt.Sprintf("%s status = ?", partialQuery)
	args = append(args, qb.status)

	if qb.providerId == 0 {
		return "", nil, fmt.Errorf("provider id is required")
	}

	partialQuery = fmt.Sprintf("%s, provider_id = ?", partialQuery)
	args = append(args, qb.providerId)

	if qb.providerMappingId == 0 {
		return "", nil, fmt.Errorf("provider mapping id is required")
	}

	partialQuery = fmt.Sprintf("%s, providermapping_id = ?", partialQuery)
	args = append(args, qb.providerMappingId)

	return partialQuery, args, nil
}

func (qb *UpdateBookingQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "WHERE 1=1"
	args := []interface{}{}

	if qb.bookingReferenceId == "" {
		return "", nil, fmt.Errorf("booking reference id is required")
	}

	partialQuery = fmt.Sprintf("%s AND bookingreferenceid = ?", partialQuery)
	args = append(args, qb.bookingReferenceId)

	return partialQuery, args, nil
}

func (qb *UpdateBookingQueryBuilder) BuildQuery() (string, []interface{}, error) {

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
	UPDATE bookings
	%s
	%s
	`, setClauses, whereClauses)

	return query, args, nil
}

type UpdateProductStockQueryBuilder struct {
	builders.MySQLQueryBuilder

	products []dto.StockInput
}

func (qb *UpdateProductStockQueryBuilder) SetProducts(products []dto.StockInput) {
	qb.products = products
}

func (qb *UpdateProductStockQueryBuilder) BuildQuery() (string, []interface{}, error) {

	var values string
	var args []interface{}

	for idx, product := range qb.products {

		valueQuery, valueArgs, err := product.Build()
		if err != nil {
			return "", nil, err
		}

		if idx == 0 {
			values = valueQuery
		} else {
			values = fmt.Sprintf("%s, %s", values, valueQuery)
		}

		args = append(args, valueArgs...)
	}

	query := fmt.Sprintf(`
		INSERT INTO airplane_stock (
			airplane_id,
			airplanemapping_id,
			product_id,
			productmapping_id,
			stock
		) VALUES
		%s;
	`, values)

	return query, args, nil
}
