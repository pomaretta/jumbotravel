package mysql

import (
	"errors"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/authbuilders"
)

func (db *MySQL) PutToken(tokenId, subject, token string, agentId int, issuedAt, expiresAt time.Time) error {

	qb := &authbuilders.PutTokenQueryBuilder{}
	qb.SetTokenId(tokenId)
	qb.SetSubject(subject)
	qb.SetToken(token)
	qb.SetAgentId(agentId)
	qb.SetIssuedAt(issuedAt)
	qb.SetExpiresAt(expiresAt)

	_, err := db.Put(qb)
	if err != nil {
		return err
	}

	return nil
}

func (db *MySQL) FetchAgentAuth(dni string) (dto.AgentAuth, error) {

	qb := &authbuilders.AgentAuthQueryBuilder{}
	qb.SetDNI(dni)

	ent, err := db.Fetch(&dto.AgentAuth{}, qb)
	if err != nil {
		return dto.AgentAuth{}, err
	}

	if len(ent) == 0 {
		return dto.AgentAuth{}, errors.New("identifier not found")
	}

	return ent[0].(dto.AgentAuth), nil
}

func (db *MySQL) FetchAuthToken(agentId int, jti string, active string, expired string, single bool) (s []entity.Token, err error) {

	qb := &authbuilders.AuthTokenQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetJTI(jti)
	qb.SetActive(active)
	qb.SetExpired(expired)
	qb.SetSingle(single)

	ent, err := db.Fetch(&entity.Token{}, qb)
	if err != nil {
		return
	}

	for _, e := range ent {
		s = append(s, e.(entity.Token))
	}

	return
}
