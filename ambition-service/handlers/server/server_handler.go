package handler

// This file contains the Service definition, and a basic service
// implementation. It also includes service middlewares.

import (
	_ "errors"
	_ "time"

	"golang.org/x/net/context"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/metrics"

	pb "github.com/adamryman/ambition-truss/ambition-service"
	// Userland

	"github.com/adamryman/ambition-truss/database"
)

// NewService returns a naïve, stateless implementation of Service.

type ambitionService struct{}

func NewService() Service {
	return ambitionService{}
}

var db *database.Database

// CreateAction implements Service.
func (s ambitionService) CreateAction(ctx context.Context, in *pb.CreateActionRequest) (*pb.ActionResponse, error) {
	_ = ctx
	_ = in

	var err error
	if db == nil {
		db, err = database.New()
		if err != nil {
			return nil, err
		}
	}

	action, err := db.InsertAction(in)

	if err != nil {
		return action, err
	}

	return action, nil
}

// CreateOccurrence implements Service.
func (s ambitionService) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.OccurrenceResponse, error) {
	_ = ctx
	_ = in

	//var err error
	//if db == nil {
	//db, err = database.New()
	//if err != nil {
	//return nil, err
	//}
	//}

	//occurrence, err := db.InsertOccurrence(in)

	//if err != nil {
	//return occurrence, err
	//}

	return nil, nil
}

// CreateOccurrenceForReal implements Service.

// Items:
// Error:

// ReadActions implements Service.
func (s ambitionService) ReadActions(ctx context.Context, in *pb.ActionsRequest) (*pb.ActionResponse, error) {
	_ = ctx
	_ = in
	response := pb.ActionResponse{
	// Result:
	// Error:
	}
	return &response, nil
}

// ReadOccurrences implements Service.
func (s ambitionService) ReadOccurrences(ctx context.Context, in *pb.OccurrencesRequest) (*pb.OccurrenceResponse, error) {
	_ = ctx
	_ = in
	response := pb.OccurrenceResponse{
	// Result:
	// Error:
	}
	return &response, nil
}

type Service interface {
	ReadActions(ctx context.Context, in *pb.ActionsRequest) (*pb.ActionResponse, error)
	CreateAction(ctx context.Context, in *pb.CreateActionRequest) (*pb.ActionResponse, error)
	ReadOccurrences(ctx context.Context, in *pb.OccurrencesRequest) (*pb.OccurrenceResponse, error)
	CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.OccurrenceResponse, error)
}
