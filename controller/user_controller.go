package controller

import (
	"crud-app/config"
	"crud-app/service"
	"crud-app/docs"

	swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func UserController () {
	r := config.SetupRouter()
	docs.SwaggerInfo.BasePath = ""

	// Sign Up
	r.POST("/signup", service.SignUp)

	// Sign In
	r.POST("/signin", service.SignIn)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}