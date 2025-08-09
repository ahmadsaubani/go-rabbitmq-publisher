package register_services

import (
	"context"
	"fmt"
	"publisher-topic/src/dtos/auths/registers"
	"publisher-topic/src/utils/rabbitmqs"
)

func RegisterService(ctx context.Context, registerRequest registers.RegisterRequestDto) (map[string]interface{}, error) {
	resp, err := rabbitmqs.PublishMessage(ctx, "auth.register.request", "", registerRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}
