package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompany(t *testing.T) {
	router := gin.Default()
	// Mock your service and handlers here

	//Add also mock DB for tests - Unittest db

	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"name": "My Company", "amount_of_employees": 240, "registered": true, "type": "Corporations"}`)
	req, _ := http.NewRequest("POST", "/api/v1/companies", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
