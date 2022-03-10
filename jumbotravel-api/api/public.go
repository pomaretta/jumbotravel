package api

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func (api *API) initPublic() {

	r := api.handler.Group("/public")

	docsGroup := r.Group("/swagger")
	{
		docsGroup.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}
