package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Profile
	profileHandler "github.com/MingPV/UserService/internal/profile/handler/rest"
	profileRepository "github.com/MingPV/UserService/internal/profile/repository"
	profileUseCase "github.com/MingPV/UserService/internal/profile/usecase"

	// User
	userHandler "github.com/MingPV/UserService/internal/user/handler/rest"
	userRepository "github.com/MingPV/UserService/internal/user/repository"
	userUseCase "github.com/MingPV/UserService/internal/user/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	api := app.Group("/api/v1")

	// === Dependency Wiring ===

	// Profile
	profileRepo := profileRepository.NewGormProfileRepository(db)
	profileService := profileUseCase.NewProfileService(profileRepo)
	profileHandler := profileHandler.NewHttpProfileHandler(profileService)

	// User
	userRepo := userRepository.NewGormUserRepository(db)
	UserService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(UserService)

	// === Public Routes ===

	// Auth routes (separated from /users)
	authGroup := api.Group("/auth")
	authGroup.Post("/signup", userHandler.Register)
	authGroup.Post("/signin", userHandler.Login)

	// User routes
	userGroup := api.Group("/users")
	userGroup.Get("/", userHandler.FindAllUsers)
	userGroup.Get("/:id", userHandler.FindUserByID)
	userGroup.Patch("/:id", userHandler.PatchUser)
	userGroup.Delete("/:id", userHandler.DeleteUser)

	// Profile routes
	profileGroup := api.Group("/profiles")
	profileGroup.Get("/", profileHandler.FindAllProfiles)
	profileGroup.Get("/:id", profileHandler.FindProfileByID)
	profileGroup.Post("/", profileHandler.CreateProfile)
	profileGroup.Patch("/:id", profileHandler.PatchProfile)
	profileGroup.Delete("/:id", profileHandler.DeleteProfile)
}
