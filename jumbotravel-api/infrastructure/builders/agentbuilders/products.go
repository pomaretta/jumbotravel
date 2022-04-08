package agentbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type FlightProductsQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId  int
	flightId int
}

func (qb *FlightProductsQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *FlightProductsQueryBuilder) SetFlightId(flightId int) {
	qb.flightId = flightId
}

func (qb *FlightProductsQueryBuilder) buildWithClauses() (string, []interface{}, error) {

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

func (qb *FlightProductsQueryBuilder) buildWhereClauses() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Agent ID is required
	if qb.flightId <= 0 {
		return "", nil, fmt.Errorf("flight id is required")
	}
	partialQuery = fmt.Sprintf("%s and f.flight_id = ?", partialQuery)
	args = append(args, qb.flightId)

	partialQuery = fmt.Sprintf("%s and mp2.active = ?", partialQuery)
	args = append(args, true)

	return partialQuery, args, nil
}

func (qb *FlightProductsQueryBuilder) BuildQuery() (string, []interface{}, error) {

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

	orderClause := "ORDER BY mp2.product_code ASC"

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
			mp.productmapping_id as product_id,
			mp2.product_code,
			mp2.name,
			mp2.description,
			mp2.brand,
			mp2.type,
			as2.stock,
			mp2.max,
			mp2.saleprice
		FROM agent_flights af
		LEFT JOIN flights f
			ON f.flight_id = af.flight_id
		LEFT JOIN flight_routes fr
			ON fr.route_id = f.route_id
		LEFT JOIN master_airplanesmapping ma 
			ON ma.airplanemapping_id = fr.airplanemapping_id
		LEFT JOIN master_airplanes ma2 
			ON ma2.airplane_id = ma.airplane_id
		LEFT JOIN airplane_stock as2 -- NOTE LAST STATE OF STOCK 
			ON as2.airplanemapping_id = fr.airplanemapping_id
		LEFT JOIN airplane_stockmapping as3
			ON as3.stock_id = as2.stock_id
		LEFT JOIN master_productmapping mp 
			ON mp.productmapping_id = as2.productmapping_id 
		LEFT JOIN master_products mp2 
			ON mp2.product_id = mp.product_id
		%s
		%s
	`, withClauses, whereClauses, orderClause)

	return query, args, nil
}
