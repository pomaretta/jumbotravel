package mysql

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/stockbuilders"
)

func (db *MySQL) FetchStock(stockId, airplaneId, productId, productCode int) (s []dto.Stock, err error) {

	qb := &stockbuilders.StockQueryBuilder{}
	qb.SetStockID(stockId)
	qb.SetAirplaneID(airplaneId)
	qb.SetProductID(productId)
	qb.SetProductCode(productCode)

	ent, err := db.Fetch(&dto.Stock{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.Stock))
	}

	return
}
