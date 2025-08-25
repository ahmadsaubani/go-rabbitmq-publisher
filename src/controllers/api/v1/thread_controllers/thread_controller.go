package thread_controllers

import (
	"context"
	"fmt"
	"net/http"
	"publisher-topic/src/dtos/commons"
	"publisher-topic/src/dtos/threads"
	"publisher-topic/src/helpers"
	"publisher-topic/src/services"
	"time"

	"github.com/gin-gonic/gin"
)

type ThreadControllerInterface interface {
	CreateThreadController() gin.HandlerFunc
	GetAllThreadController() gin.HandlerFunc
	GetDetailThreadController() gin.HandlerFunc
	LikeThreadController() gin.HandlerFunc
}

type ThreadController struct {
	service services.ServiceProvider
}

func NewThreadController(service services.ServiceProvider) ThreadControllerInterface {
	return ThreadController{
		service: service,
	}
}

func (tc ThreadController) CreateThreadController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var threadRequest threads.ThreadRequestDto
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ErrorResponse(c, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if err := c.ShouldBind(&threadRequest); err != nil {
			fmt.Println("Error binding thread request:", err)
			helpers.ErrorResponse(c, "Invalid request data", http.StatusBadRequest)
			return
		}

		threadRequest.Token = authHeader

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		response, _ := tc.service.ThreadService.CreateThreadService(ctx, threadRequest)
		if success, ok := response["success"].(bool); ok && !success {
			helpers.ErrorResponse(c, fmt.Sprintf("login failed: %v", response["message"]), http.StatusUnauthorized)
			return
		}

		helpers.SuccessResponse(c, "OK", response, http.StatusOK)
	}
}

func (tc ThreadController) GetAllThreadController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenRequest commons.TokenRequestDto

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ErrorResponse(c, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenRequest.Token = authHeader

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, _ := tc.service.ThreadService.GetAllThreadService(ctx, tokenRequest)
		fmt.Println(response)
		if success, ok := response["success"].(bool); ok && !success {
			helpers.ErrorResponse(c, fmt.Sprintf("failed to get threads: %v", response["message"]), http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponse(c, "OK", response, http.StatusOK)
	}
}

func (tc ThreadController) GetDetailThreadController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenRequest threads.ThreadDetailRequestDto

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ErrorResponse(c, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenRequest.Token = authHeader

		threadID := c.Query("thread_id")
		if threadID == "" {
			helpers.ErrorResponse(c, "Missing thread_id in query params", http.StatusBadRequest)
			return
		}
		tokenRequest.ThreadID = threadID

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, _ := tc.service.ThreadService.GetDetailThreadService(ctx, tokenRequest)
		fmt.Println(response)
		if success, ok := response["success"].(bool); ok && !success {
			helpers.ErrorResponse(c, fmt.Sprintf("failed to get threads: %v", response["message"]), http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponse(c, "OK", response, http.StatusOK)
	}
}

func (tc ThreadController) LikeThreadController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var likeRequest threads.LikeThreadRequestDto
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ErrorResponse(c, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if err := c.ShouldBind(&likeRequest); err != nil {
			helpers.ErrorResponse(c, "Invalid request data", http.StatusBadRequest)
			return
		}

		likeRequest.Token = authHeader

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, _ := tc.service.ThreadService.LikeThreadService(ctx, likeRequest)
		if success, ok := response["success"].(bool); ok && !success {
			helpers.ErrorResponse(c, fmt.Sprintf("login failed: %v", response["message"]), http.StatusUnauthorized)
			return
		}

		helpers.SuccessResponse(c, "OK", response, http.StatusOK)
	}
}
