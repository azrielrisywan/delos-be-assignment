package controller

import (
	"crud-app/config"
	"crud-app/service"
	"crud-app/middleware"
)

func FarmController () {
	var hmacSecret = "pDnYuxHNGugqD6u/q20ShEFX32uIDNFTPH3CjLZjPSES/N7QvZr+v+eDOCi31F7FbQFrzCgLqngGUolnvUXzqw=="
	r := config.SetupRouter()

	// List Farm
    r.GET("/farm/list", middleware.AuthMiddleware(hmacSecret), service.FarmList)

	// List Farm By Id
	r.GET("/farm/list/:id", middleware.AuthMiddleware(hmacSecret), service.FarmListById)

	// Create Farm
	r.POST("/farm/create", middleware.AuthMiddleware(hmacSecret), service.CreateFarm)

	// Update Farm
	r.PUT("/farm/update", middleware.AuthMiddleware(hmacSecret), service.UpdateFarm)

	// Delete Farm
	r.DELETE("/farm/delete/:id", middleware.AuthMiddleware(hmacSecret), service.DeleteFarm)
}