package entity

import "go-blogrpl/utils"

type Blog struct {
	utils.Model
	Title       string     `json:"title" binding:"required"`
	Slug        string     `json:"slug" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Content     string     `json:"content" binding:"required"`
	Comments    []Comment  `json:"comments,omitempty"`
	Likes       []BlogLike `json:"likes,omitempty"`
	UserID      uint64     `gorm:"foreignKey" json:"user_id" binding:"required"`
	User        *User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
}
