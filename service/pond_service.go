package service

import (
    dao "crud-app/dao"
    dto "crud-app/dto"
    "github.com/gin-gonic/gin"
    "net/http"
	"github.com/google/uuid"
)

// PondList godoc
// @Summary List Ponds
// @Schemes
// @Description Get the list of all ponds, including associated farm details
// @Tags DELOS CRUD-APP PONDS
// @Produce json
// @Success 200 {array} dto.Pond
// @Failure 404 {object} dto.PondListErrorResponse
// @Failure 500 {object} dto.PondListErrorResponse
// @Router /pond/list [get]
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

    // Retrieve farm details for each pond
    for i := range ponds {
        farm, err := dao.FarmListById(ponds[i].FarmID)
        if err != nil {
            ponds[i].Farm = nil 
        } else {
            ponds[i].Farm = &farm 
        }
    }

    ctx.JSON(http.StatusOK, ponds)
}

// PondListById godoc
// @Summary Get Pond by ID
// @Schemes
// @Description Get details of a specific pond by its ID, including associated farm details
// @Tags DELOS CRUD-APP PONDS
// @Produce json
// @Param id path string true "Pond ID"
// @Success 200 {object} dto.Pond
// @Failure 400 {object} dto.PondListByIdErrorResponse
// @Failure 404 {object} dto.PondListByIdErrorResponse
// @Failure 500 {object} dto.PondListByIdErrorResponse
// @Router /pond/list/{id} [get]
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

    // Retrieve farm details for the specific pond
    farm, err := dao.FarmListById(pond.FarmID)
    if err != nil {
        pond.Farm = nil
    } else {
        pond.Farm = &farm
    }

    ctx.JSON(http.StatusOK, pond)
}

// CreatePond godoc
// @Summary Create Pond
// @Schemes
// @Description Create a new pond
// @Tags DELOS CRUD-APP PONDS
// @Accept json
// @Produce json
// @Param CreatePond body dto.CreatePond true "Pond creation payload"
// @Success 201 {object} dto.Pond
// @Failure 400 {object} dto.CreatePondErrorResponse
// @Failure 404 {object} dto.CreatePondErrorResponse
// @Failure 409 {object} dto.CreatePondErrorResponse
// @Failure 500 {object} dto.CreatePondErrorResponse
// @Router /pond/create [post]
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

// UpdatePond godoc
// @Summary Update Pond
// @Schemes
// @Description Update an existing pond
// @Tags DELOS CRUD-APP PONDS
// @Accept json
// @Produce json
// @Param UpdatePond body dto.UpdatePond true "Pond update payload"
// @Success 200 {object} dto.Pond
// @Failure 400 {object} dto.UpdatePondErrorResponse
// @Failure 500 {object} dto.UpdatePondErrorResponse
// @Router /pond/update [put]
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

// DeletePond godoc
// @Summary Delete Pond
// @Schemes
// @Description Delete a pond by its ID
// @Tags DELOS CRUD-APP PONDS
// @Param id path string true "Pond ID"
// @Success 200 {object} dto.DeletePondResponse
// @Failure 400 {object} dto.DeletePondErrorResponse
// @Failure 404 {object} dto.DeletePondErrorResponse
// @Failure 500 {object} dto.DeletePondErrorResponse
// @Router /pond/delete/{id} [delete]
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
