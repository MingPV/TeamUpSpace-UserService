package grpc

import (
	"context"

	"github.com/MingPV/UserService/internal/entities"
	"github.com/MingPV/UserService/internal/userreport/usecase"
	"github.com/MingPV/UserService/pkg/apperror"
	userreportpb "github.com/MingPV/UserService/proto/userreport"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)

type GrpcUserReportHandler struct {
	userReportUseCase usecase.UserReportUseCase
	userreportpb.UnimplementedUserReportServiceServer
}

func NewGrpcUserReportHandler(uc usecase.UserReportUseCase) *GrpcUserReportHandler {
	return &GrpcUserReportHandler{userReportUseCase: uc}
}

func (h *GrpcUserReportHandler) CreateUserReport(ctx context.Context, req *userreportpb.CreateUserReportRequest) (*userreportpb.CreateUserReportResponse, error) {
	reporter, err := uuid.Parse(req.Reporter)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid reporter UUID")
	}
	reportTo, err := uuid.Parse(req.ReportTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid report_to UUID")
	}

	ur := &entities.UserReport{
		Reporter: reporter,
		ReportTo: reportTo,
		Detail:   req.Detail,
		Status:   req.Status,
	}

	created, err := h.userReportUseCase.CreateUserReport(ur)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &userreportpb.CreateUserReportResponse{
		UserReport: toProtoUserReport(created),
	}, nil
}

func (h *GrpcUserReportHandler) FindUserReportByID(ctx context.Context, req *userreportpb.FindUserReportByIDRequest) (*userreportpb.FindUserReportByIDResponse, error) {
	ur, err := h.userReportUseCase.FindUserReportByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &userreportpb.FindUserReportByIDResponse{
		UserReport: toProtoUserReport(ur),
	}, nil
}

func (h *GrpcUserReportHandler) FindAllByReporter(ctx context.Context, req *userreportpb.FindAllByReporterRequest) (*userreportpb.FindAllByReporterResponse, error) {
	reports, err := h.userReportUseCase.FindAllByReporter(req.Reporter)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoReports []*userreportpb.UserReport
	for _, r := range reports {
		protoReports = append(protoReports, toProtoUserReport(r))
	}

	return &userreportpb.FindAllByReporterResponse{
		UserReports: protoReports,
	}, nil
}

func (h *GrpcUserReportHandler) FindAllByReportTo(ctx context.Context, req *userreportpb.FindAllByReportToRequest) (*userreportpb.FindAllByReportToResponse, error) {
	reports, err := h.userReportUseCase.FindAllByReportTo(req.ReportTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoReports []*userreportpb.UserReport
	for _, r := range reports {
		protoReports = append(protoReports, toProtoUserReport(r))
	}

	return &userreportpb.FindAllByReportToResponse{
		UserReports: protoReports,
	}, nil
}

func (h *GrpcUserReportHandler) FindAllUserReports(ctx context.Context, req *userreportpb.FindAllUserReportsRequest) (*userreportpb.FindAllUserReportsResponse, error) {
	reports, err := h.userReportUseCase.FindAllUserReports()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoReports []*userreportpb.UserReport
	for _, r := range reports {
		protoReports = append(protoReports, toProtoUserReport(r))
	}

	return &userreportpb.FindAllUserReportsResponse{
		UserReports: protoReports,
	}, nil
}

func (h *GrpcUserReportHandler) PatchUserReport(ctx context.Context, req *userreportpb.PatchUserReportRequest) (*userreportpb.PatchUserReportResponse, error) {
	update := &entities.UserReport{}

	if req.Detail != nil {
		update.Detail = req.GetDetail()
	}
	if req.Status != nil {
		update.Status = req.GetStatus()
	}

	updated, err := h.userReportUseCase.PatchUserReport(int(req.Id), update)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &userreportpb.PatchUserReportResponse{
		UserReport: toProtoUserReport(updated),
	}, nil
}

func (h *GrpcUserReportHandler) DeleteUserReport(ctx context.Context, req *userreportpb.DeleteUserReportRequest) (*userreportpb.DeleteUserReportResponse, error) {
	if err := h.userReportUseCase.DeleteUserReport(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &userreportpb.DeleteUserReportResponse{
		Message: "user report deleted",
	}, nil
}

func toProtoUserReport(ur *entities.UserReport) *userreportpb.UserReport {
	return &userreportpb.UserReport{
		Id:        int32(ur.ID),
		Reporter:  ur.Reporter.String(),
		ReportTo:  ur.ReportTo.String(),
		Detail:    ur.Detail,
		Status:    ur.Status,
		CreatedAt: ur.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
