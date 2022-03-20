package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/docs/response"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/lib/rsajwt"
)

// Login
//
// @Router /auth/login [post]
// @Tags Auth
// @Summary Login into the system
//
// @Security Bearer
// @Produce json
//
// @Success 200 {object} response.JSONResult{result=string} "Get token."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Login(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		var body map[string]interface{}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{
				"error": "A JSON body is required.",
			})
			return
		}

		// Get the DNI and password from the body
		agentDni, ok := body["dni"].(string)
		if !ok {
			c.JSON(400, gin.H{
				"error": "dni is required",
			})
			return
		}

		agentPassword, ok := body["password"].(string)
		if !ok {
			c.JSON(400, gin.H{
				"error": "password is required",
			})
			return
		}

		agent, err := application.GetAgentAuth(agentDni)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// TODO: Check if the password provided is the same as agent's password
		if *agent.Password != agentPassword {
			response.Unauthorized(c)
			return
		}

		// TODO: Check if the agent has an active token
		currentToken, err := application.GetAuthToken(*agent.AgentId)
		if err == nil && currentToken != "" {
			c.JSON(200, gin.H{
				"result": currentToken,
			})
			return
		}

		// TODO: Generate token with agent info and return it
		signer, err := rsajwt.NewSignerFromPrivateKeyFile("rsa.private")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		token, err := signer.Sign(
			rsajwt.AllowPolicy{
				Resources: []string{
					fmt.Sprintf("/agent/%d/*", *agent.AgentId),
				},
			},
			*agent.DNI,
			"user",
		)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = application.PutToken(
			token.Claims.Id,
			token.Claims.Subject,
			token.Token,
			*agent.AgentId,
			time.Unix(token.Claims.IssuedAt, 0),
			time.Unix(token.Claims.ExpiresAt, 0),
		)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": token.Token,
		})
	}
}
