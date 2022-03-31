package agentbuilders

import (
	"fmt"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type NotificationsQueryBuilder struct {
	builders.MySQLQueryBuilder

	agentId int
	seen    string
	active  string
	expired string
	popup   string
}

func (qb *NotificationsQueryBuilder) SetAgentId(agentId int) {
	qb.agentId = agentId
}

func (qb *NotificationsQueryBuilder) SetSeen(seen string) {
	qb.seen = seen
}

func (qb *NotificationsQueryBuilder) SetActive(active string) {
	qb.active = active
}

func (qb *NotificationsQueryBuilder) SetExpired(expired string) {
	qb.expired = expired
}

func (qb *NotificationsQueryBuilder) SetPopup(popup string) {
	qb.popup = popup
}

func (qb *NotificationsQueryBuilder) buildGlobalWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Set GLOBAL Scope
	partialQuery = fmt.Sprintf("%s and n.scope = ?", partialQuery)
	args = append(args, "GLOBAL")

	if qb.active != "0" && qb.active == "1" {
		partialQuery = fmt.Sprintf("%s and n.active = ?", partialQuery)
		args = append(args, true)
	}

	if qb.expired != "0" && qb.expired == "1" {
		partialQuery = fmt.Sprintf("%s and n.expires_at < ?", partialQuery)
		args = append(args, fmt.Sprintf("%v", time.Now().UTC()))
	}

	if qb.popup != "0" && qb.popup == "1" {
		partialQuery = fmt.Sprintf("%s and n.popup = ?", partialQuery)
		args = append(args, true)
	}

	return partialQuery, args, nil
}

func (qb *NotificationsQueryBuilder) buildAgentWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Set GLOBAL Scope
	partialQuery = fmt.Sprintf("%s and n.scope = ?", partialQuery)
	args = append(args, "AGENT")

	if qb.agentId <= 0 {
		return "", nil, fmt.Errorf("agentid is required")
	}

	partialQuery = fmt.Sprintf("%s and n.resource_id = ?", partialQuery)
	args = append(args, qb.agentId)

	if qb.active != "0" && qb.active == "1" {
		partialQuery = fmt.Sprintf("%s and n.active = ?", partialQuery)
		args = append(args, true)
	}

	if qb.seen != "0" && qb.seen == "1" {
		partialQuery = fmt.Sprintf("%s and n.seen = ?", partialQuery)
		args = append(args, true)
	}

	if qb.expired != "0" && qb.expired == "1" {
		partialQuery = fmt.Sprintf("%s and n.expires_at < ?", partialQuery)
		args = append(args, fmt.Sprintf("%v", time.Now().UTC()))
	}

	if qb.popup != "0" && qb.popup == "1" {
		partialQuery = fmt.Sprintf("%s and n.popup = ?", partialQuery)
		args = append(args, true)
	}

	return partialQuery, args, nil
}

func (qb *NotificationsQueryBuilder) buildFlightWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	// Set GLOBAL Scope
	partialQuery = fmt.Sprintf("%s and n.scope = ?", partialQuery)
	args = append(args, "FLIGHT")

	if qb.agentId <= 0 {
		return "", nil, fmt.Errorf("agentid is required")
	}
	partialQuery = fmt.Sprintf("%s and fa.agentmapping_id = ?", partialQuery)
	args = append(args, qb.agentId)

	if qb.active != "0" && qb.active == "1" {
		partialQuery = fmt.Sprintf("%s and n.active = ?", partialQuery)
		args = append(args, true)
	}

	if qb.seen != "0" && qb.seen == "1" {
		partialQuery = fmt.Sprintf("%s and n.seen = ?", partialQuery)
		args = append(args, true)
	}

	if qb.expired != "0" && qb.expired == "1" {
		partialQuery = fmt.Sprintf("%s and n.expires_at < ?", partialQuery)
		args = append(args, fmt.Sprintf("%v", time.Now().UTC()))
	}

	if qb.popup != "0" && qb.popup == "1" {
		partialQuery = fmt.Sprintf("%s and n.popup = ?", partialQuery)
		args = append(args, true)
	}

	return partialQuery, args, nil
}

func (qb *NotificationsQueryBuilder) BuildQuery() (string, []interface{}, error) {

	var args []interface{}
	orderClause := "order by created_at asc"

	globalClauses, globalArgs, err := qb.buildGlobalWhereClause()
	if err != nil {
		return "", nil, err
	}
	args = append(args, globalArgs...)

	agentClauses, agentArgs, err := qb.buildAgentWhereClause()
	if err != nil {
		return "", nil, err
	}
	args = append(args, agentArgs...)

	flightClauses, flightArgs, err := qb.buildFlightWhereClause()
	if err != nil {
		return "", nil, err
	}
	args = append(args, flightArgs...)

	query := fmt.Sprintf(`
	with
		global_notifications as (
			SELECT * FROM notifications n 
			%s
		),
		agent_notifications as (
			SELECT * FROM notifications n 
			%s
		),
		flight_notifications as (
			SELECT n.* FROM notifications n
			LEFT JOIN flight_agents fa
				ON fa.flight_id = n.resource_id
			%s
		)
	SELECT * FROM global_notifications
	UNION ALL
	SELECT * FROM agent_notifications
	UNION ALL
	SELECT * FROM flight_notifications
	%s
	`, globalClauses, agentClauses, flightClauses, orderClause)

	return query, args, nil
}
