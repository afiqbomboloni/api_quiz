package handler

import (
	"net/http"

	"github.com/afiqbomboloni/api_quiz/modules/auth/v1/service"
	"github.com/afiqbomboloni/api_quiz/request"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) Register(ctx *gin.Context) {
	var authRequest request.AuthRequest

	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err := h.authService.SaveUser(authRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	user, _ := h.authService.AuthValidate(authRequest.Email, authRequest.Password)
	token, _ := h.authService.GenerateAccessToken(ctx, user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": gin.H{
			"id":    user.ID,
			"email": authRequest.Email,
			"nama":  authRequest.Nama,
			"token": token,
		},
	})

}

func (h *authHandler) Login(ctx *gin.Context) {
	var request request.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	user, err := h.authService.AuthValidate(request.Email, request.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "error auth",
			"errors":  err.Error(),
		})
		ctx.Abort()
		return
	}

	token, err := h.authService.GenerateAccessToken(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "error 2",
			"errors":  err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Nama,
			"role":  user.Role,
			"email": user.Email,
			"token": token,
		},
		"message": "success",
	})

}

func (h *authHandler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "logout",
	})
}
