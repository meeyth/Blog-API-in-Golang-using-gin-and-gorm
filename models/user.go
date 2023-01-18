package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"primarykey" json:"user_id"`
	UserName string    `gorm:"type:varchar(255);not null;unique" json:"username"`
	Email    string    `gorm:"type:varchar(255);not null" json:"-"`
	Password []byte    `gorm:"not null" json:"-"`
	Birthday time.Time `json:"birthday"`
	JoinedAt time.Time `json:"joinedAt"`
	Blog     []Blog    `gorm:"foreignKey:creator;references:user_name;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

// HOOKS

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	username := u.UserName
	fmt.Println("username", username)
	tx.Where("creator = ?", username).Delete(&Blog{})

	return nil
}
