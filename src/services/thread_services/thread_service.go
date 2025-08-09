package thread_services

import (
	"context"
	"fmt"
	"publisher-topic/src/dtos/commons"
	"publisher-topic/src/dtos/threads"
	"publisher-topic/src/utils/rabbitmqs"
)

func CreateThreadService(ctx context.Context, threadRequestDto threads.ThreadRequestDto) (map[string]interface{}, error) {
	// Payload untuk request profile
	payload := map[string]interface{}{
		"token":       threadRequestDto.Token,
		"title":       threadRequestDto.Title,
		"description": threadRequestDto.Description,
	}

	// Kirim RPC dan langsung terima hasil unmarshal dalam map
	resp, err := rabbitmqs.PublishMessage(ctx, "thread.create.request", "", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}

func GetAllThreadService(ctx context.Context, tokenRequestDto commons.TokenRequestDto) (map[string]interface{}, error) {
	// Payload untuk request profile
	payload := map[string]interface{}{
		"token": tokenRequestDto.Token,
	}

	// Kirim RPC dan langsung terima hasil unmarshal dalam map
	resp, err := rabbitmqs.PublishMessage(ctx, "thread.getAll.request", "", payload)
	fmt.Println("resp", resp)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}

func GetDetailThreadService(ctx context.Context, threadDetailRequestDto threads.ThreadDetailRequestDto) (map[string]interface{}, error) {

	payload := map[string]interface{}{
		"token":     threadDetailRequestDto.Token,
		"thread_id": threadDetailRequestDto.ThreadID,
	}

	resp, err := rabbitmqs.PublishMessage(ctx, "thread.getDetail.request", "", payload)
	fmt.Println("resp", resp)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}

func LikeThreadService(ctx context.Context, likeThreadRequestDto threads.LikeThreadRequestDto) (map[string]interface{}, error) {

	payload := map[string]interface{}{
		"token":     likeThreadRequestDto.Token,
		"thread_id": likeThreadRequestDto.ThreadID,
	}

	resp, err := rabbitmqs.PublishMessage(ctx, "thread.like.request", "", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}
