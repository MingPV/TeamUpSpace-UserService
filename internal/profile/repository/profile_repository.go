package repository

import "github.com/MingPV/UserService/internal/entities"

type ProfileRepository interface {
	Save(profile *entities.Profile) error
	FindAll() ([]*entities.Profile, error)
	FindByID(user_id string) (*entities.Profile, error)
	Patch(user_id string, profile *entities.Profile) error
	Delete(user_id string) error
}
