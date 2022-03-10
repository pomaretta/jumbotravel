package builders

type Builder interface {
	GetPagination() string
	BuildQuery() (string, []interface{}, error)
}

type QueryBuilder struct {
	Page, PageSize int
}

func (qb *QueryBuilder) SetPage(page, pageSize int) {
	qb.Page, qb.PageSize = page, pageSize
}
