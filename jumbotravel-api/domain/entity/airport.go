package entity

type MasterAirport struct {
	AirportID  *int    `json:"airport_id"`
	Country    *string `json:"country"`
	City       *string `json:"city"`
	Airport    *string `json:"airport"`
	CommonName *string `json:"common_name"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
} // @name MasterAirport

func (v *MasterAirport) GetDestFields() []interface{} {
	return []interface{}{
		&v.AirportID,
		&v.Country,
		&v.City,
		&v.Airport,
		&v.CommonName,
		&v.CreatedAt,
		&v.UpdatedAt,
	}
}

func (v *MasterAirport) New() {
	*v = MasterAirport{}
}

func (v *MasterAirport) Val() interface{} {
	return *v
}
