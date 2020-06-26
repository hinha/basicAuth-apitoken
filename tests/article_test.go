package tests

import (
	setRouter "basicApi/router"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetEmptyArray(t *testing.T) {
	var emptyArray []interface{}
	body := gin.H{
		"data": emptyArray,
	}

	//	grab router
	router := setRouter.SetupRouter()

	w := performRequest(router, "GET", "/articles")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte("null"), &response)

	_, exists := response["data"]

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.False(t, exists)
	assert.Equal(t, body["data"], emptyArray)
}
