package response

import "github.com/gin-gonic/gin"

func Unauthorized(c *gin.Context) {
	c.JSON(401, gin.H{"error": "Unauthorized"})
}
