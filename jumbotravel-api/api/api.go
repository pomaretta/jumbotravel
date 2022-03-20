package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/api/middleware"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

type API struct {
	handler *gin.Engine

	application *application.Application
}

func New(application *application.Application) *API {

	api := &API{
		handler:     gin.Default(),
		application: application,
	}

	mode := gin.DebugMode
	if application.Environment == "PROD" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	api.initMiddlewares()
	api.initRoutes()

	return api
}

func (api *API) initRoutes() {

	api.initPublic()
	api.initMaster()

}

func (api *API) initMiddlewares() {

	// Cors
	api.handler.Use(middleware.CorsMiddleware())

	// Authorizer Middleware
	api.handler.Use(middleware.AuthenticationMiddleware(api.application))

	// Case Insentivive
	api.handler.Use(middleware.InsensitiveMiddleware)

}

func (api *API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	api.handler.ServeHTTP(w, req)
}
