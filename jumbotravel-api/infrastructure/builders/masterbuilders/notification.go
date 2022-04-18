package masterbuilders

import (
	"fmt"
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type NotificationQueryBuilder struct {
	builders.MySQLQueryBuilder

	notificationId   []int
	resourceId       []int
	resourceUuid     []string
	notificationType []string
	scope            []string
	seen             string
	active           string
	expired          string
	popup            string
}

func (qb *NotificationQueryBuilder) SetNotificationId(notificationId []int) {
	qb.notificationId = notificationId
}

func (qb *NotificationQueryBuilder) SetResourceId(resourceId []int) {
	qb.resourceId = resourceId
}

func (qb *NotificationQueryBuilder) SetResourceUuid(resourceUuid []string) {
	qb.resourceUuid = resourceUuid
}

func (qb *NotificationQueryBuilder) SetNotificationType(notificationType []string) {
	qb.notificationType = notificationType
}

func (qb *NotificationQueryBuilder) SetScope(scope []string) {
	qb.scope = scope
}

func (qb *NotificationQueryBuilder) SetSeen(seen string) {
	qb.seen = seen
}

func (qb *NotificationQueryBuilder) SetActive(active string) {
	qb.active = active
}

func (qb *NotificationQueryBuilder) SetExpired(expired string) {
	qb.expired = expired
}

func (qb *NotificationQueryBuilder) SetPopup(popup string) {
	qb.popup = popup
}

func (qb *NotificationQueryBuilder) buildWhereClause() (string, []interface{}, error) {

	partialQuery := "where 1=1"
	args := []interface{}{}

	if len(qb.notificationId) > 0 {
		partialQuery = fmt.Sprintf("%s and notification_id in (", partialQuery)
		for idx, notificationId := range qb.notificationId {
			if idx == 0 {
				partialQuery = fmt.Sprintf("%s?", partialQuery)
				args = append(args, notificationId)
				continue
			}
			partialQuery = fmt.Sprintf("%s,?", partialQuery)
			args = append(args, notificationId)
		}
		partialQuery = fmt.Sprintf("%s)", partialQuery)
	}

	if len(qb.resourceId) > 0 {
		partialQuery = fmt.Sprintf("%s and resource_id in (", partialQuery)
		for idx, resourceId := range qb.resourceId {
			if idx == 0 {
				partialQuery = fmt.Sprintf("%s?", partialQuery)
				args = append(args, resourceId)
				continue
			}
			partialQuery = fmt.Sprintf("%s,?", partialQuery)
			args = append(args, resourceId)
		}
		partialQuery = fmt.Sprintf("%s)", partialQuery)
	}

	if len(qb.resourceUuid) > 0 {
		partialQuery = fmt.Sprintf("%s and resource_uuid in (", partialQuery)
		for idx, resourceUuid := range qb.resourceUuid {
			if idx == 0 {
				partialQuery = fmt.Sprintf("%s?", partialQuery)
				args = append(args, resourceUuid)
				continue
			}
			partialQuery = fmt.Sprintf("%s,?", partialQuery)
			args = append(args, resourceUuid)
		}
		partialQuery = fmt.Sprintf("%s)", partialQuery)
	}

	if len(qb.notificationType) > 0 {
		partialQuery = fmt.Sprintf("%s and type in (", partialQuery)
		for idx, notifType := range qb.notificationType {
			if idx == 0 {
				partialQuery = fmt.Sprintf("%s?", partialQuery)
				args = append(args, notifType)
				continue
			}
			partialQuery = fmt.Sprintf("%s,?", partialQuery)
			args = append(args, notifType)
		}
		partialQuery = fmt.Sprintf("%s)", partialQuery)
	}

	if len(qb.scope) > 0 {
		partialQuery = fmt.Sprintf("%s and scope in (", partialQuery)
		for idx, notifScope := range qb.scope {
			if idx == 0 {
				partialQuery = fmt.Sprintf("%s?", partialQuery)
				args = append(args, notifScope)
				continue
			}
			partialQuery = fmt.Sprintf("%s,?", partialQuery)
			args = append(args, notifScope)
		}
		partialQuery = fmt.Sprintf("%s)", partialQuery)
	}

	if qb.seen != "0" {
		parsedSeen := false
		if qb.seen == "1" {
			parsedSeen = true
		}
		partialQuery = fmt.Sprintf("%s and seen = ?", partialQuery)
		args = append(args, parsedSeen)
	}

	if qb.active != "0" {
		parsedActive := false
		if qb.active == "1" {
			parsedActive = true
		}
		partialQuery = fmt.Sprintf("%s and active = ?", partialQuery)
		args = append(args, parsedActive)
	}

	if qb.expired != "0" && qb.expired == "1" {
		partialQuery = fmt.Sprintf("%s and expires_at < ?", partialQuery)
		args = append(args, fmt.Sprintf("%v", time.Now().UTC()))
	}

	if qb.popup != "0" {
		parsedPopup := false
		if qb.popup == "1" {
			parsedPopup = true
		}
		partialQuery = fmt.Sprintf("%s and popup = ?", partialQuery)
		args = append(args, parsedPopup)
	}

	return partialQuery, args, nil
}

func (qb *NotificationQueryBuilder) BuildQuery() (string, []interface{}, error) {

	whereClauses, args, err := qb.buildWhereClause()
	if err != nil {
		return "", nil, err
	}
	orderBy := "order by created_at desc"

	query := fmt.Sprintf(`
		select
			notification_id,
			scope,
			resource_id,
			resource_uuid,
			title,
			message,
			link,
			extra,
			type,
			popup,
			expires_at,
			created_at,
			seen,
			active
		FROM notifications
		%s
		%s
	`, whereClauses, orderBy)

	return query, args, nil
}
