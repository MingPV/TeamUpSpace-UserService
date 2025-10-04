package usecase

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/userfollow/repository"
	"github.com/google/uuid"
)

type UserFollowService struct {
	repo repository.UserFollowRepository
}

func NewUserFollowService(repo repository.UserFollowRepository) UserFollowUseCase {
	return &UserFollowService{repo: repo}
}

func (s *UserFollowService) FollowUser(userID, followTo string) (*entities.UserFollow, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	ft, err := uuid.Parse(followTo)
	if err != nil {
		return nil, err
	}
	uf := &entities.UserFollow{UserID: uid, FollowTo: ft}
	if err := s.repo.Save(uf); err != nil {
		return nil, err
	}
	return uf, nil
}

func (s *UserFollowService) UnfollowUser(userID, followTo string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	ft, err := uuid.Parse(followTo)
	if err != nil {
		return err
	}
	return s.repo.Delete(uid, ft)
}

func (s *UserFollowService) FindAllFollowers(followTo string) ([]*entities.UserFollow, error) {
	ft, err := uuid.Parse(followTo)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAllFollowers(ft)
}

func (s *UserFollowService) FindAllFollowings(userID string) ([]*entities.UserFollow, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAllFollowings(uid)
}