package service

import (
	dao "crud-app/dao"
	dto "crud-app/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// FarmList godoc
// @Summary List Farms
// @Schemes
// @Description Get the list of all farms
// @Tags DELOS CRUD-APP FARMS
// @Produce json
// @Success 200 {array} dto.FarmListResponse
// @Failure 404 {object} dto.FarmListErrorResponse
// @Router /farm/list [get]
func FarmList(ctx *gin.Context) {
    farms, err := dao.FarmList()
    if err != nil {
        if err.Error() == "FARMS_NOT_FOUND" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "No farms found"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, farms)
}

// FarmListById godoc
// @Summary Get Farm by ID
// @Schemes
// @Description Get details of a farm by its ID
// @Tags DELOS CRUD-APP FARMS
// @Produce json
// @Param id path string true "Farm ID"
// @Success 200 {object} dto.FarmResponse
// @Failure 404 {object} dto.FarmListErrorResponse
// @Router /farm/list/{id} [get]
func FarmListById(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    farm, err := dao.FarmListById(id)
    if err != nil {
        if err.Error() == "FARM_NOT_FOUND" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Farm not found"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, farm)
}

// CreateFarm godoc
// @Summary Create a new Farm
// @Schemes
// @Description Create a new farm with a unique name
// @Tags DELOS CRUD-APP FARMS
// @Accept json
// @Produce json
// @Param CreateFarm body dto.CreateFarm true "Create Farm Payload"
// @Success 201 {object} dto.CreateFarmResponse
// @Failure 409 {object} dto.CreateFarmErrorResponse
// @Router /farm/create [post]
func CreateFarm(ctx *gin.Context) {
    var createFarmDto dto.CreateFarm
    if err := ctx.ShouldBindJSON(&createFarmDto); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    farm, err := dao.CreateFarm(createFarmDto)
    if err != nil {
        if err.Error() == "DUPLICATE_FARM_NAME" {
            ctx.JSON(http.StatusConflict, gin.H{"error": "Duplicate farm name"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, farm)
}

// UpdateFarm godoc
// @Summary Update a Farm
// @Schemes
// @Description Update a farm by ID or create a new farm if the ID does not exist
// @Tags DELOS CRUD-APP FARMS
// @Accept json
// @Produce json
// @Param UpdateFarm body dto.UpdateFarm true "Update Farm Payload"
// @Success 200 {object} dto.UpdateFarmResponse
// @Success 201 {object} dto.CreateFarmResponse
// @Failure 400 {object} dto.UpdateFarmErrorResponse
// @Failure 500 {object} dto.UpdateFarmErrorResponse
// @Router /farm/update [put]
func UpdateFarm(ctx *gin.Context) {
    var updateFarmDto dto.UpdateFarm

    // Bind JSON payload to updateFarmDto
    if err := ctx.ShouldBindJSON(&updateFarmDto); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
        return
    }

    // Validate UUID
    if _, err := uuid.Parse(updateFarmDto.ID.String()); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    // Call DAO to update farm
    rowsAffected, err := dao.UpdateFarm(updateFarmDto)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if rowsAffected == 0 {
        // Create a new farm if no rows were affected
        newFarmDto := dto.CreateFarm{Name: updateFarmDto.Name}
        newFarm, err := dao.CreateFarm(newFarmDto)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusCreated, newFarm)
        return
    }

    ctx.JSON(http.StatusOK, updateFarmDto)
}

// DeleteFarm godoc
// @Summary Delete a Farm
// @Schemes
// @Description Delete a farm by ID
// @Tags DELOS CRUD-APP FARMS
// @Produce json
// @Param id path string true "Farm ID"
// @Success 200 {object} dto.DeleteFarmResponse
// @Failure 400 {object} dto.DeleteFarmErrorResponse
// @Failure 404 {object} dto.DeleteFarmErrorResponse
// @Router /farm/delete/{id} [delete]
func DeleteFarm(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    err = dao.DeleteFarm(id)
    if err != nil {
        if err.Error() == "FARM_NOT_FOUND_OR_ALREADY_DELETED" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Farm deleted successfully"})
}