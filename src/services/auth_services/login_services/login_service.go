package login_services

import (
	"context"
	"fmt"
	"publisher-topic/src/dtos/auths/logins"
	"publisher-topic/src/utils/rabbitmqs"
)

func LoginService(ctx context.Context, loginRequest logins.LoginRequestDto) (map[string]interface{}, error) {
	resp, err := rabbitmqs.PublishMessage(ctx, "auth.login.request", "", loginRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}

func GetProfileService(ctx context.Context, token string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"token": token,
	}

	resp, err := rabbitmqs.PublishMessage(ctx, "user.profile.request", "", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}
