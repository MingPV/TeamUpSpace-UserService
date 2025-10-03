package repository

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/google/uuid"
)

type UserFollowRepository interface {
	Save(follow *entities.UserFollow) error
	FindByID(id int) (*entities.UserFollow, error)
	FindAllByUser(userID uuid.UUID) ([]*entities.UserFollow, error)
	FindAllByFollowTo(followTo uuid.UUID) ([]*entities.UserFollow, error)
	FindAll() ([]*entities.UserFollow, error)
	Delete(id int) error
}