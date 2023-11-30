package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllPlayers(t *testing.T) {
	router :=  setupRouter()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/players", nil)
	router.ServeHTTP(resp, req)
	

	assert.Equal(t, 200, resp.Code)
	var response map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.Equal(t, "Success!", response["message"]) 
	assert.NotNil(t, response["data"]) 

	dataBytes, err := json.Marshal(response["data"])
    assert.NoError(t, err)

    var data []Player
    err = json.Unmarshal(dataBytes, &data)
    assert.NoError(t, err)

 
}

func TestGetAllQueues(t *testing.T) {
	router :=  setupRouter()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/queues", nil)
	router.ServeHTTP(resp, req)
	

	assert.Equal(t, 200, resp.Code)
	var response map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.Equal(t, "Success!", response["message"]) 
	assert.NotNil(t, response["data"]) 

	dataBytes, err := json.Marshal(response["data"])
    assert.NoError(t, err)

    var data []Queue
    err = json.Unmarshal(dataBytes, &data)
    assert.NoError(t, err)

 
}

func TestGetAllSessions(t *testing.T) {
	router :=  setupRouter()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/sessions", nil)
	router.ServeHTTP(resp, req)
	

	assert.Equal(t, 200, resp.Code)
	var response map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	

	assert.Equal(t, "Success!", response["message"]) 
	assert.NotNil(t, response["data"]) 

	dataBytes, err := json.Marshal(response["data"])
    assert.NoError(t, err)

    var data []Session
    err = json.Unmarshal(dataBytes, &data)
    assert.NoError(t, err)


}