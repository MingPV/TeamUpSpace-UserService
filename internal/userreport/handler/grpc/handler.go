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

func (h *GrpcUserFollowHandler) FollowUser(ctx context.Context, req *userfollowpb.FollowUserRequest) (*userfollowpb.FollowUserResponse, error) {
	uf, err := h.userFollowUseCase.FollowUser(req.UserId, req.FollowTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &userfollowpb.FollowUserResponse{
		UserFollow: toProtoUserFollow(uf),
	}, nil
}

func (h *GrpcUserFollowHandler) UnfollowUser(ctx context.Context, req *userfollowpb.UnfollowUserRequest) (*userfollowpb.UnfollowUserResponse, error) {
	if err := h.userFollowUseCase.UnfollowUser(req.UserId, req.FollowTo); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &userfollowpb.UnfollowUserResponse{
		Message: "unfollowed successfully",
	}, nil
}

func (h *GrpcUserFollowHandler) FindAllFollowers(ctx context.Context, req *userfollowpb.FindAllFollowersRequest) (*userfollowpb.FindAllFollowersResponse, error) {
	follows, err := h.userFollowUseCase.FindAllFollowers(req.FollowTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	var result []*userfollowpb.UserFollow
	for _, f := range follows {
		result = append(result, toProtoUserFollow(f))
	}
	return &userfollowpb.FindAllFollowersResponse{
		UserFollows: result,
	}, nil
}

func (h *GrpcUserFollowHandler) FindAllFollowings(ctx context.Context, req *userfollowpb.FindAllFollowingsRequest) (*userfollowpb.FindAllFollowingsResponse, error) {
	follows, err := h.userFollowUseCase.FindAllFollowings(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	var result []*userfollowpb.UserFollow
	for _, f := range follows {
		result = append(result, toProtoUserFollow(f))
	}
	return &userfollowpb.FindAllFollowingsResponse{
		UserFollows: result,
	}, nil
}

func toProtoUserFollow(uf *entities.UserFollow) *userfollowpb.UserFollow {
	return &userfollowpb.UserFollow{
		UserId:    uf.UserID.String(),
		FollowTo:  uf.FollowTo.String(),
		CreatedAt: uf.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}