package mysql

import (
	"errors"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
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
		return dto.AgentAuth{}, errors.New("Agent not found")
	}

	return ent[0].(dto.AgentAuth), nil
}

func (db *MySQL) FetchAuthToken(agentId int) (string, error) {

	qb := &authbuilders.AuthTokenQueryBuilder{}
	qb.SetAgentId(agentId)

	query, args, err := qb.BuildQuery()
	if err != nil {
		return "", err
	}

	rows, err := db.con.Query(query, args...)
	if err != nil {
		return "", err
	}

	var token string
	for rows.Next() {
		err = rows.Scan(&token)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}
