package usecase

import (
	"github.com/MingPV/UserService/internal/entities"
)

type UserFollowUseCase interface {
	CreateUserFollow(follow *entities.UserFollow) (*entities.UserFollow, error)
	FindUserFollowByID(id int) (*entities.UserFollow, error)
	FindAllByUser(userID string) ([]*entities.UserFollow, error)
	FindAllByFollowTo(followTo string) ([]*entities.UserFollow, error)
	FindAllUserFollows() ([]*entities.UserFollow, error)
	DeleteUserFollow(id int) error
}