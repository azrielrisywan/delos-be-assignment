package controller

import (
	"crud-app/config"
	"crud-app/service"
	"crud-app/middleware"
)

func PondController () {
	var hmacSecret = "pDnYuxHNGugqD6u/q20ShEFX32uIDNFTPH3CjLZjPSES/N7QvZr+v+eDOCi31F7FbQFrzCgLqngGUolnvUXzqw=="
	r := config.SetupRouter()

	// List Pond
	r.GET("/pond/list", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.PondList)

	// List Pond By Id
	r.GET("/pond/list/:id", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.PondListById)

	// Create Pond
	r.POST("/pond/create", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.CreatePond)

	// Update Pond
	r.PUT("/pond/update", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.UpdatePond)

	// Delete Pond
	r.DELETE("/pond/delete/:id", middleware.AuthMiddleware(hmacSecret), middleware.TrackUsage(), service.DeletePond)
}