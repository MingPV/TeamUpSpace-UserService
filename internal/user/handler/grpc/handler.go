package grpc

// import (
// 	"context"

// 	"github.com/MingPV/UserService/internal/entities"
// 	"github.com/MingPV/UserService/internal/user/usecase"
// 	"github.com/MingPV/UserService/pkg/apperror"
// 	userpb "github.com/MingPV/UserService/proto/user"

// 	"github.com/google/uuid"
// 	"google.golang.org/grpc/status"
// 	"google.golang.org/protobuf/types/known/timestamppb"
// )

// type GrpcUserHandler struct {
// 	userUseCase usecase.UserUseCase
// 	userpb.UnimplementedUserServiceServer
// }

// func NewGrpcUserHandler(uc usecase.UserUseCase) *GrpcUserHandler {
// 	return &GrpcUserHandler{userUseCase: uc}
// }

// // ---------- Auth ----------

// func (h *GrpcUserHandler) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
// 	userID := uuid.New()

// 	user := &entities.User{
// 		ID:       userID,
// 		Email:    req.Email,
// 		Username: req.Username,
// 		Password: req.Password,
// 	}

// 	profile := &entities.Profile{
// 		UserID:        userID,
// 		DisplayName:   req.Profile.DisplayName,
// 		Description:   req.Profile.Description,
// 		Age:           int(req.Profile.Age),
// 		University:    req.Profile.University,
// 		Year:          int(req.Profile.Year),
// 		IsGraduated:   req.Profile.IsGraduated,
// 		ProfileURL:    req.Profile.ProfileUrl,
// 		BackgroundURL: req.Profile.BackgroundUrl,
// 	}

// 	createdUser, err := h.userUseCase.Register(user, profile)
// 	if err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
// 	}

// 	return &userpb.RegisterResponse{User: toProtoUser(createdUser)}, nil
// }

// func (h *GrpcUserHandler) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
// 	token, user, err := h.userUseCase.Login(req.Email, req.Password)
// 	if err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "invalid email or password")
// 	}

// 	return &userpb.LoginResponse{
// 		User:  toProtoUser(user),
// 		Token: token,
// 	}, nil
// }

// func (h *GrpcUserHandler) Logout(ctx context.Context, _ *userpb.Empty) (*userpb.LogoutResponse, error) {
// 	// สำหรับ gRPC จะไม่ได้ลบ cookie แต่สามารถ return message ได้
// 	return &userpb.LogoutResponse{Message: "logged out successfully"}, nil
// }

// // ---------- Users ----------

// func (h *GrpcUserHandler) GetUser(ctx context.Context, req *userpb.FindByIDRequest) (*userpb.UserResponse, error) {
// 	user, err := h.userUseCase.FindUserByID(req.Id)
// 	if err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
// 	}
// 	return &userpb.UserResponse{User: toProtoUser(user)}, nil
// }

// func (h *GrpcUserHandler) FindUserByID(ctx context.Context, req *userpb.FindByIDRequest) (*userpb.UserResponse, error) {
// 	user, err := h.userUseCase.FindUserByID(req.Id)
// 	if err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
// 	}
// 	return &userpb.UserResponse{User: toProtoUser(user)}, nil
// }

// func (h *GrpcUserHandler) FindUserByEmail(ctx context.Context, req *userpb.FindByEmailRequest) (*userpb.UserResponse, error) {
// 	user, err := h.userUseCase.FindUserByEmail(req.Email)
// 	if err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
// 	}
// 	return &userpb.UserResponse{User: toProtoUser(user)}, nil
// }

// func (h *GrpcUserHandler) FindUserByUsername(ctx context.Context, req *userpb.FindByUsernameRequest) (*userpb.UserResponse, error) {
// 	user, err := h.userUseCase.FindUserByUsername(req.Username)
// 	if err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
// 	}
// 	return &userpb.UserResponse{User: toProtoUser(user)}, nil
// }

// func (h *GrpcUserHandler) FindAllUsers(ctx context.Context, _ *userpb.Empty) (*userpb.UsersResponse, error) {
// 	users, err := h.userUseCase.FindAllUsers()
// 	if err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
// 	}

// 	var protoUsers []*userpb.User
// 	for _, u := range users {
// 		protoUsers = append(protoUsers, toProtoUser(u))
// 	}
// 	return &userpb.UsersResponse{Users: protoUsers}, nil
// }

// // ---------- Update & Delete ----------

// func (h *GrpcUserHandler) PatchUser(ctx context.Context, req *userpb.PatchUserRequest) (*userpb.PatchUserResponse, error) {
// 	user := &entities.User{IsBan: req.IsBan}

// 	updatedUser, err := h.userUseCase.PatchUser(req.Id, user)
// 	if err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
// 	}

// 	return &userpb.PatchUserResponse{User: toProtoUser(updatedUser)}, nil
// }

// func (h *GrpcUserHandler) DeleteUser(ctx context.Context, req *userpb.FindByIDRequest) (*userpb.DeleteUserResponse, error) {
// 	if err := h.userUseCase.DeleteUser(req.Id); err != nil {
// 		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
// 	}
// 	return &userpb.DeleteUserResponse{Message: "user deleted"}, nil
// }

// // ---------- Helpers ----------

// func toProtoUser(u *entities.User) *userpb.User {
// 	if u == nil {
// 		return nil
// 	}
// 	return &userpb.User{
// 		Id:        u.ID.String(),
// 		Email:     u.Email,
// 		Username:  u.Username,
// 		Password:  u.Password,
// 		IsAdmin:   u.IsAdmin,
// 		IsBan:     u.IsBan,
// 		BanUntil:  timestamppb.New(u.BanUntil),
// 		CreatedAt: timestamppb.New(u.CreatedAt),
// 		UpdatedAt: timestamppb.New(u.UpdatedAt),
// 		Profile: &userpb.Profile{
// 			UserId:        u.Profile.UserID.String(),
// 			DisplayName:   u.Profile.DisplayName,
// 			Description:   u.Profile.Description,
// 			Age:           int32(u.Profile.Age),
// 			University:    u.Profile.University,
// 			Year:          int32(u.Profile.Year),
// 			IsGraduated:   u.Profile.IsGraduated,
// 			ProfileUrl:    u.Profile.ProfileURL,
// 			BackgroundUrl: u.Profile.BackgroundURL,
// 		},
// 	}
// }
