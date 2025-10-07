package app

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/MingPV/UserService/internal/entities"
	GrpcProfileHandler "github.com/MingPV/UserService/internal/profile/handler/grpc"
	profileRepository "github.com/MingPV/UserService/internal/profile/repository"
	profileUseCase "github.com/MingPV/UserService/internal/profile/usecase"
	profilepb "github.com/MingPV/UserService/proto/profile"

	GrpcUserReportHandler "github.com/MingPV/UserService/internal/userreport/handler/grpc"
	userreportRepository "github.com/MingPV/UserService/internal/userreport/repository"
	userreportUseCase "github.com/MingPV/UserService/internal/userreport/usecase"
	userreportpb "github.com/MingPV/UserService/proto/userreport"

	GrpcUserFollowHandler "github.com/MingPV/UserService/internal/userfollow/handler/grpc"
	userfollowRepository "github.com/MingPV/UserService/internal/userfollow/repository"
	userfollowUseCase "github.com/MingPV/UserService/internal/userfollow/usecase"
	userfollowpb "github.com/MingPV/UserService/proto/userfollow"

	"github.com/MingPV/UserService/pkg/config"
	"github.com/MingPV/UserService/pkg/database"
	"github.com/MingPV/UserService/pkg/middleware"
	"github.com/MingPV/UserService/pkg/mq"
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

	// MQ Publisher
	rabbitURL := cfg.RabbitMQUrl
	mqPublisher := mq.NewRabbitMQPublisher(rabbitURL)

	// === Dependency Wiring ===
	// Profile
	profileRepo := profileRepository.NewGormProfileRepository(db)
	profileService := profileUseCase.NewProfileService(profileRepo)

	// UserReport
	userReportRepo := userreportRepository.NewGormUserReportRepository(db)
	userReportService := userreportUseCase.NewUserReportService(userReportRepo)

	// UserFollow
	userFollowRepo := userfollowRepository.NewGormUserFollowRepository(db)
	userFollowService := userfollowUseCase.NewUserFollowService(userFollowRepo, mqPublisher)

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
