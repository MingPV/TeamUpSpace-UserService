package usecase

import (
	"github.com/MingPV/UserService/internal/entities"
)

type UserReportUseCase interface {
	CreateUserReport(report *entities.UserReport) (*entities.UserReport, error)
	FindUserReportByID(id int) (*entities.UserReport, error)
	FindAllByReporter(reporter string) ([]*entities.UserReport, error)
	FindAllByReportTo(reportTo string) ([]*entities.UserReport, error)
	FindAllUserReports() ([]*entities.UserReport, error)
	PatchUserReport(id int, updatedFields *entities.UserReport) (*entities.UserReport, error)
	DeleteUserReport(id int) error
}