package flight

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Route
//
// @Router /flight/route [get]
// @Tags Flight
// @Summary Get route.
//
// @Security Bearer
// @Produce json
//
// @Param routeid query int false "Route ID"
// @Param airplaneid query int false "Airplane ID"
// @Param flightid query int false "Flight ID"
// @Param departurecountry query string false "Departure country"
// @Param arrivalcountry query string false "Arrival country"
// @Param departurecity query string false "Departure city"
// @Param arrivalcity query string false "Arrival city"
// @Param departureairport query string false "Departure airport"
// @Param arrivalairport query string false "Arrival airport"
// @Param status query string false "Status"
// @Param carrier query string false "Carrier"
//
// @Success 200 {object} response.JSONResult{result=[]dto.Route} "Get route."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Route(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		routeId := c.DefaultQuery("routeid", "-1")
		parsedRouteId, err := strconv.Atoi(routeId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		airplaneId := c.DefaultQuery("airplaneid", "-1")
		parsedAirplaneId, err := strconv.Atoi(airplaneId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		flightId := c.DefaultQuery("flightid", "-1")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		departureCountry := c.Query("departurecountry")
		arrivalCountry := c.Query("arrivalcountry")
		departureCity := c.Query("departurecity")
		arrivalCity := c.Query("arrivalcity")
		departureAirport := c.Query("departureairport")
		arrivalAirport := c.Query("arrivalairport")

		status := c.Query("status")
		carrier := c.Query("carrier")

		routes, err := application.GetRoute(parsedRouteId, parsedAirplaneId, parsedFlightId, departureCountry, arrivalCountry, departureCity, arrivalCity, departureAirport, arrivalAirport, status, carrier)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": routes,
		})
	}
}
