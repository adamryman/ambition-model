package handler

// This file contains the Service definition, and a basic service
// implementation. It also includes service middlewares.

import (
	_ "errors"
	_ "time"

	"golang.org/x/net/context"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/metrics"

	pb "github.com/adamryman/ambition-model/ambition-service"
	// Userland

	"github.com/adamryman/ambition-model/model"
)

// NewService returns a naïve, stateless implementation of Service.

type ambitionService struct{}

func NewService() Service {
	return ambitionService{}
}

// CreateAction implements Service.
func (s ambitionService) CreateAction(ctx context.Context, in *pb.CreateActionRequest) (*pb.ActionResponse, error) {
	_ = ctx
	_ = in

	return model.CreateAction(in)
}

// CreateOccurrence implements Service.
func (s ambitionService) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.OccurrenceResponse, error) {

	return model.CreateOccurrence(in)
}

// ReadAction implements Service.
func (s ambitionService) ReadAction(ctx context.Context, in *pb.ReadActionRequest) (*pb.ActionResponse, error) {
	return model.ReadAction(in)
}

// ReadActions implements Service.
func (s ambitionService) ReadActions(ctx context.Context, in *pb.ReadActionsRequest) (*pb.ActionsResponse, error) {
	_ = ctx
	_ = in
	response := pb.ActionsResponse{
	// Actions:
	// Error:
	}
	return &response, nil
}

// ReadOccurrences implements Service.
func (s ambitionService) ReadOccurrences(ctx context.Context, in *pb.ReadOccurrencesRequest) (*pb.OccurrenceResponse, error) {
	_ = ctx
	_ = in
	response := pb.OccurrenceResponse{
	// Occurrence:
	// Error:
	}
	return &response, nil
}

type Service interface {
	ReadActions(ctx context.Context, in *pb.ReadActionsRequest) (*pb.ActionsResponse, error)
	ReadAction(ctx context.Context, in *pb.ReadActionRequest) (*pb.ActionResponse, error)
	CreateAction(ctx context.Context, in *pb.CreateActionRequest) (*pb.ActionResponse, error)
	ReadOccurrences(ctx context.Context, in *pb.ReadOccurrencesRequest) (*pb.OccurrenceResponse, error)
	CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.OccurrenceResponse, error)
}
