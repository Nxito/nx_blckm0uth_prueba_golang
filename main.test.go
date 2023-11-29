package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoute1(t *testing.T) {
	// Crea un enrutador Gin para las pruebas
	router := gin.Default()

	// Configura la ruta en el enrutador
	// router.GET("/route1", Route1)

	// Simula una solicitud GET a /route1
	req, err := http.NewRequest("GET", "/route1", nil)
	assert.NoError(t, err)

	// Crea un ResponseRecorder para obtener la respuesta
	w := httptest.NewRecorder()

	// Ejecuta la solicitud a través del enrutador
	router.ServeHTTP(w, req)

	// Verifica el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"message":"Hello from Route 1!"}`
	assert.Equal(t, expectedResponse, w.Body.String())
}