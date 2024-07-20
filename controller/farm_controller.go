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
    r.GET("/farm/list", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.FarmList)

	// List Farm By Id
	r.GET("/farm/list/:id", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.FarmListById)

	// Create Farm
	r.POST("/farm/create", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.CreateFarm)

	// Update Farm
	r.PUT("/farm/update", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.UpdateFarm)

	// Delete Farm
	r.DELETE("/farm/delete/:id", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.DeleteFarm)
}