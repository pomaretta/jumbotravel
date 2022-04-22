package agent

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain/entity"
)

// Notifications
//
// @Router /agent/:id/notifications [get]
// @Tags Agent
// @Summary Get agent notifications.
//
// @Security Bearer
// @Produce json
//
// @Param seen query string false "Seen" default(2)
// @Param active query string false "Active" default(1)
// @Param expired query string false "Expired" default(2)
// @Param popup query string false "Popup" default(0)
//
// @Success 200 {object} response.JSONResult{result=[]entity.Notification} "Get agent notifications."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Notifications(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		agentType := c.GetString("subtype")

		seen := c.DefaultQuery("seen", "2")
		if seen != "0" && seen != "1" && seen != "2" {
			c.JSON(400, gin.H{
				"error": "Invalid seen value",
			})
			return
		}

		active := c.DefaultQuery("active", "1")
		if active != "0" && active != "1" && active != "2" {
			c.JSON(400, gin.H{
				"error": "Invalid active value",
			})
			return
		}

		expired := c.DefaultQuery("expired", "2")
		if expired != "0" && expired != "1" && expired != "2" {
			c.JSON(400, gin.H{
				"error": "Invalid expired value",
			})
			return
		}

		popup := c.DefaultQuery("popup", "0")
		if popup != "0" && popup != "1" && popup != "2" {
			c.JSON(400, gin.H{
				"error": "Invalid popup value",
			})
			return
		}

		notifications, err := application.GetAgentNotifications(parsedAgentId, agentType, seen, active, expired, popup)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Sign the notifications
		for idx, notification := range notifications {
			signature, err := signNotification(&notification, parsedAgentId)
			notifications[idx].Signature = &signature
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
		}

		c.JSON(200, gin.H{
			"result": notifications,
		})
	}
}

func signNotification(notification *entity.Notification, agentId int) (string, error) {

	file, err := os.Open("rsa.private")
	if err != nil {
		return "", err
	}

	privatePEMData, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	block, rest := pem.Decode(privatePEMData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", errors.New("failed to decode PEM block containing private key " + block.Type)
	}

	if len(rest) != 0 {
		return "", errors.New("passed public key contains more than just the private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	rsaJWT := jwt.SigningMethodRS256

	now := time.Now().UTC()

	ttl := notification.ExpiresAt

	claims := jwt.MapClaims{
		"jti":        notification.NotificationId,
		"iat":        now.Unix(),
		"exp":        ttl.Unix(),
		"sub":        agentId,
		"iss":        "jumbotravel",
		"resourceId": notification.ResourceId,
		"scope":      notification.Scope,
		"seen":       notification.Seen,
		"active":     notification.Active,
	}

	token := jwt.NewWithClaims(rsaJWT, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ReadNotifications(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		pattern := regexp.MustCompile("[ \n\t\r]+")

		// Get the agent id
		agentId := c.Param("id")
		parsedAgentId, err := strconv.Atoi(agentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Get the notifications token from body
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if string(body) == "" {
			c.JSON(400, gin.H{
				"error": "No notifications token provided",
			})
			return
		}

		notifications := pattern.Split(string(body), -1)
		// Verify the notifications
		var parsedNotifications []entity.Notification
		for _, notification := range notifications {
			parsedNotification, err := verifyNotification(notification, parsedAgentId)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			parsedNotifications = append(parsedNotifications, parsedNotification)
		}

		// Mark the notifications as read
		var notificationIds []int
		for _, notification := range parsedNotifications {
			notificationIds = append(notificationIds, *notification.NotificationId)
		}

		rowsAffected, err := application.PostUpdateAgentNotifications(notificationIds)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": rowsAffected,
		})
	}
}

func verifyNotification(token string, agentId int) (entity.Notification, error) {

	file, err := os.Open("rsa.public")
	if err != nil {
		return entity.Notification{}, err
	}

	pubPEMData, err := io.ReadAll(file)
	if err != nil {
		return entity.Notification{}, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(pubPEMData)
	if err != nil {
		return entity.Notification{}, err
	}
	rsaJWT := jwt.SigningMethodRS256

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return entity.Notification{}, errors.New("invalid notification token")
	}

	err = rsaJWT.Verify(strings.Join(parts[0:2], "."), parts[2], key)
	if err != nil {
		return entity.Notification{}, err
	}
	claimsBytes, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return entity.Notification{}, err
	}

	dec := json.NewDecoder(bytes.NewBuffer(claimsBytes))

	c := jwt.MapClaims{
		"jti":        0,
		"iat":        0,
		"exp":        0,
		"sub":        0,
		"iss":        "",
		"resourceId": 0,
		"scope":      "",
		"seen":       false,
		"active":     false,
	}
	err = dec.Decode(&c)
	if err != nil {
		return entity.Notification{}, err
	}

	// Cast float64 to int
	notificationAgentId := int(c["sub"].(float64))

	if notificationAgentId != agentId {
		return entity.Notification{}, errors.New("invalid notification token")
	}

	parsedId := int(c["jti"].(float64))
	parsedResourceId := int(c["resourceId"].(float64))
	parsedScope := c["scope"].(string)
	parsedSeen := c["seen"].(bool)
	parsedActive := c["active"].(bool)

	return entity.Notification{
		NotificationId: &parsedId,
		ResourceId:     &parsedResourceId,
		Scope:          &parsedScope,
		Seen:           &parsedSeen,
		Active:         &parsedActive,
	}, nil
}
