package auth_controllers

import (
	"context"
	"fmt"
	"net/http"
	"publisher-topic/src/dtos/auths/registers"
	"publisher-topic/src/helpers"
	"publisher-topic/src/services"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterControllerInterface interface {
	RegisterController() gin.HandlerFunc
}

type registerController struct {
	Service services.ServiceProvider
}

func NewRegisterController(service services.ServiceProvider) RegisterControllerInterface {
	return registerController{
		Service: service,
	}
}

func (r registerController) RegisterController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerRequest registers.RegisterRequestDto
		if err := c.ShouldBind(&registerRequest); err != nil {
			helpers.ErrorResponse(c, "Invalid request data", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, _ := r.Service.RegisterService.Register(ctx, registerRequest)
		if success, ok := response["success"].(bool); ok && !success {

			helpers.ErrorResponse(c, fmt.Sprintf("register failed: %v", response["message"]), http.StatusUnauthorized)
			return
		}

		helpers.SuccessResponse(c, "OK", response, http.StatusOK)
	}
}
