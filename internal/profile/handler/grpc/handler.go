package grpc

import (
	"context"

	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/profile/usecase"
	"github.com/MingPV/UserService/pkg/apperror"
	profilepb "github.com/MingPV/UserService/proto/profile"

	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcProfileHandler struct {
	profileUseCase usecase.ProfileUseCase
	profilepb.UnimplementedProfileServiceServer
}

func NewGrpcProfileHandler(uc usecase.ProfileUseCase) *GrpcProfileHandler {
	return &GrpcProfileHandler{profileUseCase: uc}
}

func (h *GrpcProfileHandler) CreateProfile(ctx context.Context, req *profilepb.CreateProfileRequest) (*profilepb.CreateProfileResponse, error) {

	userUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid user id: %s", err.Error())
	}
	profile := &entities.Profile{
		UserID:        userUUID,
		DisplayName:   req.DisplayName,
		Description:   req.Description,
		Age:           int(req.Age),
		University:    req.University,
		Year:          int(req.Year),
		Major:         req.Major,
		IsGraduated:   req.IsGraduated,
		ProfileURL:    req.ProfileUrl,
		BackgroundURL: req.BackgroundUrl,
		Location:      req.Location,
		Country:       req.Country,
	}

	if err := h.profileUseCase.CreateProfile(profile); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &profilepb.CreateProfileResponse{Profile: toProtoProfile(profile)}, nil
}

func (h *GrpcProfileHandler) FindProfileByID(ctx context.Context, req *profilepb.FindProfileByIDRequest) (*profilepb.FindProfileByIDResponse, error) {
	profile, err := h.profileUseCase.FindProfileByID(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &profilepb.FindProfileByIDResponse{Profile: toProtoProfile(profile)}, nil
}

func (h *GrpcProfileHandler) FindAllProfiles(ctx context.Context, req *profilepb.FindAllProfilesRequest) (*profilepb.FindAllProfilesResponse, error) {
	profiles, err := h.profileUseCase.FindAllProfiles()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoProfiles []*profilepb.Profile
	for _, o := range profiles {
		protoProfiles = append(protoProfiles, toProtoProfile(o))
	}

	return &profilepb.FindAllProfilesResponse{Profiles: protoProfiles}, nil
}

func (h *GrpcProfileHandler) PatchProfile(ctx context.Context, req *profilepb.PatchProfileRequest) (*profilepb.PatchProfileResponse, error) {
	userUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid user id: %s", err.Error())
	}

	profile := &entities.Profile{
		UserID:        userUUID,
		DisplayName:   req.DisplayName,
		Description:   req.Description,
		Age:           int(req.Age),
		University:    req.University,
		Year:          int(req.Year),
		Major:         req.Major,
		IsGraduated:   req.IsGraduated,
		ProfileURL:    req.ProfileUrl,
		BackgroundURL: req.BackgroundUrl,
		Location:      req.Location,
		Country:       req.Country,
	}

	updatedProfile, err := h.profileUseCase.PatchProfile(req.UserId, profile)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &profilepb.PatchProfileResponse{Profile: toProtoProfile(updatedProfile)}, nil
}

func (h *GrpcProfileHandler) DeleteProfile(ctx context.Context, req *profilepb.DeleteProfileRequest) (*profilepb.DeleteProfileResponse, error) {
	if err := h.profileUseCase.DeleteProfile(req.UserId); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &profilepb.DeleteProfileResponse{Message: "profile deleted"}, nil
}

// helper function convert entities.Profile to profilepb.Profile
func toProtoProfile(o *entities.Profile) *profilepb.Profile {
	if o == nil {
		return nil
	}

	return &profilepb.Profile{
		UserId:        o.UserID.String(),
		DisplayName:   o.DisplayName,
		Description:   o.Description,
		Age:           int32(o.Age),
		University:    o.University,
		Year:          int32(o.Year),
		Major:         o.Major,
		IsGraduated:   o.IsGraduated,
		ProfileUrl:    o.ProfileURL,
		BackgroundUrl: o.BackgroundURL,
		Location:      o.Location,
		Country:       o.Country,
		CreatedAt:     timestamppb.New(o.CreatedAt),
		UpdatedAt:     timestamppb.New(o.UpdatedAt),
	}
}
