package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
	"crud-app/dto"
	"crud-app/config"
)

// SignUp godoc
// @Summary Sign Up
// @Schemes
// @Description Sign Up using email and password
// @Tags DELOS AUTH-APP
// @Accept json
// @Produce json
// @Param SignUpRequest body dto.SignUpRequest true "Sign Up Payload"
// @Success 200 {object} dto.SignUpResponse
// @Router /signup [post]
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

// Sign In godoc
// @Summary Sign In
// @Schemes
// @Description Sign In using email and password if you have signed up before
// @Tags DELOS AUTH-APP
// @Accept json
// @Produce json
// @Param SignInRequest body dto.SignInRequest true "Sign In Payload"
// @Success 200 {object} dto.SignInResponse
// @Router /signin [post]
func SignIn(ctx *gin.Context) {
	var requestBody dto.SignInRequest
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