package api

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/endpoints/agent"
)

func (api *API) initAgent() {

	agentGroup := api.handler.Group("/agent/:id")
	{
		agentGroup.GET("/data", agent.Data(api.application))
		agentGroup.GET("/validate", agent.Validate(api.application))

		// Notifications
		notificationsGroup := agentGroup.Group("/notifications")
		{
			notificationsGroup.GET("", agent.Notifications(api.application))
			notificationsGroup.POST("", agent.ReadNotifications(api.application))
		}

	}

}
