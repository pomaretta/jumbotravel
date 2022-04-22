package master

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Invoices
//
// @Router /master/invoices [get]
// @Tags Master
// @Summary Get master invoices.
//
// @Security Bearer
// @Produce json
//
// @Param invoiceid query int false "Invoice ID"
// @Param agentid query int false "Agent ID"
// @Param providerid query int false "Provider ID"
// @Param bookingreferenceid query string false "BookingReference ID"
//
// @Success 200 {object} response.JSONResult{result=[]dto.Invoice} "Get master invoices"
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Invoices(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		invoiceId := c.DefaultQuery("invoiceid", "0")
		parsedInvoiceId, err := strconv.Atoi(invoiceId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invoiceid must be an integer",
			})
			return
		}

		agentId := c.DefaultQuery("agentid", "0")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "agentid must be an integer",
			})
			return
		}

		providerId := c.DefaultQuery("providerid", "0")
		parsedProviderId, err := strconv.Atoi(providerId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "providerid must be an integer",
			})
			return
		}

		bookingReferenceId := c.Query("bookingreferenceid")

		invoices, err := application.GetInvoices(parsedInvoiceId, parsedAgentId, parsedProviderId, bookingReferenceId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": invoices,
		})
	}
}
