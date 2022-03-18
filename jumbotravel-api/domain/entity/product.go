package entity

import "time"

type Product struct {
	ProductID   *int       `json:"product_id"`
	ProductCode *int       `json:"product_code"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Brand       *string    `json:"brand"`
	Type        *string    `json:"type"`
	Max         *int       `json:"max"`
	SalePrice   *float64   `json:"saleprice"`
	Active      *bool      `json:"active"`
	CreatedAt   *time.Time `json:"created_at"`
} // @name Product

func (v *Product) GetDestFields() []interface{} {
	return []interface{}{
		&v.ProductID,
		&v.ProductCode,
		&v.Name,
		&v.Description,
		&v.Brand,
		&v.Type,
		&v.Max,
		&v.SalePrice,
		&v.Active,
		&v.CreatedAt,
	}
}

func (v *Product) New() {
	*v = Product{}
}

func (v *Product) Val() interface{} {
	return *v
}
