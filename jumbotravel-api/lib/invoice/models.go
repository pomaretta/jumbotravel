package invoice

import (
	"fmt"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
)

type Assistant struct {
	Id    int
	Name  string
	DNI   string
	Email string
}

type Provider struct {
	Id      int
	Name    string
	DNI     string
	Airport string
	Email   string
}

type BookingItem struct {
	Code     int
	Name     string
	Quantity int
	Price    float64
}

func (i BookingItem) GetContent(referenceId string) [][]string {
	return [][]string{
		{referenceId, fmt.Sprintf("%s (%d)", i.Name, i.Code), fmt.Sprintf("%d", i.Quantity), fmt.Sprintf("%.2f", i.Price)},
	}
}

type Booking struct {
	ReferenceId string
	Items       []BookingItem
	Total       float64
}

func (b Booking) GetContent() (content [][]string) {
	for idx, item := range b.Items {
		referenceId := ""
		if idx == 0 {
			referenceId = b.ReferenceId
		}
		content = append(content, item.GetContent(referenceId)...)
	}
	return
}

type Invoice struct {
	Id           int
	Date         time.Time
	Signature    string
	SignatureUrl string
	Assistant    Assistant
	Provider     Provider
	Bookings     []Booking
}

func (i *Invoice) Total() (total float64) {
	for _, booking := range i.Bookings {
		total += booking.Total
	}
	return
}

func (i *Invoice) GetContent() (content [][]string) {
	for _, booking := range i.Bookings {
		content = append(content, booking.GetContent()...)
	}
	return
}

func Parse(invoice []dto.Invoice) (*Invoice, error) {

	if len(invoice) == 0 {
		return nil, fmt.Errorf("no invoice found")
	}
	metadataInvoice := invoice[0]

	// Get first metadata (invoiceid, assistant, provider)
	assistant := Assistant{
		Id:    *metadataInvoice.AgentId,
		Name:  *metadataInvoice.AgentName,
		DNI:   *metadataInvoice.AgentDNI,
		Email: *metadataInvoice.AgentEmail,
	}
	provider := Provider{
		Id:      *metadataInvoice.ProviderId,
		Name:    *metadataInvoice.ProviderName,
		DNI:     *metadataInvoice.ProviderDNI,
		Airport: *metadataInvoice.ProviderAirport,
		Email:   *metadataInvoice.ProviderEmail,
	}

	// Group invoices by referenceId
	bookings := make(map[string][]dto.Invoice)
	for _, invoice := range invoice {
		bookings[*invoice.BookingReferenceId] = append(bookings[*invoice.BookingReferenceId], invoice)
	}

	// Parse each booking
	parsedBookings := make([]Booking, 0)
	for referenceId, invoices := range bookings {
		booking := Booking{
			ReferenceId: referenceId,
			Items:       make([]BookingItem, 0),
		}
		for _, invoice := range invoices {
			booking.Items = append(booking.Items, BookingItem{
				Code:     *invoice.ProductCode,
				Name:     *invoice.ProductName,
				Quantity: *invoice.Items,
				Price:    *invoice.Price,
			})
			booking.Total = booking.Total + (*invoice.Price * float64(*invoice.Items))
		}
		parsedBookings = append(parsedBookings, booking)
	}

	return &Invoice{
		Id:        *metadataInvoice.InvoiceId,
		Date:      *metadataInvoice.CreatedAt,
		Assistant: assistant,
		Provider:  provider,
		Bookings:  parsedBookings,
	}, nil
}
