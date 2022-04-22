package masterbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type MasterAirportQueryBuilder struct {
	builders.MySQLQueryBuilder

	AirportID int
	Country   string
	City      string
	Airport   string
}

func (qb *MasterAirportQueryBuilder) SetAirportID(airportID int) {
	qb.AirportID = airportID
}

func (qb *MasterAirportQueryBuilder) SetCountry(country string) {
	qb.Country = country
}

func (qb *MasterAirportQueryBuilder) SetCity(city string) {
	qb.City = city
}

func (qb *MasterAirportQueryBuilder) SetAirport(airport string) {
	qb.Airport = airport
}

func (qb *MasterAirportQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.AirportID > 0 {
		partialQuery = fmt.Sprintf("%s and airport_id = ?", partialQuery)
		args = append(args, qb.AirportID)
	}

	if qb.Country != "" {
		partialQuery = fmt.Sprintf("%s and country = ?", partialQuery)
		args = append(args, qb.Country)
	}

	if qb.City != "" {
		partialQuery = fmt.Sprintf("%s and city = ?", partialQuery)
		args = append(args, qb.City)
	}

	if qb.Airport != "" {
		partialQuery = fmt.Sprintf("%s and airport = ?", partialQuery)
		args = append(args, qb.Airport)
	}

	return partialQuery, args, nil
}

func (qb *MasterAirportQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf(`
	select 
		airport_id,
		country,
		city,
		airport,
		common_name,
		created_at,
		updated_at
	from master_airports
	%s
	`, whereClauses)

	return query, args, nil
}
