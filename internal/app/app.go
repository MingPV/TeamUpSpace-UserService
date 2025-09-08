package app

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/MingPV/UserService/internal/entities"
	GrpcProfileHandler "github.com/MingPV/UserService/internal/profile/handler/grpc"
	profileRepository "github.com/MingPV/UserService/internal/profile/repository"
	profileUseCase "github.com/MingPV/UserService/internal/profile/usecase"
	"github.com/MingPV/UserService/pkg/config"
	"github.com/MingPV/UserService/pkg/database"
	"github.com/MingPV/UserService/pkg/middleware"
	"github.com/MingPV/UserService/pkg/routes"
	profilepb "github.com/MingPV/UserService/proto/profile"
)

// rest
func SetupRestServer(db *gorm.DB, cfg *config.Config) (*fiber.App, error) {
	app := fiber.New()
	middleware.FiberMiddleware(app)
	// comment out Swagger when testing
	// routes.SwaggerRoute(app)
	routes.RegisterPublicRoutes(app, db)
	routes.RegisterPrivateRoutes(app, db)
	routes.RegisterNotFoundRoute(app)
	return app, nil
}

// grpc
func SetupGrpcServer(db *gorm.DB, cfg *config.Config) (*grpc.Server, error) {
	s := grpc.NewServer()

	// === Dependency Wiring ===
	// Profile
	profileRepo := profileRepository.NewGormProfileRepository(db)
	profileService := profileUseCase.NewProfileService(profileRepo)

	// === Register gRPC Services ===
	// Profile
	profileHandler := GrpcProfileHandler.NewGrpcProfileHandler(profileService)
	profilepb.RegisterProfileServiceServer(s, profileHandler)
	return s, nil
}

// dependencies
func SetupDependencies(env string) (*gorm.DB, *config.Config, error) {
	cfg := config.LoadConfig(env)

	db, err := database.Connect(cfg.DatabaseDSN)
	if err != nil {
		return nil, nil, err
	}

	if env == "test" {
		db.Migrator().DropTable(&entities.Profile{}, &entities.User{})
	}
	if err := db.AutoMigrate(&entities.Profile{}, &entities.User{}); err != nil {
		return nil, nil, err
	}

	return db, cfg, nil
}
