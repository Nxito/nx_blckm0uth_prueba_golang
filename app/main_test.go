package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"golang-gin-crud-api/model"
	"golang-gin-crud-api/router"  
	"github.com/stretchr/testify/assert"
)

func TestGetAllPlayers(t *testing.T) {
	router :=  router.SetupRouter()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/players", nil)
	router.ServeHTTP(resp, req)
	

	assert.True(t, resp.Code == 200 || resp.Code == 204, "Esperado 200 o 204. Se obtuvo %d", resp.Code)

	var response map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.Equal(t, "Success!", response["message"]) 
	assert.NotNil(t, response["data"]) 

	dataBytes, err := json.Marshal(response["data"])
    assert.NoError(t, err)

    var data []model.Player
    err = json.Unmarshal(dataBytes, &data)
    assert.NoError(t, err)

 
}

func TestGetAllQueues(t *testing.T) {
	router :=  router.SetupRouter()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/queues", nil)
	router.ServeHTTP(resp, req)
	

	assert.True(t, resp.Code == 200 || resp.Code == 204, "Esperado 200 o 204. Se obtuvo %d", resp.Code)

	var response map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.Equal(t, "Success!", response["message"]) 
	assert.NotNil(t, response["data"]) 

	dataBytes, err := json.Marshal(response["data"])
    assert.NoError(t, err)

    var data []model.Queue
    err = json.Unmarshal(dataBytes, &data)
    assert.NoError(t, err)

 
}

func TestGetAllSessions(t *testing.T) {
	router :=  router.SetupRouter()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/sessions", nil)
	router.ServeHTTP(resp, req)
	

	assert.True(t, resp.Code == 200 || resp.Code == 204, "Esperado 200 o 204. Se obtuvo %d", resp.Code)

	var response map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	

	assert.Equal(t, "Success!", response["message"]) 
	assert.NotNil(t, response["data"]) 

	dataBytes, err := json.Marshal(response["data"])
    assert.NoError(t, err)

    var data []model.Session
    err = json.Unmarshal(dataBytes, &data)
    assert.NoError(t, err)


}

func TestCreatePlayer(t *testing.T) {
	router :=  router.SetupRouter()

	requestBody := map[string]interface{}{
		"name": "scavender",
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	assert.NoError(t, err)


	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/players", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)
	
	assert.Equal(t, 201, resp.Code)

	var response map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	

	assert.Equal(t, "Success!", response["message"]) 
	assert.NotNil(t, response["data"]) 

	dataBytes, err := json.Marshal(response["data"])
    assert.NoError(t, err)

    var data []model.Player
    err = json.Unmarshal(dataBytes, &data)
    assert.NoError(t, err)

}

func TestCreateQueue(t *testing.T) {
	router :=  router.SetupRouter()

	requestBody := map[string]interface{}{
		"name": "CustomBattle",
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	assert.NoError(t, err)


	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/queues", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)
	
	assert.Equal(t, 201, resp.Code)

	var response map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	

	assert.Equal(t, "Success!", response["message"]) 
	assert.NotNil(t, response["data"]) 

	dataBytes, err := json.Marshal(response["data"])
    assert.NoError(t, err)

    var data []model.Queue
    err = json.Unmarshal(dataBytes, &data)
    assert.NoError(t, err)

}