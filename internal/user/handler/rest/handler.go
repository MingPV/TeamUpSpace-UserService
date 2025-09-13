package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/user/dto"
	"github.com/MingPV/UserService/internal/user/usecase"
	"github.com/MingPV/UserService/pkg/apperror"
	"github.com/MingPV/UserService/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type HttpUserHandler struct {
	userUseCase       usecase.UserUseCase
	googleOauthConfig *oauth2.Config
}

func NewHttpUserHandler(useCase usecase.UserUseCase, clientID, clientSecret, redirectURL string) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: useCase, googleOauthConfig: &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}}
}

// Register godoc
// @Summary Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body entities.User true "User registration payload"
// @Success 201 {object} entities.User
// @Router /auth/signup [post]
func (h *HttpUserHandler) Register(c *fiber.Ctx) error {
	req := new(dto.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}

	user_id := uuid.New()
	userEntity := &entities.User{
		ID:       user_id,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}
	profileEntity := &entities.Profile{
		UserID:        user_id,
		DisplayName:   req.Profile.DisplayName,
		Description:   req.Profile.Description,
		Age:           req.Profile.Age,
		University:    req.Profile.University,
		Year:          req.Profile.Year,
		IsGraduated:   req.Profile.IsGraduated,
		ProfileURL:    req.Profile.ProfileURL,
		BackgroundURL: req.Profile.BackgroundURL,
	}

	createdUser, err := h.userUseCase.Register(userEntity, profileEntity)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToUserResponse(createdUser))
}

// Login godoc
// @Summary Authenticate user and return token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Login credentials (email & password)"
// @Success 200 {object} map[string]interface{} "Authenticated user and JWT token"
// @Router /auth/signin [post]
func (h *HttpUserHandler) Login(c *fiber.Ctx) error {
	loginReq := new(dto.LoginRequest)
	if err := c.BodyParser(loginReq); err != nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}

	token, userEntity, err := h.userUseCase.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return responses.ErrorWithMessage(c, apperror.ErrUnauthorized, "invalid email or password")
	}

	// Set JWT token as cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: false,
		Secure:   false, // set to false if not using HTTPS in development
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"user":  dto.ToUserResponse(userEntity),
		"token": token,
	})
}

func (h *HttpUserHandler) Logout(c *fiber.Ctx) error {
	// Invalidate the token on the client side by removing it from storage (e.g., localStorage, cookies)
	// remove cookies
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	// remove refresh token cookies
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return responses.Message(c, fiber.StatusOK, "logged out successfully")
}

// GetUser godoc
// @Summary Get currently authenticated user
// @Tags users
// @Produce json
// @Success 200 {object} entities.User
// @Router /users/me [get]
func (h *HttpUserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return responses.Error(c, apperror.ErrInvalidData)
	}

	userEntity, err := h.userUseCase.FindUserByID(fmt.Sprint(userID))
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToUserResponse(userEntity))
}

// FindUserByID godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entities.User
// @Router /users/{id} [get]
func (h *HttpUserHandler) FindUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return responses.ErrorWithMessage(c, apperror.ErrInvalidData, "id is required")
	}

	userEntity, err := h.userUseCase.FindUserByID(id)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToUserResponse(userEntity))
}

// find user by email
func (h *HttpUserHandler) FindUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return responses.ErrorWithMessage(c, apperror.ErrInvalidData, "email is required")
	}

	userEntity, err := h.userUseCase.FindUserByEmail(email)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToUserResponse(userEntity))
}

// find user by username
func (h *HttpUserHandler) FindUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return responses.ErrorWithMessage(c, apperror.ErrInvalidData, "username is required")
	}

	userEntity, err := h.userUseCase.FindUserByUsername(username)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToUserResponse(userEntity))
}

// FindAllUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} entities.User
// @Router /users [get]
func (h *HttpUserHandler) FindAllUsers(c *fiber.Ctx) error {
	users, err := h.userUseCase.FindAllUsers()
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToUserResponseList(users))
}

// PatchUser godoc
// @Summary Update an user partially
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body entities.User true "User update payload"
// @Success 200 {object} entities.User
// @Router /users/{id} [patch]
func (h *HttpUserHandler) PatchUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.PatchUserRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	user := &entities.User{IsBan: req.IsBan}

	msg, err := validatePatchUser(user)
	if err != nil {
		return responses.ErrorWithMessage(c, err, msg)
	}

	updatedUser, err := h.userUseCase.PatchUser(id, user)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToUserResponse(updatedUser))
}

// DeleteUser godoc
// @Summary Delete an user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.MessageResponse
// @Router /users/{id} [delete]
func (h *HttpUserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.userUseCase.DeleteUser(id); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "user deleted")
}

func validatePatchUser(user *entities.User) (string, error) {

	if user.IsBan {
		// return "username is invalid", apperror.ErrInvalidData
	}

	return "", nil
}

func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// GoogleLogin godoc
// @Summary Redirect to Google OAuth
// @Tags auth
// @Success 302
// @Router /auth/google/login [get]
func (h *HttpUserHandler) GoogleLogin(c *fiber.Ctx) error {

	fmt.Println("Initiating Google OAuth Login")

	state := generateState()
	c.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HTTPOnly: true,
	})

	url := h.googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce, oauth2.SetAuthURLParam("prompt", "select_account"))
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

// GoogleCallback godoc
// @Summary Handle Google OAuth callback
// @Tags auth
// @Success 200 {object} map[string]interface{}
// @Router /auth/google/callback [get]
func (h *HttpUserHandler) GoogleCallback(c *fiber.Ctx) error {
	// check state
	state := c.Query("state")
	cookie := c.Cookies("oauthstate")
	if state != cookie {
		return responses.ErrorWithMessage(c, apperror.ErrUnauthorized, "invalid oauth state")
	}

	code := c.Query("code")
	if code == "" {
		return responses.ErrorWithMessage(c, apperror.ErrInvalidData, "missing code")
	}

	// exchange code -> token
	token, err := h.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return responses.ErrorWithMessage(c, apperror.ErrUnauthorized, "code exchange failed")
	}

	// fetch userinfo
	client := h.googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return responses.ErrorWithMessage(c, err, "failed to fetch userinfo")
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to parse userinfo")
	}

	// send login/register to use case
	jwtToken, _, err := h.userUseCase.LoginOrRegisterWithGoogle(userInfo, token)
	if err != nil {
		return responses.Error(c, err)
	}

	// Set JWT token as cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: false,
		Secure:   false, // set to false if not using HTTPS in development
		SameSite: "Lax",
	})

	frontendURL := os.Getenv("FRONTEND_URL")
	return c.Redirect(frontendURL, http.StatusSeeOther)

	// return c.JSON(fiber.Map{
	// 	"user":  userEntity,
	// 	"token": jwtToken,
	// })
}
