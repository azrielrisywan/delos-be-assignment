package controller

import (
    "crud-app/config"
    "crud-app/service"
    "crud-app/middleware"
)

func StatisticsController() {
    var hmacSecret = "pDnYuxHNGugqD6u/q20ShEFX32uIDNFTPH3CjLZjPSES/N7QvZr+v+eDOCi31F7FbQFrzCgLqngGUolnvUXzqw=="

    r := config.SetupRouter()
    r.GET("/stats", middleware.AuthMiddleware(hmacSecret), service.GetEndpointStats)
}