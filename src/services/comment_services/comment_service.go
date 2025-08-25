package comment_services

import (
	"context"
	"fmt"
	"publisher-topic/src/dtos/comments"
	"publisher-topic/src/utils/rabbitmqs"
)

type CommentServiceInterface interface {
	CreateComment(ctx context.Context, commentRequestDto comments.CommentRequestDto) (map[string]interface{}, error)
}

type CommentService struct{}

func NewCommentService() CommentServiceInterface {
	return CommentService{}
}

func (s CommentService) CreateComment(ctx context.Context, commentRequestDto comments.CommentRequestDto) (map[string]interface{}, error) {

	payload := map[string]interface{}{
		"token":     commentRequestDto.Token,
		"thread_id": commentRequestDto.ThreadID,
		"parent_id": commentRequestDto.ParentID,
		"comment":   commentRequestDto.Comment,
	}

	resp, err := rabbitmqs.PublishMessage(ctx, "comment.create.request", "", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	return resp, nil
}
