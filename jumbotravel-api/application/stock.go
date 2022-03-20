package application

import "github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"

func (app *Application) GetStock(stockId, airplaneId, productId, productCode int) ([]dto.Stock, error) {
	return app.MySQLFetcher.FetchStock(stockId, airplaneId, productId, productCode)
}
