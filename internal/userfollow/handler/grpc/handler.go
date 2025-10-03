package grpc

import (
	"context"

	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/userfollow/usecase"
	userfollowpb "github.com/MingPV/UserService/proto/userfollow"
	"google.golang.org/grpc/status"
)

type GrpcUserFollowHandler struct {
	userfollowpb.UnimplementedUserFollowServiceServer
	usecase usecase.UserFollowUseCase
}

func NewGrpcUserFollowHandler(uc usecase.UserFollowUseCase) *GrpcUserFollowHandler {
	return &GrpcUserFollowHandler{usecase: uc}
}

func (h *GrpcUserFollowHandler) FollowUser(ctx context.Context, req *userfollowpb.FollowUserRequest) (*userfollowpb.FollowUserResponse, error) {
	uf, err := h.usecase.FollowUser(req.UserId, req.FollowTo)
	if err != nil {
		return nil, status.Errorf(500, err.Error())
	}
	return &userfollowpb.FollowUserResponse{
		UserFollow: toProto(uf),
	}, nil
}

func (h *GrpcUserFollowHandler) UnfollowUser(ctx context.Context, req *userfollowpb.UnfollowUserRequest) (*userfollowpb.UnfollowUserResponse, error) {
	if err := h.usecase.UnfollowUser(req.UserId, req.FollowTo); err != nil {
		return nil, status.Errorf(500, err.Error())
	}
	return &userfollowpb.UnfollowUserResponse{Message: "unfollowed successfully"}, nil
}

func (h *GrpcUserFollowHandler) FindAllFollowers(ctx context.Context, req *userfollowpb.FindAllFollowersRequest) (*userfollowpb.FindAllFollowersResponse, error) {
	list, err := h.usecase.FindAllFollowers(req.FollowTo)
	if err != nil {
		return nil, status.Errorf(500, err.Error())
	}
	var result []*userfollowpb.UserFollow
	for _, f := range list {
		result = append(result, toProto(f))
	}
	return &userfollowpb.FindAllFollowersResponse{UserFollows: result}, nil
}

func (h *GrpcUserFollowHandler) FindAllFollowings(ctx context.Context, req *userfollowpb.FindAllFollowingsRequest) (*userfollowpb.FindAllFollowingsResponse, error) {
	list, err := h.usecase.FindAllFollowings(req.UserId)
	if err != nil {
		return nil, status.Errorf(500, err.Error())
	}
	var result []*userfollowpb.UserFollow
	for _, f := range list {
		result = append(result, toProto(f))
	}
	return &userfollowpb.FindAllFollowingsResponse{UserFollows: result}, nil
}

func toProto(uf *entities.UserFollow) *userfollowpb.UserFollow {
	return &userfollowpb.UserFollow{
		UserId:    uf.UserID.String(),
		FollowTo:  uf.FollowTo.String(),
		CreatedAt: uf.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}