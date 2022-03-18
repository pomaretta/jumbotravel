package application

import "github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"

type MySQLFetcher interface {
	FetchMasterAirports(airpotId int, country, city, airport string) ([]entity.MasterAirport, error)
	FetchMasterAgents(agentId int, dni, agentType, email string, active bool) ([]entity.Agent, error)
}

func (app *Application) GetMasterAirports(airpotId int, country, city, airport string) ([]entity.MasterAirport, error) {
	return app.MySQLFetcher.FetchMasterAirports(airpotId, country, city, airport)
}

func (app *Application) GetMasterAgents(agentId int, dni, agentType, email string, active bool) ([]entity.Agent, error) {
	return app.MySQLFetcher.FetchMasterAgents(agentId, dni, agentType, email, active)
}
