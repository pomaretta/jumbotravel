package dto

import "fmt"

type StockInput struct {
	AirplaneId        int `json:"airplane_id"`
	AirplaneMappingId int `json:"airplanemapping_id"`
	ProductId         int `json:"product_id"`
	ProductMappingId  int `json:"productmapping_id"`
	Stock             int `json:"stock"`
}

func (s StockInput) Build() (string, []interface{}, error) {

	partialQuery := "("
	args := make([]interface{}, 0)

	partialQuery = fmt.Sprintf("%s?", partialQuery)
	args = append(args, s.AirplaneId)

	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, s.AirplaneMappingId)

	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, s.ProductId)

	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	args = append(args, s.ProductMappingId)

	partialQuery = fmt.Sprintf("%s,?)", partialQuery)
	args = append(args, s.Stock)

	return partialQuery, args, nil
}
