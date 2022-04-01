package agent

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Validate
//
// @Router /agent/:id/validate [get]
// @Tags Agent
// @Summary Check if session is valid.
//
// @Security Bearer
// @Produce json
//
// @Success 200 {object} response.JSONResult{result=string} "Success"
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Validate(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": fmt.Sprintf("agent %d is valid", parsedAgentId),
		})
	}
}
