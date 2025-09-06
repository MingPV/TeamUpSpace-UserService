package routes

import (
	"os"

	profileRepository "github.com/MingPV/UserService/internal/profile/repository"
	userHandler "github.com/MingPV/UserService/internal/user/handler/rest"
	userRepository "github.com/MingPV/UserService/internal/user/repository"
	userUseCase "github.com/MingPV/UserService/internal/user/usecase"
	middleware "github.com/MingPV/UserService/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app fiber.Router, db *gorm.DB) {

	route := app.Group("/api/v1", middleware.JWTMiddleware())

	// Profile
	profileRepo := profileRepository.NewGormProfileRepository(db)

	// User
	userRepo := userRepository.NewGormUserRepository(db)
	UserService := userUseCase.NewUserService(userRepo, profileRepo)
	userHandler := userHandler.NewHttpUserHandler(UserService, os.Getenv("GOOGLE_OAUTH_CLIENT_ID"), os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"), os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"))

	route.Get("/me", userHandler.GetUser)
}
