package entities

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	UserID        uuid.UUID `gorm:"primaryKey;autoIncrement" json:"user_id"`
	DisplayName   string    `json:"display_name"`
	Description   string    `json:"description"`
	Age           int       `json:"age"`
	University    string    `json:"university"`
	Year          int       `json:"year"`
	IsGraduated   bool      `json:"is_graduated"`
	ProfileURL    string    `json:"profile_url"`
	BackgroundURL string    `json:"background_url"`
	Location      string    `json:"location"`
	Country       string    `json:"country"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
