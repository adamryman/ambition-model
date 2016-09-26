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
func (s ambitionService) CreateAction(ctx context.Context, in *pb.CreateActionRequest) (*pb.CreateActionResponse, error) {
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
func (s ambitionService) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.CreateOccurrenceResponse, error) {
	_ = ctx
	_ = in

	var err error
	if db == nil {
		db, err = database.New()
		if err != nil {
			return nil, err
		}
	}

	occurrence, err := db.InsertOccurrence(in)

	if err != nil {
		return occurrence, err
	}

	return occurrence, nil
}

type Service interface {
	CreateAction(ctx context.Context, in *pb.CreateActionRequest) (*pb.CreateActionResponse, error)
	CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.CreateOccurrenceResponse, error)
}
