package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string    `json:"password"`
	IsAdmin   bool      `gorm:"default:false" json:"is_admin"`
	IsBan     bool      `gorm:"default:false" json:"is_ban"`
	BanUntil  time.Time `json:"ban_until"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Profile Profile `gorm:"foreignKey:ID;references:UserID" json:"profile"` // this.ID -> Profile.UserID
}
