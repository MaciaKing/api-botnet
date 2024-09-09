package test

import (
	"api-botnet/cmd/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingEndpoint(t *testing.T) {
	// Router configuration
	router := router.SetupRouter()

	// Petition creation
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Send petition
	router.ServeHTTP(w, req)

	// Assert code
	assert.Equal(t, 200, w.Code)

	// Verify body message
	expectedBody := `{"message":"pong"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
