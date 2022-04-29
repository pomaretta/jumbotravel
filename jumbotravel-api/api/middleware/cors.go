package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/utils"
)

func CorsMiddleware() gin.HandlerFunc {

	origins := []string{
		"https://jumbotravel.carlospomares.es",
		"https://api.jumbotravel.carlospomares.es",
	}
	if !utils.IsWorker() {
		origins = []string{"*"}
	}
	return cors.New(
		cors.Config{
			AllowOrigins:  origins,
			AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders: []string{"Content-Length"},
			MaxAge:        12 * time.Hour,
		},
	)
}
