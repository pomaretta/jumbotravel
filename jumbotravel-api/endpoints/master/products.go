package master

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Products
//
// @Router /master/products [get]
// @Tags Master
// @Summary Get master products.
//
// @Security Bearer
// @Produce json
//
// @Param productid query int false "Product ID"
// @Param productcode query int false "Product code"
//
// @Success 200 {object} response.JSONResult{result=[]entity.Product} "Get master products"
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Products(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		productId := c.DefaultQuery("productid", "-1")
		parsedProductId, err := strconv.Atoi(productId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "agentid must be an integer",
			})
			return
		}

		productCode := c.DefaultQuery("productcode", "-1")
		parsedProductCode, err := strconv.Atoi(productCode)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "agentid must be an integer",
			})
			return
		}

		products, err := application.GetMasterProducts(parsedProductId, parsedProductCode)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": products,
		})
	}
}
