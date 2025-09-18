package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserReport struct {
	ID 		  int 		`gorm:"primaryKey;autoIncrement" json:"id"`
	Reporter  uuid.UUID `gorm:"type:uuid;not null" json:"reporter"`
	ReportTo  uuid.UUID `gorm:"type:uuid;not null" json:"report_to"`
	Detail    string    `gorm:"type:text;not null" json:"detail"`
	Status    string    `gorm:"size:50;default:'pending'" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	ReporterUser *User `gorm:"foreignKey:Reporter;references:ID;constraint:OnDelete:CASCADE" json:"reporter_user"`
	ReportedUser *User `gorm:"foreignKey:ReportTo;references:ID;constraint:OnDelete:CASCADE" json:"reported_user"`
}