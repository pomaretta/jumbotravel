package dto

import "time"

type Route struct {
	RouteID *int `json:"route_id"`

	// Airplane
	AirplaneID   *int    `json:"airplane_id"`
	Carrier      *string `json:"carrier"`
	FlightNumber *string `json:"flight_number"`
	Seats        *int    `json:"seats"`

	// Flight
	FlightID           *int       `json:"flight_id"`
	Status             *string    `json:"status"`
	DepartureTime      *time.Time `json:"departure_time"`
	ArrivalTime        *time.Time `json:"arrival_time"`
	FlightCreation     *time.Time `json:"flight_creation"`
	FlightLastModified *time.Time `json:"flight_lastmodified"`

	// Route Data
	DepartureCountry    *string `json:"departure_country"`
	ArrivalCountry      *string `json:"arrival_country"`
	DepartureCity       *string `json:"departure_city"`
	ArrivalCity         *string `json:"arrival_city"`
	DepartureAirport    *string `json:"departure_airport"`
	ArrivalAirport      *string `json:"arrival_airport"`
	DepartureCommonName *string `json:"departure_commonname"`
	ArrivalCommonName   *string `json:"arrival_commonname"`

	// Timestamps
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
} // @name Route

func (v *Route) GetDestFields() []interface{} {
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
		&v.FlightCreation,
		&v.FlightLastModified,
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

func (v *Route) New() {
	*v = Route{}
}

func (v *Route) Val() interface{} {
	return *v
}
