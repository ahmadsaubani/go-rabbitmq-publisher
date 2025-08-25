package routes

import (
	"publisher-topic/src/providers"

	"github.com/gin-gonic/gin"
)

func API(ginEngine *gin.Engine) *gin.Engine {
	app := providers.Register()

	// gin.DisableConsoleColor()
	// r := gin.Default()

	// Auth routes
	v1 := ginEngine.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", app.Controllers.LoginController.LoginController())
			auth.POST("/register", app.Controllers.RegisterController.RegisterController())
			auth.GET("/profile", app.Controllers.LoginController.GetProfileController())
			// auth.POST("/login", auth_controllers.LoginController())
			// auth.POST("/register", auth_controllers.RegisterController())
			// auth.GET("/profile", auth_controllers.GetProfileController())
		}

		v1.GET("/threads", app.Controllers.ThreadController.GetAllThreadController())
		// v1.GET("/threads", thread_controllers.GetAllThreadController())

		thread := v1.Group("/thread")
		{
			thread.POST("/", app.Controllers.ThreadController.CreateThreadController())
			thread.POST("/like", app.Controllers.ThreadController.LikeThreadController())
			thread.GET("/detail", app.Controllers.ThreadController.GetDetailThreadController())
			// thread.POST("/", thread_controllers.CreateThreadController())
			// thread.POST("/like", thread_controllers.LikeThreadController())
			// thread.GET("/detail", thread_controllers.GetDetailThreadController())

			comment := thread.Group("/comment")
			{
				comment.POST("/", app.Controllers.CommentController.CreateCommentController())
				// comment.POST("/", comment_controllers.CreateCommentController())

			}
		}

	}

	return ginEngine
}
