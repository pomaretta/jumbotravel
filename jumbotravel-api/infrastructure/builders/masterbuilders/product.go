package masterbuilders

import (
	"fmt"
	"strings"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type MasterProductQueryBuilder struct {
	builders.MySQLQueryBuilder

	ProductID   []int
	ProductCode []int
}

func (qb *MasterProductQueryBuilder) SetProductID(productID []int) {
	qb.ProductID = productID
}

func (qb *MasterProductQueryBuilder) SetProductCode(productCode []int) {
	qb.ProductCode = productCode
}

func (qb *MasterProductQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if len(qb.ProductID) > 0 {
		partialQuery = fmt.Sprintf("%s and mp.product_id in (", partialQuery)
		productQuery := ""
		for _, v := range qb.ProductID {
			productQuery = fmt.Sprintf("%s,?", productQuery)
			args = append(args, v)
		}
		productQuery = strings.TrimPrefix(productQuery, ",")
		partialQuery = fmt.Sprintf("%s%s)", partialQuery, productQuery)
	}

	if len(qb.ProductCode) > 0 {
		partialQuery = fmt.Sprintf("%s and mp.product_code in (", partialQuery)
		productQuery := ""
		for _, v := range qb.ProductCode {
			productQuery = fmt.Sprintf("%s,?", productQuery)
			args = append(args, v)
		}
		productQuery = strings.TrimPrefix(productQuery, ",")
		partialQuery = fmt.Sprintf("%s%s)", partialQuery, productQuery)
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
