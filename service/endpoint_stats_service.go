package service

import (
    "crud-app/dao"
    "github.com/gin-gonic/gin"
    "net/http"
)

// GetEndpointStats godoc
// @Summary Retrieve Endpoint Stats
// @Schemes
// @Description Get the count of how many times each endpoint is called
// @Tags Endpoint Stats
// @Produce json
// @Success 200 {array} dto.EndpointStatsResponse
// @Failure 404 {object} dto.EndpointStatsErrorResponse
// @Router /stats [get]
func GetEndpointStats(ctx *gin.Context) {
    stats, err := dao.GetEndpointStats()
    if err != nil {
        if err.Error() == "no endpoints tracked" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, stats)
}
