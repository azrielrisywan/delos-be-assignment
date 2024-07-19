package service

import (
    dao "crud-app/dao"
    dto "crud-app/dto"
    "github.com/gin-gonic/gin"
    "net/http"
	"github.com/google/uuid"
)

func PondList(ctx *gin.Context) {
    ponds, err := dao.PondList()
    if err != nil {
        if err.Error() == "PONDS_NOT_FOUND" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "No ponds found"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, ponds)
}

func PondListById(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    pond, err := dao.PondListById(id)
    if err != nil {
        if err.Error() == "POND_NOT_FOUND" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Pond not found"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, pond)
}

func CreatePond(ctx *gin.Context) {
    var createPondDto dto.CreatePond
    if err := ctx.ShouldBindJSON(&createPondDto); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
        return
    }
    
    pond, err := dao.CreatePond(createPondDto)
    if err != nil {
        if err.Error() == "DUPLICATE_POND_NAME" {
            ctx.JSON(http.StatusConflict, gin.H{"error": "Duplicate pond name"})
            return
        }
        if err.Error() == "FARM_NOT_FOUND" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Farm not found"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    ctx.JSON(http.StatusCreated, pond)
}

func UpdatePond(ctx *gin.Context) {
    var updatePondDto dto.UpdatePond

    // Bind JSON payload to updatePondDto
    if err := ctx.ShouldBindJSON(&updatePondDto); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
        return
    }

    // Validate UUID
    if _, err := uuid.Parse(updatePondDto.ID.String()); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    // Call DAO to update pond
    rowsAffected, err := dao.UpdatePond(updatePondDto)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if rowsAffected == 0 {
        // Create a new pond if no rows were affected
        newPondDto := dto.CreatePond{Name: updatePondDto.Name, FarmID: updatePondDto.FarmID}
        newPond, err := dao.CreatePond(newPondDto)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusCreated, newPond)
        return
    }

    ctx.JSON(http.StatusOK, updatePondDto)
}

func DeletePond(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    err = dao.DeletePond(id)
    if err != nil {
        if err.Error() == "POND_NOT_FOUND_OR_ALREADY_DELETED" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Pond deleted successfully"})
}
