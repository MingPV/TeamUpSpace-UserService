package usecase

import (
	"github.com/MingPV/UserService/internal/entities"
	"golang.org/x/oauth2"
)

type UserUseCase interface {
	Register(user *entities.User, profile *entities.Profile) (*entities.User, error)
	Login(email, password string) (string, *entities.User, error)
	LoginOrRegisterWithGoogle(userInfo map[string]interface{}, token *oauth2.Token) (string, *entities.User, error)
	FindUserByID(id string) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	FindUserByUsername(username string) (*entities.User, error)
	FindAllUsers() ([]*entities.User, error)
	PatchUser(id string, user *entities.User) (*entities.User, error)
	DeleteUser(id string) error
}
