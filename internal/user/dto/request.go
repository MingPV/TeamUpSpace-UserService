package dto

import (
	profileDTO "github.com/MingPV/UserService/internal/profile/dto"
)

type RegisterRequest struct {
	Email    string                          `json:"email" validate:"required,email"`
	Username string                          `json:"username" validate:"required,min=3"`
	Password string                          `json:"password" validate:"required,min=6"`
	Profile  profileDTO.CreateProfileRequest `json:"profile"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PatchUserRequest struct {
	IsBan    bool   `json:"is_ban"`
	BanUntil string `json:"ban_until"`
}
