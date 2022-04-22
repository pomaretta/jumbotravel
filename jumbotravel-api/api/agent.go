package api

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/endpoints/agent"
)

func (api *API) initAgent() {

	agentGroup := api.handler.Group("/agent/:id")
	{
		agentGroup.GET("/data", agent.Data(api.application))
		agentGroup.GET("/validate", agent.Validate(api.application))

		// Notifications
		notificationsGroup := agentGroup.Group("/notifications")
		{
			notificationsGroup.GET("", agent.Notifications(api.application))
			notificationsGroup.POST("", agent.ReadNotifications(api.application))
		}

		// Flights
		flightsGroup := agentGroup.Group("/flights")
		{
			flightsGroup.GET("", agent.Flights(api.application))
			// Get flight data
			flightDataGroup := flightsGroup.Group("/:flightid")
			{
				flightDataGroup.GET("/details", agent.FlightDetails(api.application))
				flightDataGroup.GET("/operations", agent.FlightOperations(api.application))
				flightDataGroup.GET("/agents", agent.FlightAgents(api.application))
				flightDataGroup.GET("/products", agent.FlightProducts(api.application))

				// Functionality
				flightDataGroup.POST("/status", agent.UpdateFlightStatus(api.application))

			}
		}

		// Bookings
		bookingsGroup := agentGroup.Group("/bookings")
		{
			bookingsGroup.PUT("", agent.BookingCreation(api.application))
			bookingsGroup.GET("/status", agent.BookingStatus(api.application))
			bookingDataGroup := bookingsGroup.Group("/:bookingid")
			{
				bookingDataGroup.GET("/details", agent.BookingDetails(api.application))
				bookingDataGroup.GET("/operations", agent.BookingOperations(api.application))
				bookingDataGroup.GET("/items", agent.BookingItems(api.application))
				bookingDataGroup.POST("/request", agent.BookingRequest(api.application))
			}
		}

	}

}
