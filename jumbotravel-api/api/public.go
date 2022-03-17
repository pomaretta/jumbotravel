package api

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/api/middleware"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/docs"
	_ "github.com/pomaretta/jumbotravel/jumbotravel-api/docs"
)

func (api *API) initPublic() {

	r := api.handler.Group("/public")

	docsGroup := r.Group("/swagger")
	{
		docsGroup.GET("*any", middleware.SwaggerMiddleware(api.application, docs.SwaggerInfo))
	}

}
