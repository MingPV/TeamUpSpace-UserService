package repository

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserFollowRepository struct {
	db *gorm.DB
}

func NewGormUserFollowRepository(db *gorm.DB) UserFollowRepository {
	return &GormUserFollowRepository{db: db}
}

func (r *GormUserFollowRepository) Save(follow *entities.UserFollow) error {
	return r.db.Create(follow).Error
}

func (r *GormUserFollowRepository) FindByID(id int) (*entities.UserFollow, error) {
	var follow entities.UserFollow
	if err := r.db.Where("id = ?", id).First(&follow).Error; err != nil {
		return nil, err
	}
	return &follow, nil
}

func (r *GormUserFollowRepository) FindAllByUser(userID uuid.UUID) ([]*entities.UserFollow, error) {
	var values []entities.UserFollow
	if err := r.db.Where("user_id = ?", userID).Find(&values).Error; err != nil {
		return nil, err
	}
	follows := make([]*entities.UserFollow, len(values))
	for i := range values {
		follows[i] = &values[i]
	}
	return follows, nil
}

func (r *GormUserFollowRepository) FindAllByFollowTo(followTo uuid.UUID) ([]*entities.UserFollow, error) {
	var values []entities.UserFollow
	if err := r.db.Where("follow_to = ?", followTo).Find(&values).Error; err != nil {
		return nil, err
	}
	follows := make([]*entities.UserFollow, len(values))
	for i := range values {
		follows[i] = &values[i]
	}
	return follows, nil
}

func (r *GormUserFollowRepository) FindAll() ([]*entities.UserFollow, error) {
	var values []entities.UserFollow
	if err := r.db.Find(&values).Error; err != nil {
		return nil, err
	}
	follows := make([]*entities.UserFollow, len(values))
	for i := range values {
		follows[i] = &values[i]
	}
	return follows, nil
}

func (r *GormUserFollowRepository) Delete(id int) error {
	result := r.db.Where("id = ?", id).Delete(&entities.UserFollow{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}