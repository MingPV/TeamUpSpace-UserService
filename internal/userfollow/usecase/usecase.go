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

func (s *UserFollowService) CreateUserFollow(follow *entities.UserFollow) (*entities.UserFollow, error) {
	if err := s.repo.Save(follow); err != nil {
		return nil, err
	}
	return follow, nil
}

func (s *UserFollowService) FindUserFollowByID(id int) (*entities.UserFollow, error) {
	return s.repo.FindByID(id)
}

func (s *UserFollowService) FindAllByUser(userID string) ([]*entities.UserFollow, error) {
	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAllByUser(parsedID)
}

func (s *UserFollowService) FindAllByFollowTo(followTo string) ([]*entities.UserFollow, error) {
	parsedID, err := uuid.Parse(followTo)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAllByFollowTo(parsedID)
}

func (s *UserFollowService) FindAllUserFollows() ([]*entities.UserFollow, error) {
	return s.repo.FindAll()
}

func (s *UserFollowService) DeleteUserFollow(id int) error {
	return s.repo.Delete(id)
}