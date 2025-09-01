package repository

import (
	"github.com/MingPV/UserService/internal/entities"
	"gorm.io/gorm"
)

type GormProfileRepository struct {
	db *gorm.DB
}

func NewGormProfileRepository(db *gorm.DB) ProfileRepository {
	return &GormProfileRepository{db: db}
}

func (r *GormProfileRepository) Save(profile *entities.Profile) error {
	return r.db.Create(&profile).Error
}

func (r *GormProfileRepository) FindAll() ([]*entities.Profile, error) {
	var profileValues []entities.Profile
	if err := r.db.Find(&profileValues).Error; err != nil {
		return nil, err
	}

	profiles := make([]*entities.Profile, len(profileValues))
	for i := range profileValues {
		profiles[i] = &profileValues[i]
	}
	return profiles, nil
}

func (r *GormProfileRepository) FindByID(user_id string) (*entities.Profile, error) {
	var profile entities.Profile
	if err := r.db.First(&profile, "user_id = ?", user_id).Error; err != nil {
		return &entities.Profile{}, err
	}
	return &profile, nil
}

func (r *GormProfileRepository) Patch(user_id string, profile *entities.Profile) error {
	result := r.db.Model(&entities.Profile{}).Where("user_id = ?", user_id).Updates(profile)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormProfileRepository) Delete(user_id string) error {
	result := r.db.Delete(&entities.Profile{}, "user_id = ?", user_id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
