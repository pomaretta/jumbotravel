package api

import (
	"github.com/pomaretta/jumbotravel/jumbotravel-api/endpoints/agent"
)

func (api *API) initAgent() {

	agentGroup := api.handler.Group("/agent/:id")
	{
		agentGroup.GET("/data", agent.AgentData(api.application))
	}

}
