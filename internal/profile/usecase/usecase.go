package usecase

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/profile/repository"
)

// ProfileService
type ProfileService struct {
	repo repository.ProfileRepository
}

// Init ProfileService function
func NewProfileService(repo repository.ProfileRepository) ProfileUseCase {
	return &ProfileService{repo: repo}
}

// ProfileService Methods - 1 create
func (s *ProfileService) CreateProfile(profile *entities.Profile) error {
	if err := s.repo.Save(profile); err != nil {
		return err
	}
	return nil
}

// ProfileService Methods - 2 find all
func (s *ProfileService) FindAllProfiles() ([]*entities.Profile, error) {
	profiles, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

// ProfileService Methods - 3 find by user_id
func (s *ProfileService) FindProfileByID(user_id string) (*entities.Profile, error) {

	profile, err := s.repo.FindByID(user_id)
	if err != nil {
		return &entities.Profile{}, err
	}
	return profile, nil
}

// ProfileService Methods - 4 patch
func (s *ProfileService) PatchProfile(user_id string, profile *entities.Profile) (*entities.Profile, error) {

	if err := s.repo.Patch(user_id, profile); err != nil {
		return nil, err
	}
	updatedProfile, _ := s.repo.FindByID(user_id)

	return updatedProfile, nil
}

// ProfileService Methods - 5 delete
func (s *ProfileService) DeleteProfile(user_id string) error {
	if err := s.repo.Delete(user_id); err != nil {
		return err
	}
	return nil
}
