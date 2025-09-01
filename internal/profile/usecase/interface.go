package usecase

import "github.com/MingPV/UserService/internal/entities"

type ProfileUseCase interface {
	FindAllProfiles() ([]*entities.Profile, error)
	CreateProfile(profile *entities.Profile) error
	PatchProfile(user_id string, profile *entities.Profile) (*entities.Profile, error)
	DeleteProfile(user_id string) error
	FindProfileByID(user_id string) (*entities.Profile, error)
}
