package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, message string, data interface{}, httpCode int) {
	c.JSON(httpCode, gin.H{
		"status":     true,
		"message":    message,
		"data":       data,
		"code":       httpCode,
		"accessTime": time.Now().UTC().Format(time.RFC3339),
	})
}

func ErrorResponse(c *gin.Context, message string, httpCode int) {
	c.JSON(httpCode, gin.H{
		"status":  false,
		"message": message,
		"data":    nil,
		// "data": gin.H{
		// 	"message": message,
		// },
		"code":       httpCode,
		"accessTime": time.Now().UTC().Format(time.RFC3339),
	})
}
