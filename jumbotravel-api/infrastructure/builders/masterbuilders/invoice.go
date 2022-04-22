package masterbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type InvoiceQueryBuilder struct {
	builders.MySQLQueryBuilder

	invoiceId          int
	bookingReferenceId string
	agentId            int
	providerId         int
}

func (qb *InvoiceQueryBuilder) SetInvoiceId(invoiceId int) {
	qb.invoiceId = invoiceId
}

func (qb *InvoiceQueryBuilder) SetBookingReferenceId(bookingReferenceId string) {
	qb.bookingReferenceId = bookingReferenceId
}

func (qb *InvoiceQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *InvoiceQueryBuilder) SetProviderId(providerId int) {
	qb.providerId = providerId
}

func (qb *InvoiceQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "WHERE 1=1"
	args := []interface{}{}

	if qb.invoiceId > 0 {
		partialQuery = fmt.Sprintf("%s AND i.id = ?", partialQuery)
		args = append(args, qb.invoiceId)
	}

	if qb.bookingReferenceId != "" {
		partialQuery = fmt.Sprintf("%s AND b.bookingreferenceid = ?", partialQuery)
		args = append(args, qb.bookingReferenceId)
	}

	if qb.agentId > 0 {
		partialQuery = fmt.Sprintf("%s AND ma4.agentmapping_id = ?", partialQuery)
		args = append(args, qb.agentId)
	}

	if qb.providerId > 0 {
		partialQuery = fmt.Sprintf("%s AND ma5.agentmapping_id = ?", partialQuery)
		args = append(args, qb.providerId)
	}

	return partialQuery, args, nil
}

func (qb *InvoiceQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClauses()
	if err != nil {
		return "", nil, err
	}

	orderClause := "ORDER BY b.bookingreferenceid, b.productcode ASC"

	query := fmt.Sprintf(`
	SELECT
		i.invoice_id,
		ma.agent_id,
		ma.dni as agent_dni,
		CONCAT(ma.name, ' ',ma.surname) as agent_name,
		ma.email as agent_email,
		ma2.agent_id as provider_id,
		ma2.dni as provider_dni,
		CONCAT(ma2.name, ' ',ma2.surname) as provider_name,
		CONCAT(ma3.common_name, ' ', '(', ma3.country, ', ', ma3.airport, ')') as provider_airport,
		ma2.email as provider_email,
		b.bookingreferenceid,
		b.productcode,
		mp.name as product_name,
		b.items,
		b.price,
		i.created_at
	FROM
		invoices i
		LEFT JOIN invoice_bookings ib ON ib.invoice_id = i.invoice_id
		INNER JOIN master_agents ma ON ma.agent_id = i.agent_id
		INNER JOIN master_agents ma2 ON ma2.agent_id = i.provider_id
		LEFT JOIN master_agentmapping ma4 ON ma4.agentmapping_id = i.agentmapping_id 
		LEFT JOIN master_agentmapping ma5 ON ma5.agentmapping_id = i.providermapping_id
		LEFT JOIN master_airports ma3 ON ma3.airport_id = ma2.airport_id 
		INNER JOIN bookings b ON b.bookingreferenceid = ib.bookingreferenceid
		INNER JOIN master_products mp ON mp.product_id = b.product_id AND mp.product_code = b.productcode
	%s
	%s
	`, whereClauses, orderClause)

	return query, args, nil
}

type PutInvoiceQueryBuilder struct {
	builders.MySQLQueryBuilder

	invoice dto.InvoiceInput
}

func (qb *PutInvoiceQueryBuilder) SetInvoice(invoice dto.InvoiceInput) {
	qb.invoice = invoice
}

func (qb *PutInvoiceQueryBuilder) BuildQuery() (string, []interface{}, error) {

	values, args, err := qb.invoice.Build()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf(`
		INSERT INTO invoices (
			agent_id,
			agentmapping_id,
			provider_id,
			providermapping_id
		) VALUES
		%s;
	`, values)

	return query, args, nil
}

type PutInvoiceBookingsQueryBuilder struct {
	builders.MySQLQueryBuilder

	bookings []dto.InvoiceBookingInput
}

func (qb *PutInvoiceBookingsQueryBuilder) SetBookings(bookings []dto.InvoiceBookingInput) {
	qb.bookings = bookings
}

func (qb *PutInvoiceBookingsQueryBuilder) BuildQuery() (string, []interface{}, error) {

	var values string
	var args []interface{}

	for idx, notification := range qb.bookings {

		valueQuery, valueArgs, err := notification.Build()
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
		INSERT INTO invoice_bookings (
			invoice_id,
			bookingreferenceid
		) VALUES
		%s;
	`, values)

	return query, args, nil
}
