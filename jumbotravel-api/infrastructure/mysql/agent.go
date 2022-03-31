package mysql

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/agentbuilders"
)

func (db *MySQL) FetchAgentNotifications(agentId int, seen, active, expired, popup string) (s []entity.Notification, err error) {

	qb := &agentbuilders.NotificationsQueryBuilder{}
	qb.SetAgentId(agentId)
	qb.SetSeen(seen)
	qb.SetActive(active)
	qb.SetExpired(expired)
	qb.SetPopup(popup)

	ent, err := db.Fetch(&entity.Notification{}, qb)
	if err != nil {
		return
	}

	for _, e := range ent {
		s = append(s, e.(entity.Notification))
	}

	return
}
