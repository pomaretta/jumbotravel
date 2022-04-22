package authbuilders

import (
	"errors"
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type AgentAuthQueryBuilder struct {
	builders.MySQLQueryBuilder

	DNI string
}

func (qb *AgentAuthQueryBuilder) SetDNI(dni string) {
	qb.DNI = dni
}

func (qb *AgentAuthQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.DNI == "" {
		return "", nil, errors.New("DNI is required")
	}

	partialQuery = fmt.Sprintf("%s and ag.dni = ?", partialQuery)
	args = append(args, qb.DNI)

	// The agent must be active
	partialQuery = fmt.Sprintf("%s and ag.active = true", partialQuery)

	return partialQuery, args, nil
}

func (qb *AgentAuthQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf(`
		select 
			map.agentmapping_id,
			ag.dni,
			ag.password,
			ag.type
		from master_agentmapping map
		left join master_agents ag 
			on map.agent_id = ag.agent_id
		%s
	`, whereClauses)

	return query, args, nil
}
