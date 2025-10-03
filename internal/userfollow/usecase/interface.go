package usecase

import "github.com/MingPV/UserService/internal/entities"

type UserFollowUseCase interface {
	FollowUser(userID, followTo string) (*entities.UserFollow, error)
	UnfollowUser(userID, followTo string) error
	FindAllFollowers(followTo string) ([]*entities.UserFollow, error)
	FindAllFollowings(userID string) ([]*entities.UserFollow, error)
}