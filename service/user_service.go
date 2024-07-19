package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
	"crud-app/dto"
	"crud-app/config"
)

func SignUp(ctx *gin.Context) {
	// Bind the request body to the SignUpRequest struct
	var requestBody dto.SignUpRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hit the Supabase API to sign up the user
	supabase := config.SetupSupabase()
	user, err := supabase.Auth.SignUp(ctx, supa.UserCredentials{
		Email: requestBody.Email,
		Password: requestBody.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the user data
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func SignIn(ctx *gin.Context) {
	var requestBody dto.SignUpRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supabase := config.SetupSupabase()
	user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email: requestBody.Email,
		Password: requestBody.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}