package agent

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/utils"
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

		agentType := c.GetString("subtype")
		flightId := c.DefaultQuery("flightid", "0")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := application.GetAgentBookingsAggregate(parsedAgentId, agentType, parsedFlightId)
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

		agentType := c.GetString("subtype")
		bookingReferenceId := c.Param("bookingid")
		if bookingReferenceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Booking reference ID is required.",
			})
			return
		}

		result, err := application.GetAgentBookingDetails(parsedAgentId, agentType, bookingReferenceId)
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

		agentType := c.GetString("subtype")
		bookingReferenceId := c.Param("bookingid")
		if bookingReferenceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Booking reference ID is required.",
			})
			return
		}

		result, err := application.GetAgentBookingItems(parsedAgentId, agentType, bookingReferenceId)
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

// BookingCreation
//
// @Router /agent/:id/bookings [put]
// @Tags Agent
// @Summary Create a new booking.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param flightid query string true "Flight ID"
//
// @Success 200 {object} response.JSONResult{result=string} "Get booking reference id."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func BookingCreation(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		flightId := c.Query("flightid")
		parsedFlightId, err := strconv.Atoi(flightId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Flight ID is required",
			})
			return
		}

		// TODO: Check input data from body.
		var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "json body is required.",
			})
			return
		}

		// TODO: Check if there's a items in the body.
		items, ok := body["items"].([]interface{})
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "items are required.",
			})
			return
		}
		products, err := parseInputItems(items)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// ====================
		// TODO: Booking data
		// ====================

		// TODO: Check if the flights is available for the agent.
		flights, err := application.FetchAgentFlights(parsedAgentId, 0, parsedFlightId, 0, "", time.Time{}, time.Time{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(flights) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "flight is not available for the agent",
			})
			return
		}
		flight := flights[0]

		// TODO: Check if flight has booking created.
		if *flight.HasBooking {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "flight has booking created",
			})
			return
		}

		// TODO: Check if the status of the flight is legal.
		if *flight.Status != "FLYING" && *flight.Status != "ARRIVAL" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "flight status is not legal",
			})
			return
		}

		// TODO: Custom fetcher for obtaining product data and current stock by flight.
		bookedProducts, err := application.FetchAgentFlightProducts(parsedAgentId, parsedFlightId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		parsedBookedProducts, err := parseBookedProducts(products, bookedProducts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// TODO: Get agent details
		agents, err := application.FetchMasterAgents(parsedAgentId, "", "", "", true)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(agents) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "agent is not available",
			})
			return
		}
		agent := agents[0]

		// ====================
		// Request
		// ====================

		// TODO: Generate booking reference id
		referenceId := uuid.NewString()

		// Generate request inputs
		var requestInputs []dto.BookingInput
		// TODO: Fill each product with quantity
		for _, product := range parsedBookedProducts {
			input := dto.BookingInput{
				BookingReferenceId: utils.String(referenceId),
				ProductCode:        product.ProductCode,
				Status:             utils.String("PENDING"),
				AgentId:            agent.AgentID,
				AgentMappingId:     agent.AgentID,
				ProductId:          product.ProductID,
				ProductMappingId:   product.ProductID,
				FlightId:           flight.FlightID,
				Items:              product.Quantity,
				Price:              product.SalePrice,
			}
			hash, err := hashstructure.Hash(input, hashstructure.FormatV2, nil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			input.Hash64 = utils.Int64(int64(hash))
			requestInputs = append(requestInputs, input)
		}

		// TODO: Create booking
		_, err = application.PutBookingCreation(requestInputs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		bookingNotification := dto.NotificationInput{
			Scope:            utils.String("BOOKING"),
			ResourceUuid:     utils.String(referenceId),
			Title:            utils.String("Booking created"),
			Message:          utils.String("Booking created"),
			NotificationType: utils.String("INFO"),
			ExpiresAt:        utils.Time(time.Now().Add(time.Hour * 24)),
		}

		flightNotification := dto.NotificationInput{
			Scope:            utils.String("FLIGHT"),
			ResourceId:       flight.FlightID,
			Title:            utils.String("New booking"),
			Message:          utils.String("Order created with id"),
			NotificationType: utils.String("INFO"),
			ExpiresAt:        utils.Time(time.Now().Add(time.Hour * 24)),
			Extra: &map[string]string{
				"agent":   fmt.Sprintf("%s %s", *agent.Name, *agent.Surname),
				"agentid": fmt.Sprintf("%d", *agent.AgentID),
				"booking": referenceId,
			},
		}

		agentNotification := dto.NotificationInput{
			Scope:            utils.String("AGENT"),
			ResourceId:       agent.AgentID,
			Title:            utils.String("Booking created successfully"),
			NotificationType: utils.String("SUCCESS"),
			Link:             utils.String(fmt.Sprintf("/bookings/%s", referenceId)),
			ExpiresAt:        utils.Time(time.Now().Add(time.Hour * 1)),
		}

		// TODO: Provider notification
		notifications := []dto.NotificationInput{
			bookingNotification,
			flightNotification,
			agentNotification,
		}

		_, err = application.PutNotifications(notifications)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": referenceId,
		})
	}
}

func parseInputItems(items []interface{}) ([]dto.BookingItemInput, error) {
	var result []dto.BookingItemInput
	// TODO: Check if the items are valid.
	for _, item := range items {
		// Parse item as json map
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("items are required")
		}
		// Check if the item has a product_code
		parsedProductCode, ok := itemMap["product_code"].(float64)
		if !ok {
			return nil, fmt.Errorf("product_code is required")
		}
		// Check if the item has a quantity
		parsedProductQuantity, ok := itemMap["quantity"].(float64)
		if !ok {
			return nil, fmt.Errorf("quantity is required")
		}
		if parsedProductQuantity <= 0 {
			return nil, fmt.Errorf("quantity must be greater than 0")
		}

		result = append(result, dto.BookingItemInput{
			ProductCode: utils.Int(int(parsedProductCode)),
			Quantity:    utils.Int(int(parsedProductQuantity)),
		})
	}
	return result, nil
}

func parseBookedProducts(inputProducts []dto.BookingItemInput, booked []dto.FlightProduct) ([]dto.BookingProduct, error) {
	var result []dto.BookingProduct
	for _, inputProduct := range inputProducts {
		for _, bookedProduct := range booked {
			if *bookedProduct.ProductCode == *inputProduct.ProductCode {

				// // TODO: Check if the quantity is legal.
				if *bookedProduct.Stock+*inputProduct.Quantity > *bookedProduct.Max {
					return nil, fmt.Errorf("illegal quantity")
				}

				result = append(result, dto.BookingProduct{
					FlightProduct: bookedProduct,
					Quantity:      inputProduct.Quantity,
				})
				break
			}
		}
	}
	return result, nil
}

// BookingRequest
//
// @Router /agent/:id/bookings/:bookingid/request [post]
// @Tags Agent
// @Summary Request review of booking.
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param bookingid path string true "Booking Reference ID"
//
// @Success 200 {object} response.JSONResult{result=string} "Successfull request."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func BookingRequest(application *application.Application) func(*gin.Context) {
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

		agentType := c.GetString("subtype")
		if agentType != "ASSISTANT" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "agent is not an assistant",
			})
			return
		}

		// TODO: Get booking airport
		bookingDetails, err := application.GetAgentBookingDetails(parsedAgentId, agentType, bookingReferenceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// TODO: Get flight details
		flightDetails, err := application.GetAgentFlights(
			parsedAgentId, 0, *bookingDetails.FlightId, 0, "", time.Time{}, time.Time{},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(flightDetails) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "flight not found",
			})
			return
		}
		flight := flightDetails[0]
		// TODO: Check if the flight has booking
		if !*flight.HasBooking {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "flight has no booking",
			})
			return
		}

		airports, err := application.GetMasterAirports(
			0, *flight.ArrivalCountry, *flight.ArrivalCity, *flight.ArrivalAirport,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(airports) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "arrival airport not found",
			})
			return
		}
		airport := airports[0]

		notificationId := []int{*airport.AirportID}
		notificationUuid := []string{*bookingDetails.BookingReferenceId}
		notificationScope := []string{"AIRPORT"}
		// TODO: Check if there's a current notification for the request
		notifications, err := application.GetNotifications(
			nil, notificationId, notificationUuid, nil, notificationScope, "0", "0", "0", "0",
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(notifications) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "There's a notification for this request",
			})
			return
		}

		notification := dto.NotificationInput{
			Scope:            utils.String("AIRPORT"),
			ResourceId:       airport.AirportID,
			ResourceUuid:     bookingDetails.BookingReferenceId,
			Title:            utils.String(fmt.Sprintf("Booking request for %d", *flight.FlightID)),
			Message:          utils.String(fmt.Sprintf("Booking request for %d", *flight.FlightID)),
			Link:             utils.String(fmt.Sprintf("/bookings/%s", bookingReferenceId)),
			NotificationType: utils.String("INFO"),
			ExpiresAt:        utils.Time(time.Now().UTC().Add(time.Hour * 2)),
		}
		bookingNotification := dto.NotificationInput{
			Scope:            utils.String("BOOKING"),
			ResourceUuid:     bookingDetails.BookingReferenceId,
			Title:            utils.String(fmt.Sprintf("Requested review for %s airport providers", *airport.CommonName)),
			Message:          utils.String(fmt.Sprintf("Requested review for %s airport providers", *airport.CommonName)),
			NotificationType: utils.String("INFO"),
			ExpiresAt:        utils.Time(time.Now().UTC().Add(time.Hour * 2)),
		}
		_, err = application.PutNotifications([]dto.NotificationInput{notification, bookingNotification})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "booking request sent",
		})
	}
}

// BookingComplete
//
// @Router /agent/:id/bookings/:bookingid/complete [post]
// @Tags Agent
// @Summary Change booking status to complete
//
// @Security Bearer
// @Produce json
//
// @Param id path int true "Agent ID"
// @Param bookingid path string true "Booking Reference ID"
//
// @Success 200 {object} response.JSONResult{result=string} "Successfull request."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func BookingComplete(app *application.Application) func(*gin.Context) {
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

		agentType := c.GetString("subtype")
		if agentType != "PROVIDER" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "agent is not a provider",
			})
			return
		}

		// TODO: Get booking airport
		bookingDetails, err := app.GetAgentBookingDetails(parsedAgentId, agentType, bookingReferenceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if bookingDetails == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "booking not found",
			})
			return
		}

		if *bookingDetails.Status == "COMPLETED" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "booking already completed",
			})
			return
		}

		flightDetails, err := app.GetAgentFlights(
			*bookingDetails.AgentId, 0, *bookingDetails.FlightId, 0, "", time.Time{}, time.Time{},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(flightDetails) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "flight not found",
			})
			return
		}
		flight := flightDetails[0]

		bookingItems, err := app.GetAgentBookingItems(parsedAgentId, agentType, bookingReferenceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(bookingItems) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "booking items not found",
			})
			return
		}

		currentStock, err := app.GetAgentFlightProducts(*bookingDetails.AgentId, *bookingDetails.FlightId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		parsedStock, err := parseCurrentStockItems(bookingItems, currentStock, *flight.AirplaneID, *flight.AirplaneID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = app.UpdateStock(parsedStock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// TODO: Update booking status
		_, err = app.UpdateBookingStatus(bookingReferenceId, "COMPLETED", parsedAgentId, parsedAgentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Mark booking as completed
		bookingNotification := dto.NotificationInput{
			Scope:            utils.String("BOOKING"),
			ResourceUuid:     utils.String(bookingReferenceId),
			Title:            utils.String("Booking completed, stock updated"),
			Message:          utils.String("Booking completed, stock updated"),
			NotificationType: utils.String("SUCCESS"),
			ExpiresAt:        utils.Time(time.Now().Add(time.Hour * 24)),
		}
		agentNotification := dto.NotificationInput{
			Scope:            utils.String("AGENT"),
			ResourceId:       bookingDetails.AgentId,
			Title:            utils.String(fmt.Sprintf("Booking %s completed", bookingReferenceId)),
			NotificationType: utils.String("SUCCESS"),
			Link:             utils.String(fmt.Sprintf("/bookings/%s", bookingReferenceId)),
			ExpiresAt:        utils.Time(time.Now().Add(time.Hour * 1)),
		}

		_, err = app.PutNotifications([]dto.NotificationInput{bookingNotification, agentNotification})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// TODO: Register notification for booking complete
		c.JSON(http.StatusOK, gin.H{
			"result": "booking completed",
		})
	}
}

func parseCurrentStockItems(inputProducts []dto.BookingItem, current []dto.FlightProduct, airplaneId, airplaneMappingId int) (s []dto.StockInput, err error) {
	for _, inputProduct := range inputProducts {
		for _, currentProduct := range current {
			parsedProductCode, err := strconv.Atoi(*inputProduct.ProductCode)
			if err != nil {
				return nil, err
			}
			if *currentProduct.ProductCode != parsedProductCode {
				continue
			}

			// Check if quantity is legal
			if *inputProduct.Items+*currentProduct.Stock > *currentProduct.Max {
				return nil, errors.New("quantity is not legal")
			}

			s = append(s, dto.StockInput{
				AirplaneId:        airplaneId,
				AirplaneMappingId: airplaneMappingId,
				ProductId:         *currentProduct.ProductID,
				ProductMappingId:  *currentProduct.ProductID,
				Stock:             (*inputProduct.Items + *currentProduct.Stock),
			})
		}
	}
	return
}
