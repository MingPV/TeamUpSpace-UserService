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

func (r *GormUserFollowRepository) Delete(userID, followTo uuid.UUID) error {
	return r.db.Where("user_id = ? AND follow_to = ?", userID, followTo).Delete(&entities.UserFollow{}).Error
}

func (r *GormUserFollowRepository) FindAllFollowers(followTo uuid.UUID) ([]*entities.UserFollow, error) {
	var values []entities.UserFollow
	if err := r.db.Where("follow_to = ?", followTo).Find(&values).Error; err != nil {
		return nil, err
	}
	result := make([]*entities.UserFollow, len(values))
	for i := range values {
		result[i] = &values[i]
	}
	return result, nil
}

func (r *GormUserFollowRepository) FindAllFollowings(userID uuid.UUID) ([]*entities.UserFollow, error) {
	var values []entities.UserFollow
	if err := r.db.Where("user_id = ?", userID).Find(&values).Error; err != nil {
		return nil, err
	}
	result := make([]*entities.UserFollow, len(values))
	for i := range values {
		result[i] = &values[i]
	}
	return result, nil
}