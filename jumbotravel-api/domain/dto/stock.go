package dto

import "time"

type Stock struct {
	StockID *int `json:"stock_id"`

	// Airplane
	AirplaneID   *int    `json:"airplane_id"`
	Carrier      *string `json:"carrier"`
	FlightNumber *string `json:"flight_number"`
	Seats        *int    `json:"seats"`

	// Product
	ProductID   *int     `json:"product_id"`
	ProductCode *int     `json:"product_code"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Brand       *string  `json:"brand"`
	Type        *string  `json:"type"`
	Max         *int     `json:"max"`
	SalePrice   *float64 `json:"saleprice"`
	Active      *bool    `json:"active"`

	// Stock Data
	Stock     *int       `json:"stock"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
} // @name Stock

func (v *Stock) GetDestFields() []interface{} {
	return []interface{}{
		&v.StockID,
		&v.AirplaneID,
		&v.Carrier,
		&v.FlightNumber,
		&v.Seats,
		&v.ProductID,
		&v.ProductCode,
		&v.Name,
		&v.Description,
		&v.Brand,
		&v.Type,
		&v.Max,
		&v.SalePrice,
		&v.Active,
		&v.Stock,
		&v.UpdatedAt,
		&v.CreatedAt,
	}
}

func (v *Stock) New() {
	*v = Stock{}
}

func (v *Stock) Val() interface{} {
	return *v
}
