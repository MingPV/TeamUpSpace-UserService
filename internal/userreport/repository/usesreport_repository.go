package repository

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/google/uuid"
)

type UserReportRepository interface {
	Save(report *entities.UserReport) error
	FindByID(id int) (*entities.UserReport, error)
	FindAllByReporter(reporter uuid.UUID) ([]*entities.UserReport, error)
	FindAllByReportTo(reportTo uuid.UUID) ([]*entities.UserReport, error)
	FindAll() ([]*entities.UserReport, error)
	Patch(id int, update *entities.UserReport) error
	Delete(id int) error
}