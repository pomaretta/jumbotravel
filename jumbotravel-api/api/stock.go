package api

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/endpoints/stock"
)

func (api *API) initStock() {

	r := api.handler.Group("/stock")

	r.GET("/latest", stock.Stock(api.application))

}
