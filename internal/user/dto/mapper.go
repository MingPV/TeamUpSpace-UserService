package dto

import (
	"github.com/MingPV/UserService/internal/entities"

	profileDTO "github.com/MingPV/UserService/internal/profile/dto"
)

// From entity.User to UserResponse
func ToUserResponse(user *entities.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		IsBan:    user.IsBan,
		BanUntil: user.BanUntil,
		Profile:  *profileDTO.ToProfileResponse(&user.Profile),
	}
}

func ToUserResponseList(users []*entities.User) []*UserResponse {
	responses := make([]*UserResponse, len(users))
	for i, u := range users {
		responses[i] = ToUserResponse(u)
	}
	return responses
}
