package agent

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/lib/invoice"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/lib/rsajwt"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/utils"
)

func ObtainPDF(application *application.Application) func(*gin.Context) {
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
				"error": "booking reference id is required",
			})
			return
		}

		signature := c.Query("signature")
		if signature == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "signature is required",
			})
			return
		}

		verifier, err := rsajwt.NewVerifierFromPublicKeyFile("rsa.public")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		claims, err := verifier.GetVerifiedValue(signature, application.Environment, false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		parsedInvoiceId, err := strconv.Atoi(claims.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		parsedSignatureAgentId, err := strconv.Atoi(claims.AllowPolicy.Extra["agent"])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		parsedSignatureProviderId, err := strconv.Atoi(claims.AllowPolicy.Extra["provider"])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		invoices, err := application.GetInvoices(parsedInvoiceId, parsedSignatureAgentId, parsedSignatureProviderId, bookingReferenceId, time.Time{}, time.Time{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(invoices) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "invoice not found",
			})
			return
		}

		parsedInvoice, err := invoice.Parse(invoices)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		schema := "https"
		hostname := "api.jumbotravel.carlospomares.es"
		if application.Environment == "DEV" {
			hostname = "api.jumbotravel.dev.carlospomares.es"
		}
		if !utils.IsWorker() {
			schema = "http"
			hostname = "localhost:3000"
		}
		parsedInvoice.SignatureUrl = fmt.Sprintf("%s://%s/agent/%d/bookings/%s/invoice?signature=%s", schema, hostname, parsedAgentId, bookingReferenceId, signature)

		creator := invoice.New(parsedInvoice)
		res, err := creator.Create()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Header("Content-Type", "application/pdf")
		res.Buffer.WriteTo(c.Writer)
	}
}

func PutInvoice(app *application.Application) func(*gin.Context) {
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
				"error": "booking reference id is required",
			})
			return
		}

		schema := "https"
		if !utils.IsWorker() {
			schema = "http"
		}

		invoices, err := app.GetInvoices(0, parsedAgentId, 0, bookingReferenceId, time.Time{}, time.Time{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(invoices) > 0 {
			parsedInvoice, err := invoice.Parse(invoices)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			signature, err := signInvoice(parsedInvoice)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			parsedInvoice.Signature = signature
			parsedInvoice.SignatureUrl = fmt.Sprintf("%s://%s/agent/%d/bookings/%s/invoice?signature=%s", schema, c.Request.Host, parsedAgentId, bookingReferenceId, signature)

			creator := invoice.New(parsedInvoice)
			res, err := creator.Create()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.Header("Content-Type", "application/pdf")
			res.Buffer.WriteTo(c.Writer)
			return
		}

		// TODO: Get booking details
		bookingDetails, err := app.GetAgentBookingDetails(parsedAgentId, agentType, bookingReferenceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if bookingDetails == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "booking not found",
			})
			return
		}

		if *bookingDetails.Status != "COMPLETED" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "booking not completed",
			})
			return
		}

		invoiceInput := dto.InvoiceInput{
			AgentId:           bookingDetails.AgentId,
			AgentMappingId:    bookingDetails.AgentId,
			ProviderId:        bookingDetails.ProviderId,
			ProviderMappingId: bookingDetails.ProviderId,
		}
		// TODO: Put Invoice and retrieve invoice id
		invoiceId, err := app.RegisterInvoice(invoiceInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		bookingInput := dto.InvoiceBookingInput{
			BookingReferenceId: bookingDetails.BookingReferenceId,
			InvoiceId:          utils.Int(int(invoiceId)),
		}
		_, err = app.RegisterInvoiceBookings([]dto.InvoiceBookingInput{bookingInput})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		invoices, err = app.GetInvoices(0, parsedAgentId, 0, bookingReferenceId, time.Time{}, time.Time{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(invoices) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "invoice not found",
			})
			return
		}

		parsedInvoice, err := invoice.Parse(invoices)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Sign the invoice
		signature, err := signInvoice(parsedInvoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		parsedInvoice.Signature = signature
		parsedInvoice.SignatureUrl = fmt.Sprintf("%s://%s/agent/%d/bookings/%s/invoice?signature=%s", schema, c.Request.Host, parsedAgentId, bookingReferenceId, signature)

		creator := invoice.New(parsedInvoice)
		res, err := creator.Create()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Header("Content-Type", "application/pdf")
		res.Buffer.WriteTo(c.Writer)
	}
}

func InvoiceReport(app *application.Application) func(*gin.Context) {
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
		// Day of report
		day := c.Query("day")
		parsedDay, err := time.Parse("2006-01-02", day)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid day",
			})
			return
		}

		if agentType != "PROVIDER" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid agent type",
			})
			return
		}

		if parsedDay.Equal(time.Now().UTC().Truncate(24 * time.Hour)) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invoice cannot be generated for today",
			})
			return
		}

		// TODO: Check if invoice is already generated
		invoices, err := app.GetInvoices(0, parsedAgentId, 0, "", parsedDay, parsedDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(invoices) > 0 {
			parsedInvoice, err := invoice.Parse(invoices)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			signature, err := signInvoice(parsedInvoice)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			parsedInvoice.Signature = signature
			creator := invoice.New(parsedInvoice)
			res, err := creator.Create()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.Header("Content-Type", "application/pdf")
			res.Buffer.WriteTo(c.Writer)
			return
		}

		// Obtain all bookings by day
		bookings, err := app.GetAgentBookingsAggregateWithDays(parsedAgentId, agentType, 0, 0, parsedDay, parsedDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(bookings) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "bookings not found",
			})
			return
		}

		invoiceInput := dto.InvoiceInput{
			AgentId:           &parsedAgentId,
			AgentMappingId:    &parsedAgentId,
			ProviderId:        &parsedAgentId,
			ProviderMappingId: &parsedAgentId,
			ReportDate:        &parsedDay,
		}
		invoiceId, err := app.RegisterInvoice(invoiceInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		var invoiceBookings []dto.InvoiceBookingInput
		for _, booking := range bookings {
			invoiceBookings = append(invoiceBookings, dto.InvoiceBookingInput{
				InvoiceId:          utils.Int(int(invoiceId)),
				BookingReferenceId: booking.BookingReferenceId,
			})
		}
		_, err = app.RegisterInvoiceBookings(invoiceBookings)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		invoices, err = app.GetInvoices(int(invoiceId), parsedAgentId, 0, "", parsedDay, parsedDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(invoices) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "invoice not found",
			})
			return
		}

		parsedInvoice, err := invoice.Parse(invoices)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Sign the invoice
		signature, err := signInvoice(parsedInvoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		parsedInvoice.Signature = signature

		creator := invoice.New(parsedInvoice)
		res, err := creator.Create()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Header("Content-Type", "application/pdf")
		res.Buffer.WriteTo(c.Writer)
	}
}

func signInvoice(invoice *invoice.Invoice) (string, error) {

	file, err := os.Open("rsa.private")
	if err != nil {
		return "", err
	}

	privatePEMData, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	block, rest := pem.Decode(privatePEMData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", errors.New("failed to decode PEM block containing private key " + block.Type)
	}

	if len(rest) != 0 {
		return "", errors.New("passed public key contains more than just the private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	rsaJWT := jwt.SigningMethodRS256
	claims := &rsajwt.Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        fmt.Sprintf("%d", invoice.Id),
			Subject:   fmt.Sprintf("%d", invoice.Id),
			ExpiresAt: invoice.Date.Unix(),
			IssuedAt:  invoice.Date.Unix(),
			Issuer:    "jumbotravel",
		},
		AllowPolicy: rsajwt.AllowPolicy{
			Extra: map[string]string{
				"agent":    fmt.Sprintf("%d", invoice.Assistant.Id),
				"provider": fmt.Sprintf("%d", invoice.Provider.Id),
			},
		},
	}

	token := jwt.NewWithClaims(rsaJWT, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return ss, nil
}
