package controller

import (
	"bytes"
	"crud-app/config"
	"crud-app/middleware"
	"crud-app/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// SetupRouter initializes the router and registers routes
func SetupFarmTestRouter() *gin.Engine {
    r := config.SetupRouter()

    // List Farm
    r.GET("/farm/list", middleware.TrackUsage(), service.FarmList)

    // List Farm By Id
    r.GET("/farm/list/:id", middleware.TrackUsage(), service.FarmListById)

    // Create Farm
    r.POST("/farm/create", middleware.TrackUsage(), service.CreateFarm)

    // Update Farm
    r.PUT("/farm/update", middleware.TrackUsage(), service.UpdateFarm)

    // Delete Farm
    r.DELETE("/farm/delete/:id", middleware.TrackUsage(), service.DeleteFarm)

    return r
}

// TestFarmController tests the endpoints of FarmController
func TestFarmController(t *testing.T) {
    router := SetupFarmTestRouter()

    // Create Farm
    uuid := uuid.New() // Generate a random UUID for unique name
    reqBody := `{"name": "New Farm ` + uuid.String() + `"}`
    req, _ := http.NewRequest(http.MethodPost, "/farm/create", bytes.NewBufferString(reqBody))
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

    // List Farm
    req, _ = http.NewRequest(http.MethodGet, "/farm/list", nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code, "List Farm failed")

    // List Farm By Id
    req, _ = http.NewRequest(http.MethodGet, "/farm/list/"+farmID, nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code, "List Farm By Id failed")

    // Update Farm
    updatedName := "Updated Farm " + uuid.String()
    reqBody = `{"id": "` + farmID + `", "name": "` + updatedName + `"}`
    req, _ = http.NewRequest(http.MethodPut, "/farm/update", bytes.NewBufferString(reqBody))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code, "Update Farm failed")

    // Delete Farm
    req, _ = http.NewRequest(http.MethodDelete, "/farm/delete/"+farmID, nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code, "Delete Farm failed")
}
