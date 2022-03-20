package middleware

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/docs/response"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/lib/rsajwt"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/utils"
)

var (
	publicResources = []string{
		"\\/public\\/.*",
		"\\/swagger.*",
	}
)

func AuthenticationMiddleware(application *application.Application) gin.HandlerFunc {
	return func(c *gin.Context) {

		if !utils.IsWorker() {
			c.Next()
			return
		}

		method := c.Request.Method
		endpoint := c.Request.URL.Path

		// Check if the endpoint is public or not
		if IsPublic(endpoint) {
			c.Next()
			return
		}

		authorizationToken := c.Request.Header.Get("Authorization")

		// TODO: Check if the token is valid
		verifier, err := rsajwt.NewVerifierFromPublicKeyFile("rsa.public")
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		splitToken := strings.Split(authorizationToken, "Bearer ")
		if len(splitToken) < 2 {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		jwtToken := splitToken[1]

		// get token claims and see if has expired
		claims, err := verifier.GetVerifiedValue(jwtToken, application.Environment)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		for _, resource := range claims.Resources {

			resourceMethod := strings.SplitN(resource, "/", 2)[0]
			resourceEndpoint := strings.SplitN(resource, "/", 2)[1]

			if !IsAllowed(method, endpoint, resourceMethod, resourceEndpoint) {
				continue
			}

			c.Next()
			return
		}

		response.Unauthorized(c)
		c.Abort()
		return

	}
}

func IsPublic(endpoint string) bool {
	for _, resource := range publicResources {
		if matched, _ := regexp.MatchString(resource, endpoint); matched {
			return true
		}
	}
	return false
}

func IsAllowed(method, endpoint, resourceMethod, resourceEndpoint string) bool {

	resourceMethod = strings.ReplaceAll(resourceMethod, "/", "\\/")
	resourceEndpoint = strings.ReplaceAll("/"+resourceEndpoint, "/", "\\/")
	resourceMethod = strings.ReplaceAll(resourceMethod, "*", ".*")
	resourceEndpoint = strings.ReplaceAll(resourceEndpoint, "*", ".*")

	// Check if the method is allowed
	methodExpression := regexp.MustCompile(resourceMethod)
	if !methodExpression.MatchString(method) {
		return false
	}

	resourceExpression := regexp.MustCompile(resourceEndpoint)
	if !resourceExpression.MatchString(endpoint) {
		return false
	}

	return true
}
