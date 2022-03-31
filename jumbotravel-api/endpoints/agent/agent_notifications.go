package agent

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
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

		notifications, err := application.GetAgentNotifications(parsedAgentId, seen, active, expired, popup)
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
