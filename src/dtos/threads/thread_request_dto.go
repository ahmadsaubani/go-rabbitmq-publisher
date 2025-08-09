package threads

type ThreadRequestDto struct {
	Token       string `form:"token" json:"token"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

type ThreadDetailRequestDto struct {
	Token    string `form:"token" json:"token"`
	ThreadID string `form:"thread_id" json:"thread_id" binding:"required"`
}

type LikeThreadRequestDto struct {
	Token    string `form:"token" json:"token"`
	ThreadID string `form:"thread_id" json:"thread_id" binding:"required"`
}
