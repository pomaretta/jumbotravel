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
