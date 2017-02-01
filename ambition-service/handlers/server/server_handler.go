package handler

import (
	"fmt"
	"golang.org/x/net/context"
	"time"

	"github.com/pkg/errors"

	pb "github.com/adamryman/ambition-model/ambition-service"
	mysql "github.com/adamryman/ambition-model/mysql"
	"github.com/adamryman/kit/dbconn"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AmbitionServer {
	database, err := mysql.Open(dbconn.FromENV("MYSQL").MySQL())
	if err != nil {
		// TODO: Do not panic, start something to try connection over and over.
		// Maybe 100 times?
		// DEBUG_SVC=1 then do like 3.
		// There will also need to be retry logic for the database methods
		panic(err)
	}
	return ambitionService{
		db: database,
	}
}

type ambitionService struct {
	db pb.Database
}

// CreateAction implements Service.
func (s ambitionService) CreateAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	// TODO: Input validation
	a, err := s.db.CreateAction(in)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create action")
	}
	return a, nil
}

// CreateOccurrence implements Service.
func (s ambitionService) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.Occurrence, error) {
	// TODO: Make sure database accepts this time format
	now := time.Now().String()

	occurrence := in.GetOccurrence()
	if occurrence == nil {
		return nil, errors.New("cannot create nil occurrence")
	}
	if occurrence.GetDatetime() == "" {
		occurrence.Datetime = now
	}

	action, err := s.db.ReadActionByID(occurrence.GetActionID())
	if err != nil {
		return nil, errors.Wrap(err, "cannot read action")
	}
	if action.GetUserID() != in.GetUserID() {
		// TODO: Replace this "logging" with real logging
		fmt.Printf("action-user-id: %d\nuser-id:%d\n", action.GetUserID(), in.GetUserID())
		return nil, errors.New("cannot create occurrence for action not owned by user")
	}

	o, err := s.db.CreateOccurrence(occurrence)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create occurrence")
	}
	return o, nil
}

// ReadAction implements Service.
func (s ambitionService) ReadAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	if in.GetID() != 0 {
		a, err := s.db.ReadActionByID(in.GetID())
		if err != nil {

			return nil, errors.Wrap(err, "cannot read action")
		}
		return a, nil
	}
	if name, userID := in.GetName(), in.GetUserID(); name != "" && userID != 0 {
		a, err := s.db.ReadActionByNameAndUserID(name, userID)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read action")
		}
		return a, nil
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
