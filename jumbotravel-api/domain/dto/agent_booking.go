package dto

import (
	"fmt"
	"time"
)

type BookingAggregate struct {
	BookingReferenceId *string    `json:"bookingreferenceid"`
	Status             *string    `json:"status"`
	FlightId           *int       `json:"flight_id"`
	AgentId            *int       `json:"agent_id"`
	AgentName          *string    `json:"agent_name"`
	AgentSurname       *string    `json:"agent_surname"`
	ProviderId         *int       `json:"provider_id"`
	ProviderName       *string    `json:"provider_name"`
	ProviderSurname    *string    `json:"provider_surname"`
	Items              *int       `json:"items"`
	Total              *float64   `json:"total"`
	CreatedAt          *time.Time `json:"created_at"`
	HasInvoice         *bool      `json:"has_invoice"`
} // @name BookingAggregate

func (v *BookingAggregate) GetDestFields() []interface{} {
	return []interface{}{
		&v.BookingReferenceId,
		&v.Status,
		&v.FlightId,
		&v.AgentId,
		&v.AgentName,
		&v.AgentSurname,
		&v.ProviderId,
		&v.ProviderName,
		&v.ProviderSurname,
		&v.Items,
		&v.Total,
		&v.CreatedAt,
		&v.HasInvoice,
	}
}

func (v *BookingAggregate) New() {
	*v = BookingAggregate{}
}

func (v *BookingAggregate) Val() interface{} {
	return *v
}

type BookingItem struct {
	BookingReferenceId *string    `json:"bookingreferenceid"`
	ProductCode        *string    `json:"productcode"`
	Status             *string    `json:"status"`
	Items              *int       `json:"items"`
	Price              *float64   `json:"price"`
	Name               *string    `json:"name"`
	Brand              *string    `json:"brand"`
	SalePrice          *float64   `json:"saleprice"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
} // @name BookingItem

func (v *BookingItem) GetDestFields() []interface{} {
	return []interface{}{
		&v.BookingReferenceId,
		&v.ProductCode,
		&v.Status,
		&v.Items,
		&v.Price,
		&v.Name,
		&v.Brand,
		&v.SalePrice,
		&v.CreatedAt,
		&v.UpdatedAt,
	}
}

func (v *BookingItem) New() {
	*v = BookingItem{}
}

func (v *BookingItem) Val() interface{} {
	return *v
}

type BookingInput struct {
	BookingReferenceId *string  `json:"bookingreferenceid"`
	ProductCode        *int     `json:"productcode"`
	Status             *string  `json:"status"`
	AgentId            *int     `json:"agent_id"`
	AgentMappingId     *int     `json:"agentmapping_id"`
	ProductId          *int     `json:"product_id"`
	ProductMappingId   *int     `json:"productmapping_id"`
	FlightId           *int     `json:"flight_id"`
	Items              *int     `json:"items"`
	Price              *float64 `json:"price"`
	ProviderId         *int     `json:"provider_id"`
	ProviderMappingId  *int     `json:"providermapping_id"`
	Hash64             *int64   `json:"hash64"`
}

func (b BookingInput) Build() (string, []interface{}, error) {

	partialQuery := "("
	args := make([]interface{}, 0)

	// BookingReferenceId
	partialQuery = fmt.Sprintf("%s?", partialQuery)
	args = append(args, *b.BookingReferenceId)

	// ProductCode
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.ProductCode)

	// Status
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.Status)

	// AgentId
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.AgentId)

	// AgentMappingId
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.AgentMappingId)

	// ProductId
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.ProductId)

	// ProductMappingId
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.ProductMappingId)

	// FlightId
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.FlightId)

	// Items
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.Items)

	// Price
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, *b.Price)

	// ProviderId
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if b.ProviderId != nil {
		args = append(args, *b.ProviderId)
	} else {
		args = append(args, nil)
	}

	// ProviderMappingId
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if b.ProviderMappingId != nil {
		args = append(args, *b.ProviderMappingId)
	} else {
		args = append(args, nil)
	}

	// Hash64
	partialQuery = fmt.Sprintf("%s,?)", partialQuery)
	args = append(args, *b.Hash64)

	return partialQuery, args, nil
}

type BookingItemInput struct {
	ProductCode *int `json:"productcode"`
	Quantity    *int `json:"quantity"`
}
