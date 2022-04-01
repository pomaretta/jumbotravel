package authbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type AuthTokenQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId int
	jti     string
	active  string
	expired string
	single  bool
}

func (qb *AuthTokenQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *AuthTokenQueryBuilder) SetJTI(jti string) {
	qb.jti = jti
}

func (qb *AuthTokenQueryBuilder) SetActive(active string) {
	qb.active = active
}

func (qb *AuthTokenQueryBuilder) SetExpired(expired string) {
	qb.expired = expired
}

func (qb *AuthTokenQueryBuilder) SetSingle(single bool) {
	qb.single = single
}

func (qb *AuthTokenQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.agentId > -1 {
		partialQuery = fmt.Sprintf("%s and au.agent_id = ?", partialQuery)
		args = append(args, qb.agentId)
	}

	if qb.jti != "" {
		partialQuery = fmt.Sprintf("%s and au.id = ?", partialQuery)
		args = append(args, qb.jti)
	}

	if qb.active != "" && qb.active == "1" {
		partialQuery = fmt.Sprintf("%s and au.active = true", partialQuery)
	}

	if qb.expired != "" && qb.expired == "1" {
		partialQuery = fmt.Sprintf("%s and au.expires_at > now()", partialQuery)
	}

	return partialQuery, args, nil
}

func (qb *AuthTokenQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	orderClause := "order by au.issued_at desc"
	limitClause := ""
	if qb.single {
		limitClause = "limit 1"
	}

	query := fmt.Sprintf(`
		select 
			au.id,
			au.subjet,
			au.token,
			au.agent_id,
			au.issued_at,
			au.expires_at,
			au.active,
			au.created_at
		from agent_auth au
		%s
		%s
		%s
	`, whereClauses, orderClause, limitClause)

	return query, args, nil
}
