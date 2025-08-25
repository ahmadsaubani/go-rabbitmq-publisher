package auth_controllers

import (
	"context"
	"fmt"
	"net/http"
	"publisher-topic/src/dtos/auths/logins"
	"publisher-topic/src/helpers"
	"publisher-topic/src/services"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginControllerInterface interface {
	LoginController() gin.HandlerFunc
	GetProfileController() gin.HandlerFunc
}

type loginController struct {
	Service services.ServiceProvider
}

func NewLoginController(service services.ServiceProvider) LoginControllerInterface {
	return loginController{
		Service: service,
	}
}

func (l loginController) LoginController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest logins.LoginRequestDto
		if err := c.ShouldBind(&loginRequest); err != nil {
			helpers.ErrorResponse(c, "Invalid request data", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, _ := l.Service.LoginService.Login(ctx, loginRequest)
		if success, ok := response["success"].(bool); ok && !success {

			helpers.ErrorResponse(c, fmt.Sprintf("login failed: %v", response["message"]), http.StatusUnauthorized)
			return
		}

		helpers.SuccessResponse(c, "OK", response, http.StatusOK)
	}
}

func (l loginController) GetProfileController() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ErrorResponse(c, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, _ := l.Service.LoginService.GetProfile(ctx, authHeader)
		if success, ok := response["success"].(bool); ok && !success {
			helpers.ErrorResponse(c, fmt.Sprintf("login failed: %v", response["message"]), http.StatusUnauthorized)
			return
		}

		helpers.SuccessResponse(c, "OK", response, http.StatusOK)
	}
}
