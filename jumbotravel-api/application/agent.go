package application

import "github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"

func (app *Application) GetAgentNotifications(agentId int, seen, active, expired, popup string) (s []entity.Notification, err error) {
	return app.FetchAgentNotifications(agentId, seen, active, expired, popup)
}
