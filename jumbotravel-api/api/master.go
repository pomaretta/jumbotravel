package api

import "github.com/pomaretta/jumbotravel/jumbotravel-api/endpoints/master"

func (api *API) initMaster() {

	r := api.handler.Group("/master")

	r.GET("/airports", master.Airports(api.application))

}
