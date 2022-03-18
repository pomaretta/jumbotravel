package masterbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type AirplaneQueryBuilder struct {
	builders.MySQLQueryBuilder

	AirplaneID   int
	Carrier      string
	FlightNumber int
}

func (qb *AirplaneQueryBuilder) SetAirplaneID(airplaneID int) {
	qb.AirplaneID = airplaneID
}

func (qb *AirplaneQueryBuilder) SetCarrier(carrier string) {
	qb.Carrier = carrier
}

func (qb *AirplaneQueryBuilder) SetFlightNumber(flightNumber int) {
	qb.FlightNumber = flightNumber
}

func (qb *AirplaneQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.AirplaneID > -1 {
		partialQuery = fmt.Sprintf("%s and ag.airplane_id = ?", partialQuery)
		args = append(args, qb.AirplaneID)
	}

	if qb.Carrier != "" {
		partialQuery = fmt.Sprintf("%s and ag.carrier = ?", partialQuery)
		args = append(args, qb.Carrier)
	}

	if qb.FlightNumber > -1 {
		partialQuery = fmt.Sprintf("%s and ag.flight_number = ?", partialQuery)
		args = append(args, qb.FlightNumber)
	}

	return partialQuery, args, nil
}

func (qb *AirplaneQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf(`
	select 
		ag.airplane_id,
		ag.carrier,
		ag.flight_number,
		ag.seats,
		ag.created_at
	from master_airplanesmapping map
	left join master_airplanes ag
		on map.airplane_id = ag.airplane_id
	%s
	`, whereClauses)

	return query, args, nil
}
