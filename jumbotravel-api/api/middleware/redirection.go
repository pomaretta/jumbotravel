package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/utils"
)

func RedirectToEndpoint(application *application.Application, to string) func(*gin.Context) {
	return func(c *gin.Context) {
		hostname := "https://pws.prod.carlospomares.com"
		if application.Environment == "DEV" {
			hostname = "https://pws.dev.carlospomares.com"
		}
		if !utils.IsWorker() {
			hostname = "http://localhost:3000"
		}
		// Redirect to Endpoint
		c.Redirect(301, hostname+to)
	}
}
