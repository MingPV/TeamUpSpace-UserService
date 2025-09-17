package usecase

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/userreport/repository"
	"github.com/google/uuid"
)

type UserReportService struct {
	repo repository.UserReportRepository
}

func NewUserReportService(repo repository.UserReportRepository) UserReportUseCase {
	return &UserReportService{repo: repo}
}

func (s *UserReportService) CreateUserReport(report *entities.UserReport) (*entities.UserReport, error) {
	if err := s.repo.Save(report); err != nil {
		return nil, err
	}
	return report, nil
}

func (s *UserReportService) FindUserReportByID(id int) (*entities.UserReport, error) {
	return s.repo.FindByID(id)
}

func (s *UserReportService) FindAllByReporter(reporter string) ([]*entities.UserReport, error) {
	parsedID, err := uuid.Parse(reporter)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAllByReporter(parsedID)
}

func (s *UserReportService) FindAllByReportTo(reportTo string) ([]*entities.UserReport, error) {
	parsedID, err := uuid.Parse(reportTo)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAllByReportTo(parsedID)
}

func (s *UserReportService) FindAllUserReports() ([]*entities.UserReport, error) {
	return s.repo.FindAll()
}

func (s *UserReportService) PatchUserReport(id int, updatedFields *entities.UserReport) (*entities.UserReport, error) {
	if err := s.repo.Patch(id, updatedFields); err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *UserReportService) DeleteUserReport(id int) error {
	return s.repo.Delete(id)
}