package infrastructure

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type DB interface {
	getSlice(domain.Entity, string, ...interface{}) ([]interface{}, error)
	getIntValue(string, ...interface{}) (int, error)
	Fetch(domain.Entity, builders.Builder) ([]interface{}, error)
	FetchCount(builders.Builder) (int, error)
}
