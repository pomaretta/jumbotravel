package master

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Agents
//
// @Router /master/agents [get]
// @Tags Master
// @Summary Get master agents.
//
// @Security Bearer
// @Produce json
//
// @Param agentid query int false "Agent ID"
// @Param dni query string false "DNI"
// @Param agenttype query string false "Agent type"
// @Param email query string false "Email"
// @Param active query bool false "Active"
//
// @Success 200 {object} response.JSONResult{result=[]entity.Agent} "Get master agents"
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Agents(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.DefaultQuery("agentid", "-1")
		parsedAirportId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "agentid must be an integer",
			})
			return
		}
		dni := c.Query("dni")

		agentType := c.Query("agenttype")
		email := c.Query("email")
		active := c.DefaultQuery("active", "true")
		parsedActive, err := strconv.ParseBool(active)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "active must be a boolean",
			})
			return
		}

		masterAgents, err := application.GetMasterAgents(parsedAirportId, dni, agentType, email, parsedActive)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": masterAgents,
		})

	}
}
