package service

import (
	dao "crud-app/dao"
	dto "crud-app/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

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