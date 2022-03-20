package api

import "github.com/pomaretta/jumbotravel/jumbotravel-api/endpoints/flight"

func (api *API) initFlight() {

	r := api.handler.Group("/flight")

	// Route group
	routeGropup := r.Group("/route")
	{
		routeGropup.GET("", flight.Route(api.application))
	}

}
