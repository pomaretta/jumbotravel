package api

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/api/middleware"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/docs"
	_ "github.com/pomaretta/jumbotravel/jumbotravel-api/docs"
)

func (api *API) initPublic() {

	r := api.handler
	publicGroup := r.Group("/public")

	docsGroup := publicGroup.Group("/swagger")
	{
		docsGroup.GET("*any", middleware.SwaggerMiddleware(api.application, docs.SwaggerInfo))
	}

	r.GET("/swagger", middleware.RedirectToEndpoint(api.application, "/public/swagger/index.html"))

}
