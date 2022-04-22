package master

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Notifications
//
// @Router /master/notifications [get]
// @Tags Master
// @Summary Get master notifications.
//
// @Security Bearer
// @Produce json
//
// @Param notificationid query int false "Notification ID"
// @Param resourceid query int false "Resource ID"
// @Param resourceuuid query string false "Resource UUID"
// @Param notificationtype query string false "Notification type"
// @Param scope query string false "Scope"
// @Param seen query string false "Seen"
// @Param active query string false "Active"

//
// @Success 200 {object} response.JSONResult{result=[]entity.Notification} "Get master notifications."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Notifications(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		pattern := regexp.MustCompile("[ ,;\n\t\r]+")

		notificationId := c.Query("notificationid")
		parsedNotificationsIds := make([]int, 0)
		if notificationId != "" {
			notificationIds := pattern.Split(notificationId, -1)
			for _, notificationId := range notificationIds {
				parsedNotificationId, err := strconv.Atoi(notificationId)
				if err != nil {
					c.JSON(400, gin.H{
						"error": "notification id must be an integer",
					})
					return
				}
				parsedNotificationsIds = append(parsedNotificationsIds, parsedNotificationId)
			}
		}

		resourceId := c.Query("resourceid")
		parsedResourceIds := make([]int, 0)
		if resourceId != "" {
			resourceIds := pattern.Split(resourceId, -1)
			for _, resourceId := range resourceIds {
				parsedResourceId, err := strconv.Atoi(resourceId)
				if err != nil {
					c.JSON(400, gin.H{
						"error": "resourceid id must be an integer",
					})
					return
				}
				parsedResourceIds = append(parsedResourceIds, parsedResourceId)
			}
		}

		resourceUuid := c.Query("resourceuuid")
		parsedResourceUuids := make([]string, 0)
		if resourceUuid != "" {
			resourceUuids := pattern.Split(resourceUuid, -1)
			parsedResourceUuids = append(parsedResourceUuids, resourceUuids...)
		}

		notificationType := c.Query("notificationtype")
		parsedNotificationTypes := make([]string, 0)
		if notificationType != "" {
			notificationTypes := pattern.Split(notificationType, -1)
			parsedNotificationTypes = append(parsedNotificationTypes, notificationTypes...)
		}

		scope := c.Query("scope")
		parsedScopes := make([]string, 0)
		if scope != "" {
			scopes := pattern.Split(scope, -1)
			parsedScopes = append(parsedScopes, scopes...)
		}

		seen := strings.ToLower(c.DefaultQuery("seen", "0"))
		if seen != "0" && seen != "1" && seen != "2" {
			c.JSON(400, gin.H{
				"error": "seen must be 0, 1 or 2",
			})
			return
		}

		active := strings.ToLower(c.DefaultQuery("active", "0"))
		if active != "0" && active != "1" && active != "2" {
			c.JSON(400, gin.H{
				"error": "active must be 0, 1 or 2",
			})
			return
		}

		expired := strings.ToLower(c.DefaultQuery("expired", "0"))
		if expired != "0" && expired != "1" && expired != "2" {
			c.JSON(400, gin.H{
				"error": "expired must be 0, 1 or 2",
			})
			return
		}

		popup := strings.ToLower(c.DefaultQuery("popup", "0"))
		if popup != "0" && popup != "1" && popup != "2" {
			c.JSON(400, gin.H{
				"error": "popup must be 0, 1 or 2",
			})
			return
		}

		notifications, err := application.GetNotifications(
			parsedNotificationsIds,
			parsedResourceIds,
			parsedResourceUuids,
			parsedNotificationTypes,
			parsedScopes,
			seen,
			active,
			expired,
			popup,
		)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": notifications,
		})
	}
}
