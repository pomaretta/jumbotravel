package mysql

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/flightbuilders"
)

func (db *MySQL) FetchRoute(routeId, airplaneId, flightId int, departureCountry, arrivalCountry, departureCity, arrivalCity, departureAirport, arrivalAirport, status, carrier string) (s []dto.Route, err error) {

	qb := &flightbuilders.RouteQueryBuilder{}
	qb.SetRouteID(routeId)
	qb.SetAirplaneID(airplaneId)
	qb.SetFlightID(flightId)
	qb.SetDepartureCountry(departureCountry)
	qb.SetArrivalCountry(arrivalCountry)
	qb.SetDepartureCity(departureCity)
	qb.SetArrivalCity(arrivalCity)
	qb.SetDepartureAirport(departureAirport)
	qb.SetArrivalAirport(arrivalAirport)
	qb.SetStatus(status)
	qb.SetCarrier(carrier)

	ent, err := db.Fetch(&dto.Route{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(dto.Route))
	}

	return
}
