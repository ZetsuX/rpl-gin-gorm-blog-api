package entity

import "go-blogrpl/utils"

type BlogLike struct {
	utils.Model
	UserID uint64 `gorm:"foreignKey" json:"user_id" binding:"required"`
	User   *Blog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	BlogID uint64 `gorm:"foreignKey" json:"blog_id" binding:"required"`
	Blog   *Blog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog,omitempty"`
}

type CommentLike struct {
	utils.Model
	UserID    uint64   `gorm:"foreignKey" json:"user_id" binding:"required"`
	User      *Blog    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	CommentID uint64   `gorm:"foreignKey" json:"comment_id" binding:"required"`
	Comment   *Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment,omitempty"`
}
