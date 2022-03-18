package masterbuilders

import (
	"fmt"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type MasterAgentQueryBuilder struct {
	builders.MySQLQueryBuilder

	AgentID int
	DNI     string
	Type    string
	Active  bool
	Email   string
}

func (qb *MasterAgentQueryBuilder) SetAgentID(agentID int) {
	qb.AgentID = agentID
}

func (qb *MasterAgentQueryBuilder) SetDNI(dni string) {
	qb.DNI = dni
}

func (qb *MasterAgentQueryBuilder) SetType(agentType string) {
	qb.Type = agentType
}

func (qb *MasterAgentQueryBuilder) SetActive(active bool) {
	qb.Active = active
}

func (qb *MasterAgentQueryBuilder) SetEmail(email string) {
	qb.Email = email
}

func (qb *MasterAgentQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if qb.AgentID > -1 {
		partialQuery = fmt.Sprintf("%s and ag.agent_id = ?", partialQuery)
		args = append(args, qb.AgentID)
	}

	if qb.DNI != "" {
		partialQuery = fmt.Sprintf("%s and ag.dni = ?", partialQuery)
		args = append(args, qb.DNI)
	}

	if qb.Type != "" {
		partialQuery = fmt.Sprintf("%s and ag.type = ?", partialQuery)
		args = append(args, qb.Type)
	}

	if qb.Active {
		partialQuery = fmt.Sprintf("%s and ag.active = true", partialQuery)
	} else {
		partialQuery = fmt.Sprintf("%s and ag.active = false", partialQuery)
	}

	if qb.Email != "" {
		partialQuery = fmt.Sprintf("%s and ag.email = ?", partialQuery)
		args = append(args, qb.Email)
	}

	return partialQuery, args, nil
}

func (qb *MasterAgentQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}

	query := fmt.Sprintf(`
		select 
			ag.agent_id,
			ag.dni,
			ag.name,
			ag.surname,
			ag.email,
			ag.password,
			ag.type,
			ag.airport_id,
			ag.created_at,
			ag.active
		from master_agentmapping map
		left join master_agents ag 
			on map.agent_id = ag.agent_id
		%s
	`, whereClauses)

	return query, args, nil
}
