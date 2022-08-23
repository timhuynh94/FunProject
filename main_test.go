package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHealthHandler(t *testing.T) {
	mockResponse := `{"message":"Healthy","status":"success"}`
	r := SetUpRouter()
	r.GET("/health", getHealth)
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGETHandlerNotFound(t *testing.T) {
	mockResponse := `"Item 1234 not found"`
	r := SetUpRouter()
	r.GET("/products/:id", getProductByID)
	req, _ := http.NewRequest("GET", "/products/1234", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusNotFound, w.Code)
}
