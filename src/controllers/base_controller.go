package controllers

import (
	"publisher-topic/src/controllers/api/v1/auth_controllers"
	"publisher-topic/src/controllers/api/v1/comment_controllers"
	"publisher-topic/src/controllers/api/v1/thread_controllers"
	"publisher-topic/src/services"
)

type BaseController struct {
	LoginController    auth_controllers.LoginControllerInterface
	RegisterController auth_controllers.RegisterControllerInterface
	CommentController  comment_controllers.CommentControllerInterface
	ThreadController   thread_controllers.ThreadControllerInterface
}

func InitControllers(services services.ServiceProvider) BaseController {
	return BaseController{
		LoginController:    auth_controllers.NewLoginController(services),
		RegisterController: auth_controllers.NewRegisterController(services),
		CommentController:  comment_controllers.NewCommentController(services),
		ThreadController:   thread_controllers.NewThreadController(services),
	}
}
