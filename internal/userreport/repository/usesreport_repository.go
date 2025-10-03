package repository

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/google/uuid"
)

type UserFollowRepository interface {
	Save(follow *entities.UserFollow) error
	Delete(userID, followTo uuid.UUID) error
	FindAllFollowers(followTo uuid.UUID) ([]*entities.UserFollow, error)
	FindAllFollowings(userID uuid.UUID) ([]*entities.UserFollow, error)
}