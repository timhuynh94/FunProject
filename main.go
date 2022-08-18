package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/timhuynh94/TargetChallenge/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products/:id", getProductByID)
	r.PUT("/products/:id", updateProductByID)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getProductByID(c *gin.Context) {
	var id string
	id = c.Param("id")

	c.JSON(http.StatusOK, getProductIDFromRedsky(id))
}

func updateProductByID(c *gin.Context) {
	var id string
	id = c.Param("id")

	c.JSON(http.StatusOK, updateProductFromRedsky(id))
}

func getProductIDFromRedsky(id string) *models.RespBody {
	var response models.RespBody
	qId := fmt.Sprintf("https://redsky-uat.perf.target.com/redsky_aggregations/v1/redsky/case_study_v1?key=3yUxt7WltYG7MFKPp7uyELi1K40ad2ys&tcin=%s", id)
	resp, err := http.Get(qId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		log.Fatal("Can not unmarshal JSON")
	}
	return &response
}

func updateProductFromRedsky(id string) *models.RespBody {
	var body models.RespBody
	body.Data.Product.Tcin = id
	body.Data.Product.Item.CurrentPrice.Value = "12.9"
	var respBody models.RespBody
	qId := fmt.Sprintf("https://redsky-uat.perf.target.com/redsky_aggregations/v1/redsky/case_study_v1?key=3yUxt7WltYG7MFKPp7uyELi1K40ad2ys&tcin=%s", id)
	payload, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	resp, err := http.NewRequest(http.MethodPut, qId, bytes.NewBuffer(payload))
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(resBody, &respBody); err != nil { // Parse []byte to go struct pointer
		log.Fatal("Can not unmarshal JSON")
	}
	return &respBody
}
