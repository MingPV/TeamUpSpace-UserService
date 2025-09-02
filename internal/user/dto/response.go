package dto

import (
	"time"

	"github.com/google/uuid"

	profileDTO "github.com/MingPV/UserService/internal/profile/dto"
)

type UserResponse struct {
	ID       uuid.UUID                  `json:"id"`
	Email    string                     `json:"email"`
	IsAdmin  bool                       `json:"is_admin"`
	IsBan    bool                       `json:"is_ban"`
	BanUntil time.Time                  `json:"ban_until"`
	Profile  profileDTO.ProfileResponse `json:"profile"`
}
