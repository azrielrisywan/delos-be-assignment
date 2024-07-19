package config

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetupRouter() *gin.Engine {
	if router == nil {
        router = gin.Default()
    }
    return router
}