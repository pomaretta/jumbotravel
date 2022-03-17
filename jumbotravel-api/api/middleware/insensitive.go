package middleware

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func InsensitiveMiddleware(c *gin.Context) {

	// 1. Parse and build new query
	queryString, err := parseQueryInsensitive(c.Request.URL.RawQuery)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	// 2. Apply changes to RawQuery
	c.Request.URL.RawQuery = queryString

	c.Next()
}

func parseQueryInsensitive(query string) (queryString string, err error) {

	var values []string

	for query != "" {
		key := query
		if i := strings.IndexAny(key, "&"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if strings.Contains(key, ";") {
			err = fmt.Errorf("invalid semicolon separator in query")
			continue
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		values = append(values, fmt.Sprintf("%s=%s", strings.ToLower(key), value))
	}

	return strings.Join(values, "&"), nil
}
