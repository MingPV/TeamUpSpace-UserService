package rest

import (
	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/profile/dto"
	"github.com/MingPV/UserService/internal/profile/usecase"
	responses "github.com/MingPV/UserService/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpProfileHandler struct {
	profileUseCase usecase.ProfileUseCase
}

func NewHttpProfileHandler(useCase usecase.ProfileUseCase) *HttpProfileHandler {
	return &HttpProfileHandler{profileUseCase: useCase}
}

// CreateProfile godoc
// @Summary Create a new profile
// @Tags profiles
// @Accept json
// @Produce json
// @Param profile body entities.Profile true "Profile payload"
// @Success 201 {object} entities.Profile
// @Router /profiles [post]
func (h *HttpProfileHandler) CreateProfile(c *fiber.Ctx) error {
	var req dto.CreateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	profile := &entities.Profile{
		UserID:        req.UserID,
		DisplayName:   req.DisplayName,
		Description:   req.Description,
		Age:           req.Age,
		University:    req.University,
		Year:          req.Year,
		IsGraduated:   req.IsGraduated,
		ProfileURL:    req.ProfileURL,
		BackgroundURL: req.BackgroundURL,
	}

	if err := h.profileUseCase.CreateProfile(profile); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToProfileResponse(profile))
}

// FindAllProfiles godoc
// @Summary Get all profiles
// @Tags profiles
// @Produce json
// @Success 200 {array} entities.Profile
// @Router /profiles [get]
func (h *HttpProfileHandler) FindAllProfiles(c *fiber.Ctx) error {
	profiles, err := h.profileUseCase.FindAllProfiles()
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToProfileResponseList(profiles))
}

// FindProfileByID godoc
// @Summary Get profile by ID
// @Tags profiles
// @Produce json
// @Param id path int true "Profile ID"
// @Success 200 {object} entities.Profile
// @Router /profiles/{id} [get]
func (h *HttpProfileHandler) FindProfileByID(c *fiber.Ctx) error {
	user_id := c.Params("id")

	profile, err := h.profileUseCase.FindProfileByID(user_id)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToProfileResponse(profile))
}

// PatchProfile godoc
// @Summary Update an profile partially
// @Tags profiles
// @Accept json
// @Produce json
// @Param id path int true "Profile ID"
// @Param profile body entities.Profile true "Profile update payload"
// @Success 200 {object} entities.Profile
// @Router /profiles/{id} [patch]
func (h *HttpProfileHandler) PatchProfile(c *fiber.Ctx) error {
	user_id := c.Params("id")

	var req dto.CreateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}
	profile := &entities.Profile{
		UserID:        req.UserID,
		DisplayName:   req.DisplayName,
		Description:   req.Description,
		Age:           req.Age,
		University:    req.University,
		Year:          req.Year,
		IsGraduated:   req.IsGraduated,
		ProfileURL:    req.ProfileURL,
		BackgroundURL: req.BackgroundURL,
	}

	msg, err := validatePatchProfile(profile)
	if err != nil {
		return responses.ErrorWithMessage(c, err, msg)
	}

	updatedProfile, err := h.profileUseCase.PatchProfile(user_id, profile)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToProfileResponse(updatedProfile))
}

// DeleteProfile godoc
// @Summary Delete an profile by ID
// @Tags profiles
// @Produce json
// @Param id path int true "Profile ID"
// @Success 200 {object} response.MessageResponse
// @Router /profiles/{id} [delete]
func (h *HttpProfileHandler) DeleteProfile(c *fiber.Ctx) error {
	user_id := c.Params("id")

	if err := h.profileUseCase.DeleteProfile(user_id); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "profile deleted")
}

func validatePatchProfile(profile *entities.Profile) (string, error) {

	if profile.Year <= 0 {
		// return "year must be positive", apperror.ErrInvalidData
	}

	return "", nil
}
