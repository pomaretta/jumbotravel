package builders

import "fmt"

type MySQLQueryBuilder struct {
	QueryBuilder
}

func (qb MySQLQueryBuilder) GetPagination() string {
	return fmt.Sprintf("LIMIT %d OFFSET %d", qb.PageSize, qb.Page*qb.PageSize)
}
