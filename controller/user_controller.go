package controller

import (
	"crud-app/config"
	"crud-app/service"
)

func UserController () {
	r := config.SetupRouter()

	// Sign Up
	r.POST("/signup", service.SignUp)

	// Sign In
	r.POST("/signin", service.SignIn)
}