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
	"github.com/adamryman/ambition-model/model"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() Service {
	return ambitionService{}
}

type ambitionService struct{}

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

// ReadAction implements Service.
func (s ambitionService) ReadAction(ctx context.Context, in *pb.Action) (*pb.ActionResponse, error) {
	return model.ReadAction(in)
}

// CreateAction implements Service.
func (s ambitionService) CreateAction(ctx context.Context, in *pb.Action) (*pb.ActionResponse, error) {
	return model.CreateAction(in)
}

// ReadOccurrences implements Service.
func (s ambitionService) ReadOccurrences(ctx context.Context, in *pb.Occurrence) (*pb.OccurrenceResponse, error) {
	_ = ctx
	_ = in
	response := pb.OccurrenceResponse{
	// Occurrence:
	// Error:
	}
	return &response, nil
}

// CreateOccurrence implements Service.
func (s ambitionService) CreateOccurrence(ctx context.Context, in *pb.Occurrence) (*pb.OccurrenceResponse, error) {
	var o pb.CreateOccurrenceRequest
	o.ActionId = in.ActionId
	o.Datetime = in.Datetime

	return model.CreateOccurrence(&o)
}

type Service interface {
	ReadActions(ctx context.Context, in *pb.ReadActionsRequest) (*pb.ActionsResponse, error)
	ReadAction(ctx context.Context, in *pb.Action) (*pb.ActionResponse, error)
	CreateAction(ctx context.Context, in *pb.Action) (*pb.ActionResponse, error)
	ReadOccurrences(ctx context.Context, in *pb.Occurrence) (*pb.OccurrenceResponse, error)
	CreateOccurrence(ctx context.Context, in *pb.Occurrence) (*pb.OccurrenceResponse, error)
}
