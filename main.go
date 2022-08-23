package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/timhuynh94/TargetChallenge/models"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var rdbClient = NewRdbClient()

func main() {

	r := gin.Default()

	r.GET("/products/:id", getProductByID)
	r.PUT("/products/:id", updateProductByID)
	r.GET("/health", getHealth)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Healthy"})
}

// getProductByID handles GET calls and returns product details including pricing info in JSON format
// returns 404 if product not found
func getProductByID(c *gin.Context) {
	var id string
	id = c.Param("id")
	prodDetails := getProductIDFromRedsky(id)
	p, errNil := rdbClient.getProductFromRedis(id)
	if errNil != nil && prodDetails.Data.Product.Tcin != "" {
		// if no pricing available, then update price to be NA
		prodDetails.Data.Product.Item.CurrentPrice.Value = "NA"
		prodDetails.Data.Product.Item.CurrentPrice.CurrencyCode = "NA"
		c.JSON(http.StatusOK, prodDetails.Data)
	}
	if errNil == nil && prodDetails.Data.Product.Tcin != "" {
		// if pricing is available, update products details to have the latest pricing
		prodDetails.Data.Product.Item.CurrentPrice = p.Data.Product.Item.CurrentPrice
		c.JSON(http.StatusOK, prodDetails.Data)
	}
	if prodDetails.Data.Product.Tcin == "" {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Item %s not found", id))
	}
}

// updateProductByID handles check if product is valid from redsky before updating pricing info in Redis
func updateProductByID(c *gin.Context) {
	var id string
	var pricing models.RespBody
	id = c.Param("id")
	prodDetails := getProductIDFromRedsky(id)
	if prodDetails.Data.Product.Tcin == "" {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Item %s not found", id))
		return
	}
	// binding input json from PUT request call
	if err := c.ShouldBindBodyWith(&pricing.Data, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Setting pricing information into redis
	err := rdbClient.setProductToRedis(id, pricing)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Product %s has been updated", id))
}

// getProductIDFromRedsky calls Redsky API and returns response in struct format
func getProductIDFromRedsky(id string) *models.RespBody {
	var response models.RespBody
	qId := fmt.Sprintf("https://redsky-uat.perf.target.com/redsky_aggregations/v1/redsky/case_study_v1?key=3yUxt7WltYG7MFKPp7uyELi1K40ad2ys&tcin=%s", id)
	resp, err := http.Get(qId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal("Can not read from response")
	}
	if jErr := json.Unmarshal(body, &response); jErr != nil { // Parse []byte to go struct pointer
		log.Fatal("Can not unmarshal JSON: ", jErr)
	}
	return &response
}
