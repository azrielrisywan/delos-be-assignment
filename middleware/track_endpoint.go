package middleware

import (
	"crud-app/dao"
	"crud-app/dto"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TrackUsage() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        usage := dto.EndpointStats{
            Endpoint:    ctx.Request.URL.Path,
            Method:      ctx.Request.Method,
            UserAgent:   ctx.Request.UserAgent(),
            RequestTime: time.Now().UTC(),
        }

		fmt.Printf("Received request to %s %s from %s\n", usage.Method, usage.Endpoint, usage.UserAgent)

        // Log the request to the database
        err := dao.LogEndpointUsage(usage)
        if err != nil {
            ctx.JSON(500, gin.H{"error": "Failed to log request"})
            ctx.Abort()
            return
        }

        ctx.Next()
    }
}
