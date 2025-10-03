package app

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/MingPV/UserService/internal/entities"
	GrpcProfileHandler "TeamUpSpace-UserService/internal/profile/handler/grpc"
	profileRepository "TeamUpSpace-UserService/internal/profile/repository"
	profileUseCase "TeamUpSpace-UserService/internal/profile/usecase"
	profilepb "TeamUpSpace-UserService/proto/profile"

	GrpcUserReportHandler "TeamUpSpace-UserService/internal/userreport/handler/grpc"
	userreportRepository "TeamUpSpace-UserService/internal/userreport/repository"
	userreportUseCase "TeamUpSpace-UserService/internal/userreport/usecase"
	userreportpb "TeamUpSpace-UserService/proto/userreport"

	GrpcUserFollowHandler "TeamUpSpace-UserService/internal/userfollow/handler/grpc"
	userfollowRepository "TeamUpSpace-UserService/internal/userfollow/repository"
	userfollowUseCase "TeamUpSpace-UserService/internal/userfollow/usecase"
	userfollowpb "TeamUpSpace-UserService/proto/userfollow"

	"github.com/MingPV/UserService/pkg/config"
	"github.com/MingPV/UserService/pkg/database"
	"github.com/MingPV/UserService/pkg/middleware"
	"github.com/MingPV/UserService/pkg/routes"
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

	// UserReport
	userReportRepo := userreportRepository.NewGormUserReportRepository(db)
	userReportService := userreportUseCase.NewUserReportService(userReportRepo)

	// UserFollow
	userFollowRepo := userfollowRepository.NewGormUserFollowRepository(db)
	userFollowService := userfollowUseCase.NewUserFollowService(userFollowRepo)

	// === Register gRPC Services ===
	// Profile
	profileHandler := GrpcProfileHandler.NewGrpcProfileHandler(profileService)
	profilepb.RegisterProfileServiceServer(s, profileHandler)

	// UserReport
	userReportHandler := GrpcUserReportHandler.NewGrpcUserReportHandler(userReportService)
	userreportpb.RegisterUserReportServiceServer(s, userReportHandler)

	// UserFollow
	userFollowHandler := GrpcUserFollowHandler.NewGrpcUserFollowHandler(userFollowService)
	userfollowpb.RegisterUserFollowServiceServer(s, userFollowHandler)

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
		db.Migrator().DropTable(&entities.Profile{}, &entities.User{}, &entities.UserReport{}, &entities.UserFollow{})
	}
	if err := db.AutoMigrate(&entities.Profile{}, &entities.User{}, &entities.UserReport{}, &entities.UserFollow{}); err != nil {
		return nil, nil, err
	}

	return db, cfg, nil
}
