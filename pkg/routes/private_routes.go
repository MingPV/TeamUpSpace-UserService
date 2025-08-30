package routes

import (
	userHandler "github.com/MingPV/UserService/internal/user/handler/rest"
	userRepository "github.com/MingPV/UserService/internal/user/repository"
	userUseCase "github.com/MingPV/UserService/internal/user/usecase"
	middleware "github.com/MingPV/UserService/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app fiber.Router, db *gorm.DB) {

	route := app.Group("/api/v1", middleware.JWTMiddleware())

	userRepo := userRepository.NewGormUserRepository(db)
	UserService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(UserService)

	route.Get("/me", userHandler.GetUser)

}
