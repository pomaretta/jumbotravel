package stockbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type StockQueryBuilder struct {
	builders.MySQLQueryBuilder

	StockID     int
	AirplaneID  int
	ProductID   int
	ProductCode int
}

func (qb *StockQueryBuilder) SetStockID(stockID int) {
	qb.StockID = stockID
}

func (qb *StockQueryBuilder) SetAirplaneID(airplaneID int) {
	qb.AirplaneID = airplaneID
}

func (qb *StockQueryBuilder) SetProductID(productID int) {
	qb.ProductID = productID
}

func (qb *StockQueryBuilder) SetProductCode(productCode int) {
	qb.ProductCode = productCode
}

func (qb *StockQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.StockID > -1 {
		partialQuery += " and st.stock_id = ?"
		args = append(args, qb.StockID)
	}

	if qb.AirplaneID > -1 {
		partialQuery += " and ma.airplane_id = ?"
		args = append(args, qb.AirplaneID)
	}

	if qb.ProductID > -1 {
		partialQuery += " and mp.product_id = ?"
		args = append(args, qb.ProductID)
	}

	if qb.ProductCode > -1 {
		partialQuery += " and mp.product_code = ?"
		args = append(args, qb.ProductCode)
	}

	return partialQuery, args, nil
}

func (qb *StockQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf(`
		select
			st.stock_id,
			ma.airplane_id,
			ma.carrier,
			ma.flight_number,
			ma.seats,
			mp.product_id,
			mp.product_code,
			mp.name,
			mp.description,
			mp.brand,
			mp.type,
			mp.max,
			mp.saleprice,
			mp.active,
			st.stock,
			stm.updated_at,
			st.created_at
		from
			airplane_stockmapping stm
		left join
			airplane_stock st
			on st.stock_id = stm.stock_id
		left join
			master_airplanes ma
			on ma.airplane_id = st.airplane_id
		left join
			master_products mp 
			on mp.product_id = st.product_id
		%s
	`, whereClauses)

	return query, args, nil
}
