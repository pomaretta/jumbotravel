package dto

import "time"

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
