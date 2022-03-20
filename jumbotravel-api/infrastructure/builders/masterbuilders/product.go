package masterbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type MasterProductQueryBuilder struct {
	builders.MySQLQueryBuilder

	ProductID   int
	ProductCode int
}

func (qb *MasterProductQueryBuilder) SetProductID(productID int) {
	qb.ProductID = productID
}

func (qb *MasterProductQueryBuilder) SetProductCode(productCode int) {
	qb.ProductCode = productCode
}

func (qb *MasterProductQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.ProductID > -1 {
		partialQuery = fmt.Sprintf("%s and mp.product_id = ?", partialQuery)
		args = append(args, qb.ProductID)
	}

	if qb.ProductCode > -1 {
		partialQuery = fmt.Sprintf("%s and mp.product_code = ?", partialQuery)
		args = append(args, qb.ProductCode)
	}

	return partialQuery, args, nil
}

func (qb *MasterProductQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf(`
		select 
			mp.product_id,
			mp.product_code,
			mp.name,
			mp.description,
			mp.brand,
			mp.type,
			mp.max,
			mp.saleprice,
			mp.active,
			mp.created_at
		from master_productmapping map
		left join master_products mp 
			on map.product_id = mp.product_id
		%s
	`, whereClauses)

	return query, args, nil
}
