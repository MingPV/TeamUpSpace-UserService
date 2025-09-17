package repository

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserReportRepository struct {
	db *gorm.DB
}

func NewGormUserReportRepository(db *gorm.DB) UserReportRepository {
	return &GormUserReportRepository{db: db}
}

func (r *GormUserReportRepository) Save(report *entities.UserReport) error {
	return r.db.Create(report).Error
}

func (r *GormUserReportRepository) FindByID(id int) (*entities.UserReport, error) {
	var report entities.UserReport
	if err := r.db.Where("id = ?", id).First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *GormUserReportRepository) FindAllByReporter(reporter uuid.UUID) ([]*entities.UserReport, error) {
	var values []entities.UserReport
	if err := r.db.Where("reporter = ?", reporter).Find(&values).Error; err != nil {
		return nil, err
	}
	reports := make([]*entities.UserReport, len(values))
	for i := range values {
		reports[i] = &values[i]
	}
	return reports, nil
}

func (r *GormUserReportRepository) FindAllByReportTo(reportTo uuid.UUID) ([]*entities.UserReport, error) {
	var values []entities.UserReport
	if err := r.db.Where("report_to = ?", reportTo).Find(&values).Error; err != nil {
		return nil, err
	}
	reports := make([]*entities.UserReport, len(values))
	for i := range values {
		reports[i] = &values[i]
	}
	return reports, nil
}

func (r *GormUserReportRepository) FindAll() ([]*entities.UserReport, error) {
	var values []entities.UserReport
	if err := r.db.Find(&values).Error; err != nil {
		return nil, err
	}
	reports := make([]*entities.UserReport, len(values))
	for i := range values {
		reports[i] = &values[i]
	}
	return reports, nil
}

func (r *GormUserReportRepository) Patch(id int, update *entities.UserReport) error {
	result := r.db.Model(&entities.UserReport{}).
		Where("id = ?", id).
		Updates(update)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormUserReportRepository) Delete(id int) error {
	result := r.db.Where("id = ?", id).Delete(&entities.UserReport{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}