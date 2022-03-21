package authbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type AuthTokenQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId int
}

func (qb *AuthTokenQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *AuthTokenQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.agentId < 1 {
		return "", args, fmt.Errorf("agentId is required")
	}

	partialQuery = fmt.Sprintf("%s and au.agent_id = ?", partialQuery)
	args = append(args, qb.agentId)

	// The token must be active
	partialQuery = fmt.Sprintf("%s and au.active = true", partialQuery)

	// The token must be valid
	partialQuery = fmt.Sprintf("%s and au.expires_at > now()", partialQuery)

	return partialQuery, args, nil
}

func (qb *AuthTokenQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	orderClause := "order by au.issued_at desc"
	limitClause := "limit 1"

	query := fmt.Sprintf(`
		select 
			au.token
		from agent_auth au
		%s
		%s
		%s
	`, whereClauses, orderClause, limitClause)

	return query, args, nil
}
