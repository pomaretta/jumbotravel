package dto

import (
	"fmt"
	"time"
)

type Invoice struct {
	InvoiceId          *int       `json:"invoice_id"`
	AgentId            *int       `json:"agent_id"`
	AgentDNI           *string    `json:"agent_dni"`
	AgentName          *string    `json:"agent_name"`
	AgentEmail         *string    `json:"agent_email"`
	ProviderId         *int       `json:"provider_id"`
	ProviderDNI        *string    `json:"provider_dni"`
	ProviderName       *string    `json:"provider_name"`
	ProviderAirport    *string    `json:"provider_airport"`
	ProviderEmail      *string    `json:"provider_email"`
	BookingReferenceId *string    `json:"bookingreferenceid"`
	ProductCode        *int       `json:"product_code"`
	ProductName        *string    `json:"product_name"`
	Items              *int       `json:"items"`
	Price              *float64   `json:"price"`
	CreatedAt          *time.Time `json:"created_at"`
} // @name Invoice

func (v *Invoice) GetDestFields() []interface{} {
	return []interface{}{
		&v.InvoiceId,
		&v.AgentId,
		&v.AgentDNI,
		&v.AgentName,
		&v.AgentEmail,
		&v.ProviderId,
		&v.ProviderDNI,
		&v.ProviderName,
		&v.ProviderAirport,
		&v.ProviderEmail,
		&v.BookingReferenceId,
		&v.ProductCode,
		&v.ProductName,
		&v.Items,
		&v.Price,
		&v.CreatedAt,
	}
}

func (v *Invoice) New() {
	*v = Invoice{}
}

func (v *Invoice) Val() interface{} {
	return *v
}

type InvoiceInput struct {
	AgentId           *int       `json:"agent_id"`
	AgentMappingId    *int       `json:"agentmapping_id"`
	ProviderId        *int       `json:"provider_id"`
	ProviderMappingId *int       `json:"providermapping_id"`
	ReportDate        *time.Time `json:"report_date"`
}

func (i InvoiceInput) Build() (string, []interface{}, error) {

	partialQuery := "("
	args := make([]interface{}, 0)

	if i.AgentId == nil {
		return "", nil, fmt.Errorf("agentid is required")
	}
	partialQuery = fmt.Sprintf("%s?", partialQuery)
	args = append(args, *i.AgentId)

	if i.AgentMappingId == nil {
		return "", nil, fmt.Errorf("agentmappingid is required")
	}
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *i.AgentMappingId)

	if i.ProviderId == nil {
		return "", nil, fmt.Errorf("providerid is required")
	}
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *i.ProviderId)

	if i.ProviderMappingId == nil {
		return "", nil, fmt.Errorf("providermappingid is required")
	}
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *i.ProviderMappingId)

	if i.ReportDate != nil {
		parsedDate := *i.ReportDate
		partialQuery = fmt.Sprintf("%s,DATE(?)", partialQuery)
		args = append(args, parsedDate.Format("2006-01-02"))
	} else {
		partialQuery = fmt.Sprintf("%s,?", partialQuery)
		args = append(args, nil)
	}

	partialQuery = fmt.Sprintf("%s)", partialQuery)

	return partialQuery, args, nil
}

type InvoiceBookingInput struct {
	InvoiceId          *int    `json:"invoice_id"`
	BookingReferenceId *string `json:"bookingreferenceid"`
}

func (i InvoiceBookingInput) Build() (string, []interface{}, error) {

	partialQuery := "("
	args := make([]interface{}, 0)

	if i.InvoiceId == nil {
		return "", nil, fmt.Errorf("invoiceid is required")
	}
	partialQuery = fmt.Sprintf("%s?", partialQuery)
	args = append(args, *i.InvoiceId)

	if i.BookingReferenceId == nil {
		return "", nil, fmt.Errorf("bookingreferenceid is required")
	}
	partialQuery = fmt.Sprintf("%s,?)", partialQuery)
	args = append(args, *i.BookingReferenceId)

	return partialQuery, args, nil
}
