package master

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Airports
//
// @Router /master/airports [get]
// @Tags Master
// @Summary Get master airports.
//
// @Security Bearer
// @Produce json
//
// @Param airportId query string false "Get airport by id"
// @Param country query string false "Get airport by country"
// @Param city query string false "Get airport by city"
// @Param airport query string false "Get airport by airport"
//
// @Success 200 {object} response.JSONResult{result=[]entity.MasterAirport} "Get master airports"
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Airports(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		airpotId := c.DefaultQuery("airpotid", "-1")
		parsedAirportId, err := strconv.Atoi(airpotId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "airpotId must be an integer",
			})
			return
		}

		country := c.Query("country")
		city := c.Query("city")
		airport := c.Query("airport")

		masterAirports, err := application.GetMasterAirports(parsedAirportId, country, city, airport)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": masterAirports,
		})

	}
}
