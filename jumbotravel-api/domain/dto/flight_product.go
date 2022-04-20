package dto

type FlightProduct struct {
	ProductID   *int     `json:"product_id"`
	ProductCode *int     `json:"product_code"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Brand       *string  `json:"brand"`
	Type        *string  `json:"type"`
	Stock       *int     `json:"stock"`
	Max         *int     `json:"max"`
	SalePrice   *float64 `json:"sale_price"`
} // @name FlightProduct

func (v *FlightProduct) GetDestFields() []interface{} {
	return []interface{}{
		&v.ProductID,
		&v.ProductCode,
		&v.Name,
		&v.Description,
		&v.Brand,
		&v.Type,
		&v.Stock,
		&v.Max,
		&v.SalePrice,
	}
}

func (v *FlightProduct) New() {
	*v = FlightProduct{}
}

func (v *FlightProduct) Val() interface{} {
	return *v
}

type BookingProduct struct {
	FlightProduct
	Quantity *int `json:"quantity"`
}
