package register_services

import (
	"context"
	"fmt"
	"publisher-topic/src/dtos/auths/registers"
	"publisher-topic/src/utils/rabbitmqs"
)

type RegisterServiceInterface interface {
	Register(ctx context.Context, registerRequest registers.RegisterRequestDto) (map[string]interface{}, error)
}

type RegisterService struct{}

func NewRegisterService() RegisterServiceInterface {
	return RegisterService{}
}

func (s RegisterService) Register(ctx context.Context, registerRequest registers.RegisterRequestDto) (map[string]interface{}, error) {
	resp, err := rabbitmqs.PublishMessage(ctx, "auth.register.request", "", registerRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}
