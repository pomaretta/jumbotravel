package master

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Airplanes
//
// @Router /master/airplanes [get]
// @Tags Master
// @Summary Get master airplanes.
//
// @Security Bearer
// @Produce json
//
// @Param airplaneid query int false "Airplane ID"
// @Param carrier query string false "Carrier"
// @Param flightnumber query int false "Flight number"
//
// @Success 200 {object} response.JSONResult{result=[]entity.Airplane} "Get master airplanes"
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Airplanes(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		airplaneId := c.DefaultQuery("airplaneid", "-1")
		parsedAirplaneId, err := strconv.Atoi(airplaneId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "airplaneid must be an integer",
			})
			return
		}

		flightNumber := c.DefaultQuery("flightnumber", "-1")
		parsedFlightNumber, err := strconv.Atoi(flightNumber)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "flightnumber must be an integer",
			})
			return
		}

		carrier := c.Query("carrier")

		masterAirplanes, err := application.GetMasterAirplanes(parsedAirplaneId, parsedFlightNumber, carrier)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": masterAirplanes,
		})

	}
}
