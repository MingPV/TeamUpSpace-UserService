package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserFollowCreatedEvent struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey"` // the follower
	FollowTo  uuid.UUID `gorm:"type:uuid;primaryKey"` // the followed
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
