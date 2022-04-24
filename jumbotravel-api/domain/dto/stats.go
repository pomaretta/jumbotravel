package dto

type BookingCount struct {
	Name  *string `json:"name"`
	Value *int    `json:"value"`
} // @name BookingCount

func (v *BookingCount) GetDestFields() []interface{} {
	return []interface{}{
		&v.Name,
		&v.Value,
	}
}

func (v *BookingCount) New() {
	*v = BookingCount{}
}

func (v *BookingCount) Val() interface{} {
	return *v
}

type BookingCompositeCount struct {
	Name     *string  `json:"name"`
	Flights  *int     `json:"flights"`
	Bookings *int     `json:"bookings"`
	Total    *float64 `json:"total"`
} // @name BookingCompositeCount

func (v *BookingCompositeCount) GetDestFields() []interface{} {
	return []interface{}{
		&v.Name,
		&v.Flights,
		&v.Bookings,
		&v.Total,
	}
}

func (v *BookingCompositeCount) New() {
	*v = BookingCompositeCount{}
}

func (v *BookingCompositeCount) Val() interface{} {
	return *v
}
