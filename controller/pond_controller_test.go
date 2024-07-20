package controller

import (
	"bytes"
	"crud-app/config"
	"crud-app/middleware"
	"crud-app/service"
	// "fmt"
    "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// SetupRouter initializes the router and registers routes
func SetupPondTestRouter() *gin.Engine {
    r := config.SetupRouter()

	// List Pond
	r.GET("/pond/list", middleware.TrackUsage(), service.PondList)

	// List Pond By Id
	r.GET("/pond/list/:id", middleware.TrackUsage(), service.PondListById)

	// Create Pond
	r.POST("/pond/create", middleware.TrackUsage(), service.CreatePond)

    // Create Farm for creating Pond with valid farm_id
    r.POST("/farm/create", middleware.TrackUsage(), service.CreateFarm)

	// Update Pond
	r.PUT("/pond/update", middleware.TrackUsage(), service.UpdatePond)

	// Delete Pond
	r.DELETE("/pond/delete/:id", middleware.TrackUsage(), service.DeletePond)

    return r
}

// TestPondController tests the endpoints of PondController
func TestPondController(t *testing.T) {
    router := SetupPondTestRouter()

    // Create Farm to get a valid farm_id
    uuid := uuid.New() // Generate a random UUID for unique name
    req, _ := http.NewRequest(http.MethodPost, "/farm/create", bytes.NewBufferString(`{"name": "New Farm ` + uuid.String() + `"}`))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code, "Create Farm failed")

    // Parse the response to extract farm_id
    var farmResponse map[string]interface{}
    if err := json.NewDecoder(w.Body).Decode(&farmResponse); err != nil {
        t.Fatalf("Failed to parse farm response: %v", err)
    }
    farmID, ok := farmResponse["id"].(string)
    if !ok || farmID == "" {
        t.Fatalf("Failed to get farm ID from response")
    }

    // Create Pond using the valid farm_id
    reqBody := `{"name": "New Pond ` + uuid.String() + `", "farm_id": "` + farmID + `"}`
    req, _ = http.NewRequest(http.MethodPost, "/pond/create", bytes.NewBufferString(reqBody))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code, "Create Pond failed")

    // Parse the response to extract pond_id
    var pondResponse map[string]interface{}
    if err := json.NewDecoder(w.Body).Decode(&pondResponse); err != nil {
        t.Fatalf("Failed to parse pond response: %v", err)
    }
    pondID, ok := pondResponse["id"].(string)
    if !ok || pondID == "" {
        t.Fatalf("Failed to get pond ID from response")
    }

    // List Pond
    req, _ = http.NewRequest(http.MethodGet, "/pond/list", nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code, "List Pond failed")

    // List Pond By Id
    req, _ = http.NewRequest(http.MethodGet, "/pond/list/"+pondID, nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code, "List Pond By Id failed")

    // Update Pond using the valid farm_id
    updatedName := "Updated Pond"
    reqBody = `{"id": "` + pondID + `", "name": "` + updatedName + `", "farm_id": "` + farmID + `"}`
    req, _ = http.NewRequest(http.MethodPut, "/pond/update", bytes.NewBufferString(reqBody))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code, "Update Pond failed")

    // Delete Pond
    req, _ = http.NewRequest(http.MethodDelete, "/pond/delete/"+pondID, nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code, "Delete Pond failed")
}
