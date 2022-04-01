package application

import (
	"time"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/dto"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
)

func (app *Application) PutToken(tokenId, subject, token string, agentId int, issuedAt, expiresAt time.Time) error {
	return app.MySQLFetcher.PutToken(tokenId, subject, token, agentId, issuedAt, expiresAt)
}

func (app *Application) GetAgentAuth(dni string) (dto.AgentAuth, error) {
	return app.MySQLFetcher.FetchAgentAuth(dni)
}

func (app *Application) GetAuthToken(agentId int, jti, active, expired string, single bool) ([]entity.Token, error) {
	return app.MySQLFetcher.FetchAuthToken(agentId, jti, active, expired, single)
}
