package authbuilders

import (
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type PutTokenQueryBuilder struct {
	builders.MySQLQueryBuilder

	tokenId   string
	subjet    string
	token     string
	agentId   int
	issuedAt  time.Time
	expiresAt time.Time
}

func (qb *PutTokenQueryBuilder) SetTokenId(tokenId string) {
	qb.tokenId = tokenId
}

func (qb *PutTokenQueryBuilder) SetSubject(subject string) {
	qb.subjet = subject
}

func (qb *PutTokenQueryBuilder) SetToken(token string) {
	qb.token = token
}

func (qb *PutTokenQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *PutTokenQueryBuilder) SetIssuedAt(issuedAt time.Time) {
	qb.issuedAt = issuedAt
}

func (qb *PutTokenQueryBuilder) SetExpiresAt(expiresAt time.Time) {
	qb.expiresAt = expiresAt
}

func (qb *PutTokenQueryBuilder) createArguments() ([]interface{}, error) {

	args := []interface{}{}

	if qb.tokenId != "" {
		args = append(args, qb.tokenId)
	}

	if qb.subjet != "" {
		args = append(args, qb.subjet)
	}

	if qb.token != "" {
		args = append(args, qb.token)
	}

	if qb.agentId > -1 {
		args = append(args, qb.agentId)
	}

	if !qb.issuedAt.IsZero() {
		args = append(args, qb.issuedAt)
	}

	if !qb.expiresAt.IsZero() {
		args = append(args, qb.expiresAt)
	}

	return args, nil
}

func (qb *PutTokenQueryBuilder) BuildQuery() (string, []interface{}, error) {

	args, err := qb.createArguments()
	if err != nil {
		return "", nil, err
	}

	partialQuery := `
		INSERT INTO agent_auth (
			id,
			subjet,
			token,
			agent_id,
			issued_at,
			expires_at
		) VALUES
		(
			?,
			?,
			?,
			?,
			?,
			?
		);
	`

	return partialQuery, args, nil
}
