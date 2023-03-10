package dto

type CommentRequest struct {
	Content string `json:"content" binding:"required"`
	BlogID  uint64 `json:"blog_id"`
	UserID  uint64 `json:"user_id"`
}
