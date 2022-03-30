package agent

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/docs/response"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
)

// Agents
//
// @Router /agent/:id/data [get]
// @Tags Agent
// @Summary Get agent data.
//
// @Security Bearer
// @Produce json
//
// @Success 200 {object} response.JSONResult{result=[]response.AgentData} "Get agent data."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func AgentData(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		masterAgents, err := application.GetMasterAgents(parsedAgentId, "", "", "", true)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(masterAgents) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Agent not found",
			})
			return
		}

		// Get the unique agent data
		agent := masterAgents[0]

		airport := &entity.MasterAirport{}
		// Get the airport data if the agent is "PROVIDER" and airport is not empty
		if *agent.Type == "PROVIDER" && agent.AirportId != nil {
			airport = getAgentAiport(application, *agent.AirportId)
		}

		agentData := response.AgentData{
			AgentId:    agent.AgentID,
			Dni:        agent.DNI,
			Name:       agent.Name,
			Surname:    agent.Surname,
			Email:      agent.Email,
			Type:       agent.Type,
			AirportId:  agent.AirportId,
			Country:    airport.Country,
			City:       airport.City,
			Airport:    airport.Airport,
			CommonName: airport.CommonName,
		}
		c.JSON(200, agentData)
	}
}

func getAgentAiport(application *application.Application, airportId int) *entity.MasterAirport {

	airport, err := application.FetchMasterAirports(airportId, "", "", "")
	if err != nil {
		return &entity.MasterAirport{}
	}

	if len(airport) == 0 {
		return &entity.MasterAirport{}
	}

	return &airport[0]
}
