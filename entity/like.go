package entity

import "go-blogrpl/utils"

type BlogLike struct {
	utils.Model
	Count  int    `json:"count"`
	Users  []User `gorm:"many2many:users_bloglikes;" json:"users,omitempty"`
	BlogID uint64 `gorm:"foreignKey" json:"blog_id" binding:"required"`
	Blog   *Blog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog,omitempty"`
}

type CommentLike struct {
	utils.Model
	Count     int      `json:"count"`
	Users     []User   `gorm:"many2many:users_commentlikes;" json:"users,omitempty"`
	CommentID uint64   `gorm:"foreignKey" json:"comment_id" binding:"required"`
	Comment   *Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment,omitempty"`
}
