package mysql

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/masterbuilders"
)

func (db *MySQL) FetchMasterAirports(airpotId int, country, city, airport string) (s []entity.MasterAirport, err error) {

	qb := &masterbuilders.MasterAirportQueryBuilder{}
	qb.SetAirportID(airpotId)
	qb.SetCountry(country)
	qb.SetCity(city)
	qb.SetAirport(airport)

	ent, err := db.Fetch(&entity.MasterAirport{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.MasterAirport))
	}

	return
}

func (db *MySQL) FetchMasterAgents(agentId int, dni, agentType, email string, active bool) (s []entity.Agent, err error) {

	qb := &masterbuilders.MasterAgentQueryBuilder{}
	qb.SetAgentID(agentId)
	qb.SetDNI(dni)
	qb.SetType(agentType)
	qb.SetActive(active)
	qb.SetEmail(email)

	ent, err := db.Fetch(&entity.Agent{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.Agent))
	}

	return
}

func (db *MySQL) FetchMasterAirplanes(airplaneId, flightNumber int, carrier string) (s []entity.Airplane, err error) {

	qb := &masterbuilders.AirplaneQueryBuilder{}
	qb.SetAirplaneID(airplaneId)
	qb.SetCarrier(carrier)
	qb.SetFlightNumber(flightNumber)

	ent, err := db.Fetch(&entity.Airplane{}, qb)
	if err != nil {
		return nil, err
	}

	for _, e := range ent {
		s = append(s, e.(entity.Airplane))
	}

	return
}
