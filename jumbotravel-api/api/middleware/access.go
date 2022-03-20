package middleware

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/lib/rsajwt"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func AccessLoggingMiddleware(app *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {
		// Get the method
		method := c.Request.Method
		// Get the path
		path := c.Request.URL.Path
		// Get the query string
		query := c.Request.URL.RawQuery

		// Generate the request id
		requestId := uuid.New().String()

		// Parse token if exists
		token := c.GetHeader("Authorization")
		tokenId := "-"
		tokenName := ""
		if token != "" {
			// Get the token data if exists
			tokenId, tokenName = getTokenData(token)
		}

		// Get the ip
		ip := c.ClientIP()

		// Write the Request ID into the response header
		c.Writer.Header().Set("X-Request-Id", requestId)
		// Write the token id into the response header
		c.Writer.Header().Set("X-Token-Id", tokenId)

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// Let the request continue
		c.Next()

		body := w.body.String()

		// Get the status code
		status := c.Writer.Status()

		// If status code is error, then get the error message
		errorMessage := ""
		if status >= 400 {
			// Parse the body JSON to get the error message
			var data map[string]interface{}
			err := json.Unmarshal([]byte(body), &data)
			if err == nil && data["error"] != nil {
				errorMessage = data["error"].(string)
			}
		}

		// Put the access logging
		_ = app.PutAccessLogging(
			requestId,
			tokenId,
			tokenName,
			ip,
			method,
			path,
			query,
			errorMessage,
			status,
		)

	}
}

func getTokenData(token string) (string, string) {
	// Split the token in three parts
	parts := strings.Split(token, ".")

	claimsBytes, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return "", ""
	}

	dec := json.NewDecoder(bytes.NewBuffer(claimsBytes))

	c := rsajwt.Claims{}
	err = dec.Decode(&c)
	if err != nil {
		return "", ""
	}

	return c.Id, c.Subject
}
