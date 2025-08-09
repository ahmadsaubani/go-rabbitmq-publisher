package comments

type CommentRequestDto struct {
	Token    string  `form:"token" json:"token"`
	ThreadID string  `form:"thread_id" json:"thread_id" binding:"required"`
	ParentID *string `form:"parent_id" json:"parent_id"`
	Comment  string  `form:"comment" json:"comment" binding:"required"`
}
