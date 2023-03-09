package entity

import "go-blogrpl/utils"

type Comment struct {
	utils.Model
	Content string        `json:"content" binding:"required"`
	Likes   []CommentLike `json:"likes,omitempty"`
	BlogID  uint64        `gorm:"foreignKey" json:"blog_id" binding:"required"`
	Blog    *Blog         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog,omitempty"`
	UserID  uint64        `gorm:"foreignKey" json:"user_id" binding:"required"`
	User    *User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}
