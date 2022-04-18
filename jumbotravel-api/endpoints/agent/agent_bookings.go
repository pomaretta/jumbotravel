package agent

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// BookingStatus
//
// @Router /agent/:id/bookings/status [get]
// @Tags Agent
// @Summary Get agent bookings.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param flightid query int false "Flight ID"
//
// @Success 200 {object} response.JSONResult{result=[]dto.BookingAggregate} "Get agent bookings."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func BookingStatus(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		flightId := c.DefaultQuery("flightid", "0")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := application.GetAgentBookingsAggregate(parsedAgentId, parsedFlightId)
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

// BookingDetails
//
// @Router /agent/:id/bookings/:bookingid/details [get]
// @Tags Agent
// @Summary Get booking details.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param bookingid path string true "Booking Reference ID"
//
// @Success 200 {object} response.JSONResult{result=dto.BookingAggregate} "Get booking details."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func BookingDetails(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		bookingReferenceId := c.Param("bookingid")
		if bookingReferenceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Booking reference ID is required.",
			})
			return
		}

		result, err := application.GetAgentBookingDetails(parsedAgentId, bookingReferenceId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

// BookingOperations
//
// @Router /agent/:id/bookings/:bookingid/operations [get]
// @Tags Agent
// @Summary Get booking operations.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param bookingid path string true "Booking Reference ID"
//
// @Success 200 {object} response.JSONResult{result=[]entity.Notification} "Get booking operations."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func BookingOperations(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		bookingReferenceId := c.Param("bookingid")
		if bookingReferenceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Booking reference ID is required.",
			})
			return
		}

		result, err := application.GetAgentBookingOperations(parsedAgentId, bookingReferenceId)
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

// BookingItems
//
// @Router /agent/:id/bookings/:bookingid/items [get]
// @Tags Agent
// @Summary Get booking items.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param bookingid path string true "Booking Reference ID"
//
// @Success 200 {object} response.JSONResult{result=[]dto.BookingItem} "Get booking items."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func BookingItems(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		bookingReferenceId := c.Param("bookingid")
		if bookingReferenceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Booking reference ID is required.",
			})
			return
		}

		result, err := application.GetAgentBookingItems(parsedAgentId, bookingReferenceId)
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
