package test

import (
	"api-botnet/cmd/router"
	"api-botnet/database"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	database.Connect()
	database.Migrate()

	// Run tests
	code := m.Run()

	// Exit
	os.Exit(code)
}

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

// ****** Bots tests ******
func TestGetAllBots(t *testing.T) {
	// Router configuration
	router := router.SetupRouter()
	// Petition creation
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bots", nil)

	// Send petition
	router.ServeHTTP(w, req)

	// Assert code
	assert.Equal(t, 200, w.Code)

	// Default data on database
	expected := "[{\"id\":1,\"ip\":\"1.0.0.0\"}]"

	assert.JSONEq(t, expected, w.Body.String())
}
