package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint           `gorm:"primarykey" json:"post_id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `json:"description"`
	LikeCount   uint           `gorm:"" json:"likeCount"`
	Creator     string         `gorm:"type:varchar(255);not null" json:"creator"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// HOOKS

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	creator := p.Creator
	var user User

	if tx.Where("user_name = ?", creator).First(&user); user.ID == 0 {
		return errors.New("no user found")
	}
	return nil
}
