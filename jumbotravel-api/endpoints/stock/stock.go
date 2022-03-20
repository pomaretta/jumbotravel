package stock

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/application"
)

// Stock
//
// @Router /stock/latest [get]
// @Tags Stock
// @Summary Get latest stock.
//
// @Security Bearer
// @Produce json
//
// @Param stockid query int false "Stock ID"
// @Param airplaneid query int false "Airplane ID"
// @Param productid query int false "Product ID"
// @Param productcode query int false "Product Code"
//
// @Success 200 {object} response.JSONResult{result=[]dto.Stock} "Get latest stock."
// @Failure 400 {object} response.JSONError "Bad request"
// @Failure 500 {object} response.JSONError "Internal server error"
func Stock(application *application.Application) func(*gin.Context) {
	return func(c *gin.Context) {

		stockId := c.DefaultQuery("stockid", "-1")
		parsedStockId, err := strconv.Atoi(stockId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		airplaneId := c.DefaultQuery("airplaneid", "-1")
		parsedAirplaneId, err := strconv.Atoi(airplaneId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		productId := c.DefaultQuery("productid", "-1")
		parsedProductId, err := strconv.Atoi(productId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		productCode := c.DefaultQuery("productcode", "-1")
		parsedProductCode, err := strconv.Atoi(productCode)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		stock, err := application.GetStock(parsedStockId, parsedAirplaneId, parsedProductId, parsedProductCode)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": stock,
		})
	}
}
