package application

import "github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"

type MySQLFetcher interface {
	FetchMasterAirports(airpotId int, country, city, airport string) ([]entity.MasterAirport, error)
}

func (app *Application) GetMasterAirports(airpotId int, country, city, airport string) ([]entity.MasterAirport, error) {
	return app.MySQLFetcher.FetchMasterAirports(airpotId, country, city, airport)
}
