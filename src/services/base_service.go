package services

import (
	"publisher-topic/src/services/auth_services/login_services"
	"publisher-topic/src/services/auth_services/register_services"
	"publisher-topic/src/services/comment_services"
	"publisher-topic/src/services/thread_services"
)

type ServiceProvider struct {
	LoginService    login_services.LoginServiceInterface
	RegisterService register_services.RegisterServiceInterface
	ThreadService   thread_services.ThreadServiceInterface
	CommentService  comment_services.CommentServiceInterface
}

func InitServices() ServiceProvider {
	return ServiceProvider{
		LoginService:    login_services.NewLoginService(),
		RegisterService: register_services.NewRegisterService(),
		ThreadService:   thread_services.NewThreadService(),
		CommentService:  comment_services.NewCommentService(),
	}
}
