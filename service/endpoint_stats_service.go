package service

import (
    "crud-app/dao"
    "github.com/gin-gonic/gin"
    "net/http"
)

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
