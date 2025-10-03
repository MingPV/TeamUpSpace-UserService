package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserFollow struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	FollowTo  uuid.UUID `gorm:"type:uuid;not null" json:"follow_to"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	User      *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user"`
	Followed  *User `gorm:"foreignKey:FollowTo;references:ID;constraint:OnDelete:CASCADE" json:"followed"`
}