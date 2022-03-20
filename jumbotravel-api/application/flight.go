package application

import "github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"

func (app *Application) GetRoute(routeId, airplaneId, flightId int, departureCountry, arrivalCountry, departureCity, arrivalCity, departureAirport, arrivalAirport, status, carrier string) ([]dto.Route, error) {
	return app.MySQLFetcher.FetchRoute(routeId, airplaneId, flightId, departureCountry, arrivalCountry, departureCity, arrivalCity, departureAirport, arrivalAirport, status, carrier)
}
