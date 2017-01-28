package handler

import (
	"golang.org/x/net/context"

	pb "github.com/adamryman/ambition-model/ambition-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AmbitionServer {
	return ambitionService{}
}

type ambitionService struct{}

// CreateAction implements Service.
func (s ambitionService) CreateAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	var resp pb.Action
	resp = pb.Action{
	// ID:
	// UserID:
	// ActionName:
	}
	return &resp, nil
}

// CreateOccurrence implements Service.
func (s ambitionService) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.Occurrence, error) {
	var resp pb.Occurrence
	resp = pb.Occurrence{
	// ID:
	// ActionID:
	// Datetime:
	// Data:
	}
	return &resp, nil
}

// ReadActions implements Service.
func (s ambitionService) ReadActions(ctx context.Context, in *pb.User) (*pb.ActionsResponse, error) {
	var resp pb.ActionsResponse
	resp = pb.ActionsResponse{
	// Actions:
	}
	return &resp, nil
}

// ReadAction implements Service.
func (s ambitionService) ReadAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	var resp pb.Action
	resp = pb.Action{
	// ID:
	// UserID:
	// ActionName:
	}
	return &resp, nil
}
