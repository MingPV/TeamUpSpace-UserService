package dto

import (
	"github.com/MingPV/UserService/internal/entities"
)

func ToProfileResponse(profile *entities.Profile) *ProfileResponse {
	return &ProfileResponse{
		UserID:        profile.UserID,
		DisplayName:   profile.DisplayName,
		Description:   profile.Description,
		Age:           profile.Age,
		University:    profile.University,
		Year:          profile.Year,
		IsGraduated:   profile.IsGraduated,
		ProfileURL:    profile.ProfileURL,
		BackgroundURL: profile.BackgroundURL,
		CreatedAt:     profile.CreatedAt,
		UpdatedAt:     profile.UpdatedAt,
	}
}

func ToProfileResponseList(profiles []*entities.Profile) []*ProfileResponse {
	result := make([]*ProfileResponse, 0, len(profiles))
	for _, o := range profiles {
		result = append(result, ToProfileResponse(o))
	}
	return result
}
