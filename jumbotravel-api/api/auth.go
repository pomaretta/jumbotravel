package api

import "github.com/pomaretta/jumbotravel/jumbotravel-api/endpoints/auth"

func (api *API) initAuth() {

	authGroup := api.handler.Group("/auth")
	{
		authGroup.POST("/login", auth.Login(api.application))
	}

}
