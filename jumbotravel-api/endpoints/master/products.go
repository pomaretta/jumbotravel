package master

import (
	"regexp"
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

		// productId := c.DefaultQuery("productid", "-1")
		// parsedProductId, err := strconv.Atoi(productId)
		// if err != nil {
		// 	c.JSON(400, gin.H{
		// 		"error": "agentid must be an integer",
		// 	})
		// 	return
		// }

		// productCode := c.DefaultQuery("productcode", "-1")
		// parsedProductCode, err := strconv.Atoi(productCode)
		// if err != nil {
		// 	c.JSON(400, gin.H{
		// 		"error": "agentid must be an integer",
		// 	})
		// 	return
		// }
		pattern := regexp.MustCompile("[ ,;\n\t\r]+")

		productId := c.Query("productid")
		parsedProductsId := make([]int, 0)
		if productId != "" {
			productsId := pattern.Split(productId, -1)
			for _, productId := range productsId {
				parsedProductId, err := strconv.Atoi(productId)
				if err != nil {
					c.JSON(400, gin.H{
						"error": "productid must be an integer",
					})
					return
				}
				parsedProductsId = append(parsedProductsId, parsedProductId)
			}
		}

		productCode := c.Query("productcode")
		parsedProductsCodes := make([]int, 0)
		if productCode != "" {
			productsCode := pattern.Split(productCode, -1)
			for _, productCode := range productsCode {
				parsedProductCode, err := strconv.Atoi(productCode)
				if err != nil {
					c.JSON(400, gin.H{
						"error": "productid must be an integer",
					})
					return
				}
				parsedProductsCodes = append(parsedProductsCodes, parsedProductCode)
			}
		}

		products, err := application.GetMasterProducts(parsedProductsId, parsedProductsCodes)
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
