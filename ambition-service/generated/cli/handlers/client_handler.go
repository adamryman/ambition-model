package clienthandler

import (
	pb "github.com/adamryman/ambition-model/ambition-service"
)

// CreateAction implements Service.
func CreateAction(IDCreateAction int64, UserIDCreateAction int64, NameCreateAction string) (*pb.Action, error) {
	request := pb.Action{
		ID:     IDCreateAction,
		UserID: UserIDCreateAction,
		Name:   NameCreateAction,
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
func ReadAction(IDReadAction int64, UserIDReadAction int64, NameReadAction string) (*pb.Action, error) {
	request := pb.Action{
		ID:     IDReadAction,
		UserID: UserIDReadAction,
		Name:   NameReadAction,
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
func ReadOccurrences(IDReadOccurrences int64, UserIDReadOccurrences int64, NameReadOccurrences string) (*pb.Action, error) {
	request := pb.Action{
		ID:     IDReadOccurrences,
		UserID: UserIDReadOccurrences,
		Name:   NameReadOccurrences,
	}
	return &request, nil
}
