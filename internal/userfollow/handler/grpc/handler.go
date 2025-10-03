package grpc

import (
	"context"

	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/userfollow/usecase"
	"github.com/MingPV/UserService/pkg/apperror"
	userfollowpb "github.com/MingPV/UserService/proto/userfollow"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)

type GrpcUserFollowHandler struct {
	userFollowUseCase usecase.UserFollowUseCase
	userfollowpb.UnimplementedUserFollowServiceServer
}

func NewGrpcUserFollowHandler(uc usecase.UserFollowUseCase) *GrpcUserFollowHandler {
	return &GrpcUserFollowHandler{userFollowUseCase: uc}
}

func (h *GrpcUserFollowHandler) CreateUserFollow(ctx context.Context, req *userfollowpb.CreateUserFollowRequest) (*userfollowpb.CreateUserFollowResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid user_id UUID")
	}
	followTo, err := uuid.Parse(req.FollowTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid follow_to UUID")
	}

	uf := &entities.UserFollow{
		UserID:   userID,
		FollowTo: followTo,
	}

	created, err := h.userFollowUseCase.CreateUserFollow(uf)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &userfollowpb.CreateUserFollowResponse{
		UserFollow: toProtoUserFollow(created),
	}, nil
}

func (h *GrpcUserFollowHandler) FindUserFollowByID(ctx context.Context, req *userfollowpb.FindUserFollowByIDRequest) (*userfollowpb.FindUserFollowByIDResponse, error) {
	uf, err := h.userFollowUseCase.FindUserFollowByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &userfollowpb.FindUserFollowByIDResponse{
		UserFollow: toProtoUserFollow(uf),
	}, nil
}

func (h *GrpcUserFollowHandler) FindAllByUser(ctx context.Context, req *userfollowpb.FindAllByUserRequest) (*userfollowpb.FindAllByUserResponse, error) {
	follows, err := h.userFollowUseCase.FindAllByUser(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoFollows []*userfollowpb.UserFollow
	for _, f := range follows {
		protoFollows = append(protoFollows, toProtoUserFollow(f))
	}

	return &userfollowpb.FindAllByUserResponse{
		UserFollows: protoFollows,
	}, nil
}

func (h *GrpcUserFollowHandler) FindAllByFollowTo(ctx context.Context, req *userfollowpb.FindAllByFollowToRequest) (*userfollowpb.FindAllByFollowToResponse, error) {
	follows, err := h.userFollowUseCase.FindAllByFollowTo(req.FollowTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoFollows []*userfollowpb.UserFollow
	for _, f := range follows {
		protoFollows = append(protoFollows, toProtoUserFollow(f))
	}

	return &userfollowpb.FindAllByFollowToResponse{
		UserFollows: protoFollows,
	}, nil
}

func (h *GrpcUserFollowHandler) FindAllUserFollows(ctx context.Context, req *userfollowpb.FindAllUserFollowsRequest) (*userfollowpb.FindAllUserFollowsResponse, error) {
	follows, err := h.userFollowUseCase.FindAllUserFollows()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoFollows []*userfollowpb.UserFollow
	for _, f := range follows {
		protoFollows = append(protoFollows, toProtoUserFollow(f))
	}

	return &userfollowpb.FindAllUserFollowsResponse{
		UserFollows: protoFollows,
	}, nil
}

func (h *GrpcUserFollowHandler) DeleteUserFollow(ctx context.Context, req *userfollowpb.DeleteUserFollowRequest) (*userfollowpb.DeleteUserFollowResponse, error) {
	if err := h.userFollowUseCase.DeleteUserFollow(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &userfollowpb.DeleteUserFollowResponse{
		Message: "user follow deleted",
	}, nil
}

func toProtoUserFollow(uf *entities.UserFollow) *userfollowpb.UserFollow {
	return &userfollowpb.UserFollow{
		Id:        int32(uf.ID),
		UserId:    uf.UserID.String(),
		FollowTo:  uf.FollowTo.String(),
		CreatedAt: uf.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}