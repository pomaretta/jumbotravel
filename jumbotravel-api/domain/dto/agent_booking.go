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
