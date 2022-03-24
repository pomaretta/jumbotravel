package auth

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
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
// @Success 200 {object} response.JWTToken "Get token data."
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

		// MD5 the password
		h := md5.New()
		h.Write([]byte(agentPassword))
		agentPassword = fmt.Sprintf("%x", h.Sum(nil))

		agent, err := application.GetAgentAuth(agentDni)
		if err != nil {
			c.JSON(401, gin.H{
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

			// Split the token with "."
			tokenParts := bytes.Split([]byte(currentToken), []byte("."))

			// Decode the token
			claimsBytes, err := jwt.DecodeSegment(string(tokenParts[1]))
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			dec := json.NewDecoder(bytes.NewBuffer(claimsBytes))

			claims := rsajwt.Claims{}
			err = dec.Decode(&claims)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, gin.H{
				"jti":   claims.Id,
				"token": currentToken,
				"exp":   time.Unix(claims.ExpiresAt, 0),
				"iat":   time.Unix(claims.IssuedAt, 0),
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
			"jti":   token.Claims.Id,
			"token": token.Token,
			"exp":   time.Unix(token.Claims.ExpiresAt, 0),
			"iat":   time.Unix(token.Claims.IssuedAt, 0),
		})
	}
}
