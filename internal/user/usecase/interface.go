package usecase

import "github.com/MingPV/UserService/internal/entities"

type UserUseCase interface {
	Register(user *entities.User) error
	Login(email, password string) (string, *entities.User, error)
	FindUserByID(id string) (*entities.User, error)
	FindAllUsers() ([]*entities.User, error)
	PatchUser(id string, user *entities.User) (*entities.User, error)
	DeleteUser(id string) error
}
