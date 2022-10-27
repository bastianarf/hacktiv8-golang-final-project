package dto

type Comment struct {
	Message string `validate:"required"`
	PhotoID uint   `validate:"required" json:"photo_id"`
}

type CommentMessage struct {
	Message string `validate:"required"`
}
