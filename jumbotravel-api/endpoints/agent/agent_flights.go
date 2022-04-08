package agent

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Flights
//
// @Router /agent/:id/flights [get]
// @Tags Agent
// @Summary Get agent flights.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param routeid query int false "Route ID"
// @Param flightid query int false "Flight ID"
// @Param airplaneid query int false "Airplane ID"
// @Param status query string false "Status"
// @Param departuretime query string false "Departure time"
// @Param arrivaltime query string false "Arrival time"
//
// @Success 200 {object} response.JSONResult{result=[]dto.AgentFlight} "Get agent flights."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Flights(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		routeId := c.DefaultQuery("routeid", "-1")
		parsedRouteId, err := strconv.Atoi(routeId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		flightId := c.DefaultQuery("flightid", "-1")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		airplaneId := c.DefaultQuery("airplaneid", "-1")
		parsedAirplaneId, err := strconv.Atoi(airplaneId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		status := c.Query("status")

		departureTime := c.Query("departuretime")
		parsedDepartureTime, err := time.Parse("2006-01-02 15:04:05", departureTime)
		if err != nil && departureTime != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		arrivalTime := c.Query("arrivaltime")
		parsedArrivalTime, err := time.Parse("2006-01-02 15:04:05", arrivalTime)
		if err != nil && arrivalTime != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := application.GetAgentFlights(parsedAgentId, parsedRouteId, parsedFlightId, parsedAirplaneId, status, parsedDepartureTime, parsedArrivalTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

// FlightDetails
//
// @Router /agent/:id/flights/:flightid/details [get]
// @Tags Agent
// @Summary Get flights details.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param flightid path int true "Flight ID"
//
// @Success 200 {object} response.JSONResult{result=[]dto.AgentFlight} "Get agent flights."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func FlightDetails(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		flightId := c.Param("flightid")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := application.GetAgentFlights(parsedAgentId, -1, parsedFlightId, -1, "", time.Time{}, time.Time{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

// FlightOperations
//
// @Router /agent/:id/flights/:flightid/operations [get]
// @Tags Agent
// @Summary Get flights operations.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param flightid path int true "Flight ID"
//
// @Success 200 {object} response.JSONResult{result=[]entity.Notification} "Get flight operations."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func FlightOperations(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		flightId := c.Param("flightid")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := application.GetAgentFlightOperations(parsedAgentId, parsedFlightId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

// FlightAgents
//
// @Router /agent/:id/flights/:flightid/agents [get]
// @Tags Agent
// @Summary Get flights agents.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param flightid path int true "Flight ID"
//
// @Success 200 {object} response.JSONResult{result=[]dto.FlightAgent} "Get flight agents."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func FlightAgents(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		flightId := c.Param("flightid")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := application.GetAgentFlightAgents(parsedAgentId, parsedFlightId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

// FlightProducts
//
// @Router /agent/:id/flights/:flightid/products [get]
// @Tags Agent
// @Summary Get flights products.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param flightid path int true "Flight ID"
//
// @Success 200 {object} response.JSONResult{result=[]dto.FlightProduct} "Get flight products."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func FlightProducts(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		flightId := c.Param("flightid")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := application.GetAgentFlightProducts(parsedAgentId, parsedFlightId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}
