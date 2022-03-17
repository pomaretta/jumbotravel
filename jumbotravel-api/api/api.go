package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	gin.SetMode(gin.DebugMode)

	api.initRoutes()

	return api
}

func (api *API) initRoutes() {

	api.initPublic()
	api.initMaster()

}

func (api *API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	api.handler.ServeHTTP(w, req)
}
