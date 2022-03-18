package entity

import "time"

type Airplane struct {
	AirplaneID   *int       `json:"airplane_id"`
	Carrier      *string    `json:"carrier"`
	FlightNumber *string    `json:"flight_number"`
	Seats        *int       `json:"seats"`
	CreatedAt    *time.Time `json:"created_at"`
} // @name Airplane

func (v *Airplane) GetDestFields() []interface{} {
	return []interface{}{
		&v.AirplaneID,
		&v.Carrier,
		&v.FlightNumber,
		&v.Seats,
		&v.CreatedAt,
	}
}

func (v *Airplane) New() {
	*v = Airplane{}
}

func (v *Airplane) Val() interface{} {
	return *v
}
