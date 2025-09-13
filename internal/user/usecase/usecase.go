package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/MingPV/UserService/internal/entities"
	profileRepo "github.com/MingPV/UserService/internal/profile/repository"
	userRepo "github.com/MingPV/UserService/internal/user/repository"
	"github.com/MingPV/UserService/pkg/apperror"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

// UserService struct
type UserService struct {
	userRepository    userRepo.UserRepository
	profileRepository profileRepo.ProfileRepository
}

// Init UserService
func NewUserService(userRepository userRepo.UserRepository, profileRepository profileRepo.ProfileRepository) UserUseCase {
	return &UserService{userRepository: userRepository, profileRepository: profileRepository}
}

// UserService Methods - 1 Register user (hash password)
func (s *UserService) Register(user *entities.User, profile *entities.Profile) (*entities.User, error) {
	existingUser, _ := s.userRepository.FindByEmail(user.Email)
	if existingUser != nil {
		return nil, apperror.ErrAlreadyExists
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPwd)

	// Insert profile first because profile has UserID
	if err := s.profileRepository.Save(profile); err != nil {
		return nil, err
	}
	if err := s.userRepository.Save(user); err != nil {
		return nil, err
	}

	createdUser, err := s.userRepository.FindByID(user.ID.String())
	if err != nil {
		return nil, apperror.ErrInternalServer
	}

	return createdUser, nil
}

// UserService Methods - 2 Login user (check email + password)
func (s *UserService) Login(email string, password string) (string, *entities.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil || user == nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, err
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_info": user,
		"exp":       time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

// UserService Methods - 3 Get user by id
func (s *UserService) FindUserByID(id string) (*entities.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *UserService) FindUserByEmail(email string) (*entities.User, error) {
	return s.userRepository.FindByEmail(email)
}

func (s *UserService) FindUserByUsername(username string) (*entities.User, error) {
	return s.userRepository.FindByUsername(username)
}

// UserService Methods - 4 Get all users
func (s *UserService) FindAllUsers() ([]*entities.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UserService Methods - 5 Get user by email
func (s *UserService) GetUserByEmail(email string) (*entities.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserService Methods - 6 Patch
func (s *UserService) PatchUser(id string, user *entities.User) (*entities.User, error) {
	if err := s.userRepository.Patch(id, user); err != nil {
		return nil, err
	}
	updatedUser, _ := s.userRepository.FindByID(id)

	return updatedUser, nil
}

// UserService Methods - 7 Delete
func (s *UserService) DeleteUser(id string) error {
	if err := s.userRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *UserService) LoginOrRegisterWithGoogle(userInfo map[string]interface{}, token *oauth2.Token) (string, *entities.User, error) {
	email, _ := userInfo["email"].(string)
	name, _ := userInfo["name"].(string)
	picture, _ := userInfo["picture"].(string)
	fmt.Println(userInfo)
	// googleID, _ := userInfo["id"].(string)

	// find user by email in database
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		// not found -> create
		user_id := uuid.New()
		user = &entities.User{
			ID:       user_id,
			Email:    email,
			Username: user_id.String(),
			Password: "",
		}
		profile := &entities.Profile{
			UserID:      user.ID,
			DisplayName: name,
			ProfileURL:  picture,
		}

		// Insert profile first because profile has UserID
		if err := s.profileRepository.Save(profile); err != nil {
			return "", nil, err
		}
		if err := s.userRepository.Save(user); err != nil {
			return "", nil, err
		}
	}
	user.Password = ""

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_info": user,
		"exp":       time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	loggedInUser, err := s.userRepository.FindByID(user.ID.String())
	if err != nil {
		return "", nil, apperror.ErrInternalServer
	}

	return tokenString, loggedInUser, nil
}
