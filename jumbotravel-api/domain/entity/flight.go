package entity

import "time"

type MasterFlight struct {
	FlightId      *int       `json:"flight_id"`
	RouteId       *int       `json:"route_id"`
	Status        *string    `json:"status"`
	DepartureTime *time.Time `json:"departure_time"`
	ArrivalTime   *time.Time `json:"arrival_time"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
} // @name MasterFlight

func (v *MasterFlight) GetDestFields() []interface{} {
	return []interface{}{
		&v.FlightId,
		&v.RouteId,
		&v.Status,
		&v.DepartureTime,
		&v.ArrivalTime,
		&v.CreatedAt,
		&v.UpdatedAt,
	}
}

func (v *MasterFlight) New() {
	*v = MasterFlight{}
}

func (v *MasterFlight) Val() interface{} {
	return *v
}
