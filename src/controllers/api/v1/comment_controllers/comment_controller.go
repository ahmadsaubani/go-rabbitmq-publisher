package comment_controllers

import (
	"context"
	"fmt"
	"net/http"
	"publisher-topic/src/dtos/comments"
	"publisher-topic/src/helpers"
	"publisher-topic/src/services/comment_services"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCommentController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentRequest comments.CommentRequestDto
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ErrorResponse(c, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if err := c.ShouldBind(&commentRequest); err != nil {
			helpers.ErrorResponse(c, "Invalid request data", http.StatusBadRequest)
			return
		}

		commentRequest.Token = authHeader

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, _ := comment_services.CreateCommentService(ctx, commentRequest)
		if success, ok := response["success"].(bool); ok && !success {
			helpers.ErrorResponse(c, fmt.Sprintf("login failed: %v", response["message"]), http.StatusUnauthorized)
			return
		}

		helpers.SuccessResponse(c, "OK", response, http.StatusOK)
	}
}
