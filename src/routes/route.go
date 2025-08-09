package routes

import (
	"publisher-topic/src/controllers/api/v1/auth_controllers"
	"publisher-topic/src/controllers/api/v1/comment_controllers"
	"publisher-topic/src/controllers/api/v1/thread_controllers"

	"github.com/gin-gonic/gin"
)

func API(ginEngine *gin.Engine) *gin.Engine {
	// gin.DisableConsoleColor()
	// r := gin.Default()

	// Auth routes
	v1 := ginEngine.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", auth_controllers.LoginController())
			auth.GET("/profile", auth_controllers.GetProfileController())
		}

		v1.GET("/threads", thread_controllers.GetAllThreadController())

		thread := v1.Group("/thread")
		{
			thread.POST("/", thread_controllers.CreateThreadController())
			thread.POST("/like", thread_controllers.LikeThreadController())
			thread.GET("/detail", thread_controllers.GetDetailThreadController())

			comment := thread.Group("/comment")
			{
				comment.POST("/", comment_controllers.CreateCommentController())
				// comment.GET("/comments", thread_controllers.GetAllCommentService())

			}
		}

	}

	// Other routes can be added here

	return ginEngine
}
