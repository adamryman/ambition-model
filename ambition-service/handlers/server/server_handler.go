package handler

import (
	"golang.org/x/net/context"
	"os"

	"github.com/pkg/errors"

	pb "github.com/adamryman/ambition-model/ambition-service"
	sqlite "github.com/adamryman/ambition-model/sqlite"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AmbitionServer {
	dbLocation := os.Getenv("SQLITE")

	database, err := sqlite.InitDatabase(dbLocation)
	if err != nil {
		panic(err)
	}
	return ambitionService{
		db: database,
	}
}

type ambitionService struct {
	db AmbitionDB
}

type AmbitionDB interface {
	CreateAction(*pb.Action) (*pb.Action, error)
	CreateOccurrence(*pb.CreateOccurrenceRequest) (*pb.Occurrence, error)
	ReadActionByID(*pb.Action) (*pb.Action, error)
	ReadActionByUserIdAndName(*pb.Action) (*pb.Action, error)
	// TODO: Add queries for all rpc's
}

// CreateAction implements Service.
func (s ambitionService) CreateAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	// TODO: Input validation
	return s.db.CreateAction(in)
}

// CreateOccurrence implements Service.
func (s ambitionService) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.Occurrence, error) {
	// TODO: Input validation
	if in.GetOccurrence() == nil || in.UserID == 0 {
		return nil, errors.New("cannot create occurrence, UserID or Occurrence is nil")
	}
	return s.db.CreateOccurrence(in)
}

// ReadAction implements Service.
func (s ambitionService) ReadAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	if in.GetID() != 0 {
		return s.db.ReadActionByID(in)
	}
	if in.GetUserID() == 0 || in.GetName() == "" {
		return s.db.ReadActionByUserIdAndName(in)
	}
	return nil, errors.New("cannot read action, need ID or BOTH UserID and Name")
}

// ReadActions implements Service.
// TODO: Implement
func (s ambitionService) ReadActions(ctx context.Context, in *pb.User) (*pb.ActionsResponse, error) {
	// TODO: Input validation
	var resp pb.ActionsResponse
	resp = pb.ActionsResponse{
	// Actions:
	}
	return &resp, nil
}

// ReadOccurrences implements Service.
// TODO: Implement
func (s ambitionService) ReadOccurrences(ctx context.Context, in *pb.Action) (*pb.OccurrencesResponse, error) {
	// TODO: Input validation
	var resp pb.OccurrencesResponse
	resp = pb.OccurrencesResponse{
	// Occurrences:
	}
	return &resp, nil
}
