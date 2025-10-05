package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProfileResponse struct {
	UserID        uuid.UUID `gorm:"type:uuid" json:"user_id"`
	DisplayName   string    `json:"display_name"`
	Description   string    `json:"description"`
	Age           int       `json:"age"`
	University    string    `json:"university"`
	Year          int       `json:"year"`
	Major         string    `json:"major"`
	IsGraduated   bool      `json:"is_graduated"`
	ProfileURL    string    `json:"profile_url"`
	BackgroundURL string    `json:"background_url"`
	Location      string    `json:"location"`
	Country       string    `json:"country"`
	Resume        string    `json:"resume"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
