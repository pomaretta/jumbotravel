package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/utils"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/swag"
	"golang.org/x/net/webdav"
)

func SwaggerMiddleware(application *application.Application, spec *swag.Spec, confs ...func(c *ginSwagger.Config)) gin.HandlerFunc {

	spec.Schemes = []string{"https"}
	spec.Version = application.Version
	switch application.Environment {
	case "PROD":
		spec.Host = "pws.prod.carlospomares.com"
	case "DEV":
		spec.Host = "pws.dev.carlospomares.com"
	}

	if !utils.IsWorker() {
		spec.Host = "localhost:3000"
		spec.Schemes = []string{"http"}
	}

	return ginSwagger.WrapHandler(&webdav.Handler{
		FileSystem: swaggerFiles.FS,
		LockSystem: webdav.NewMemLS(),
	}, confs...)
}
