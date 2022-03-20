package flightbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type RouteQueryBuilder struct {
	builders.MySQLQueryBuilder

	RouteID          int
	AirplaneID       int
	FlightID         int
	DepartureCountry string
	ArrivalCountry   string
	DepartureCity    string
	ArrivalCity      string
	DepartureAirport string
	ArrivalAirport   string
	Status           string
	Carrier          string
}

func (qb *RouteQueryBuilder) SetRouteID(routeID int) {
	qb.RouteID = routeID
}

func (qb *RouteQueryBuilder) SetAirplaneID(airplaneID int) {
	qb.AirplaneID = airplaneID
}

func (qb *RouteQueryBuilder) SetFlightID(flightID int) {
	qb.FlightID = flightID
}

func (qb *RouteQueryBuilder) SetDepartureCountry(departureCountry string) {
	qb.DepartureCountry = departureCountry
}

func (qb *RouteQueryBuilder) SetArrivalCountry(arrivalCountry string) {
	qb.ArrivalCountry = arrivalCountry
}

func (qb *RouteQueryBuilder) SetDepartureCity(departureCity string) {
	qb.DepartureCity = departureCity
}

func (qb *RouteQueryBuilder) SetArrivalCity(arrivalCity string) {
	qb.ArrivalCity = arrivalCity
}

func (qb *RouteQueryBuilder) SetDepartureAirport(departureAirport string) {
	qb.DepartureAirport = departureAirport
}

func (qb *RouteQueryBuilder) SetArrivalAirport(arrivalAirport string) {
	qb.ArrivalAirport = arrivalAirport
}

func (qb *RouteQueryBuilder) SetStatus(status string) {
	qb.Status = status
}

func (qb *RouteQueryBuilder) SetCarrier(carrier string) {
	qb.Carrier = carrier
}

func (qb *RouteQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.RouteID > -1 {
		partialQuery = fmt.Sprintf("%s and fr.route_id = ?", partialQuery)
		args = append(args, qb.RouteID)
	}

	if qb.AirplaneID > -1 {
		partialQuery = fmt.Sprintf("%s and ma.airplane_id = ?", partialQuery)
		args = append(args, qb.AirplaneID)
	}

	if qb.FlightID > -1 {
		partialQuery = fmt.Sprintf("%s and f.flight_id = ?", partialQuery)
		args = append(args, qb.FlightID)
	}

	if qb.DepartureCountry != "" {
		partialQuery = fmt.Sprintf("%s and fr.departure_country = ?", partialQuery)
		args = append(args, qb.DepartureCountry)
	}

	if qb.ArrivalCountry != "" {
		partialQuery = fmt.Sprintf("%s and fr.arrival_country = ?", partialQuery)
		args = append(args, qb.ArrivalCountry)
	}

	if qb.DepartureCity != "" {
		partialQuery = fmt.Sprintf("%s and fr.departure_city = ?", partialQuery)
		args = append(args, qb.DepartureCity)
	}

	if qb.ArrivalCity != "" {
		partialQuery = fmt.Sprintf("%s and fr.arrival_city = ?", partialQuery)
		args = append(args, qb.ArrivalCity)
	}

	if qb.DepartureAirport != "" {
		partialQuery = fmt.Sprintf("%s and fr.departure_airport = ?", partialQuery)
		args = append(args, qb.DepartureAirport)
	}

	if qb.ArrivalAirport != "" {
		partialQuery = fmt.Sprintf("%s and fr.arrival_airport = ?", partialQuery)
		args = append(args, qb.ArrivalAirport)
	}

	if qb.Status != "" {
		partialQuery = fmt.Sprintf("%s and f.status = ?", partialQuery)
		args = append(args, qb.Status)
	}

	if qb.Carrier != "" {
		partialQuery = fmt.Sprintf("%s and ma.carrier = ?", partialQuery)
		args = append(args, qb.Carrier)
	}

	return partialQuery, args, nil
}

func (qb *RouteQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf(`
		select
			fr.route_id,
			ma.airplane_id,
			ma.carrier,
			ma.flight_number,
			ma.seats,
			f.flight_id,
			f.status,
			f.departure_time,
			f.arrival_time,
			f.created_at as flight_creation,
			f.updated_at as flight_lastmodified,
			fr.departure_country,
			fr.arrival_country,
			fr.departure_city,
			fr.arrival_city,
			fr.departure_airport,
			fr.arrival_airport,
			depair.common_name as departure_commonname,
			arrair.common_name as arrival_commoname,
			fr.updated_at,
			fr.created_at
		from
			flight_routes fr 
		right join
			flights f 
			on f.route_id = fr.route_id 
		left join master_airplanesmapping mam
			on mam.airplanemapping_id = fr.airplanemapping_id
		left join master_airplanes ma 
			on ma.airplane_id = mam.airplanemapping_id
		inner join master_airports depair
			on depair.country = fr.departure_country and depair.city = fr.departure_city and depair.airport = fr.departure_airport
		inner join master_airports arrair
			on arrair.country = fr.arrival_country and arrair.city = fr.arrival_city and arrair.airport = fr.arrival_airport
		%s
	`, whereClauses)

	return query, args, nil
}
