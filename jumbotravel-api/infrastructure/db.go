package infrastructure

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type DB interface {
	getSlice(domain.Data, string, ...interface{}) ([]interface{}, error)
	getIntValue(string, ...interface{}) (int, error)
	Fetch(domain.Data, builders.Builder) ([]interface{}, error)
	FetchCount(builders.Builder) (int, error)
}
