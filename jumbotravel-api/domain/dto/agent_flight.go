package dto

import "time"

type AgentFlight struct {
	RouteID             *int       `json:"route_id"`
	DepartureCountry    *string    `json:"departure_country"`
	DepartureCity       *string    `json:"departure_city"`
	ArrivalCountry      *string    `json:"arrival_country"`
	ArrivalCity         *string    `json:"arrival_city"`
	DepartureAirport    *string    `json:"departure_airport"`
	ArrivalAirport      *string    `json:"arrival_airport"`
	DepartureCommonName *string    `json:"departure_commonname"`
	ArrivalCommonName   *string    `json:"arrival_commonname"`
	AirplaneID          *int       `json:"airplane_id"`
	Carrier             *string    `json:"carrier"`
	FlightNumber        *string    `json:"flight_number"`
	Seats               *int       `json:"seats"`
	FlightID            *int       `json:"flight_id"`
	Status              *string    `json:"status"`
	DepartureTime       *time.Time `json:"departure_time"`
	ArrivalTime         *time.Time `json:"arrival_time"`
	FlightCreated       *time.Time `json:"flight_created"`
	FlightUpdated       *time.Time `json:"flight_updated"`
	UpdatedAt           *time.Time `json:"updated_at"`
	CreatedAt           *time.Time `json:"created_at"`
} // @name AgentFlight

func (v *AgentFlight) GetDestFields() []interface{} {
	return []interface{}{
		&v.RouteID,
		&v.AirplaneID,
		&v.Carrier,
		&v.FlightNumber,
		&v.Seats,
		&v.FlightID,
		&v.Status,
		&v.DepartureTime,
		&v.ArrivalTime,
		&v.FlightCreated,
		&v.FlightUpdated,
		&v.DepartureCountry,
		&v.ArrivalCountry,
		&v.DepartureCity,
		&v.ArrivalCity,
		&v.DepartureAirport,
		&v.ArrivalAirport,
		&v.DepartureCommonName,
		&v.ArrivalCommonName,
		&v.UpdatedAt,
		&v.CreatedAt,
	}
}

func (v *AgentFlight) New() {
	*v = AgentFlight{}
}

func (v *AgentFlight) Val() interface{} {
	return *v
}

type FlightAgent struct {
	AgentID *int    `json:"agent_id"`
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Email   *string `json:"email"`
} // @name FlightAgent

func (v *FlightAgent) GetDestFields() []interface{} {
	return []interface{}{
		&v.AgentID,
		&v.Name,
		&v.Surname,
		&v.Email,
	}
}

func (v *FlightAgent) New() {
	*v = FlightAgent{}
}

func (v *FlightAgent) Val() interface{} {
	return *v
}
