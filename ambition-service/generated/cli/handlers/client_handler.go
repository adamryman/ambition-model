package clienthandler

import (
	pb "github.com/adamryman/ambition-model/ambition-service"
)

// CreateAction implements Service.
func CreateAction(IDCreateAction int64, NameCreateAction string, UserIDCreateAction int64) (*pb.Action, error) {
	request := pb.Action{
		ID:     IDCreateAction,
		Name:   NameCreateAction,
		UserID: UserIDCreateAction,
	}
	return &request, nil
}

// CreateOccurrence implements Service.
func CreateOccurrence(UserIDCreateOccurrence int64, OccurrenceCreateOccurrence pb.Occurrence) (*pb.CreateOccurrenceRequest, error) {
	request := pb.CreateOccurrenceRequest{
		UserID:     UserIDCreateOccurrence,
		Occurrence: &OccurrenceCreateOccurrence,
	}
	return &request, nil
}

// ReadAction implements Service.
func ReadAction(IDReadAction int64, NameReadAction string, UserIDReadAction int64) (*pb.Action, error) {
	request := pb.Action{
		ID:     IDReadAction,
		Name:   NameReadAction,
		UserID: UserIDReadAction,
	}
	return &request, nil
}

// ReadActions implements Service.
func ReadActions(UserIDReadActions int64) (*pb.User, error) {
	request := pb.User{
		UserID: UserIDReadActions,
	}
	return &request, nil
}

// ReadOccurrences implements Service.
func ReadOccurrences(IDReadOccurrences int64, NameReadOccurrences string, UserIDReadOccurrences int64) (*pb.Action, error) {
	request := pb.Action{
		ID:     IDReadOccurrences,
		Name:   NameReadOccurrences,
		UserID: UserIDReadOccurrences,
	}
	return &request, nil
}
