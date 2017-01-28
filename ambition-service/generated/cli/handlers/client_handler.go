package clienthandler

import (
	pb "github.com/adamryman/ambition-model/ambition-service"
)

// CreateAction implements Service.
func CreateAction(IDCreateAction int64, UserIDCreateAction int64, ActionNameCreateAction string) (*pb.Action, error) {
	request := pb.Action{
		ID:         IDCreateAction,
		UserID:     UserIDCreateAction,
		ActionName: ActionNameCreateAction,
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

// ReadActions implements Service.
func ReadActions(UserIDReadActions int64) (*pb.User, error) {
	request := pb.User{
		UserID: UserIDReadActions,
	}
	return &request, nil
}

// ReadAction implements Service.
func ReadAction(IDReadAction int64, UserIDReadAction int64, ActionNameReadAction string) (*pb.Action, error) {
	request := pb.Action{
		ID:         IDReadAction,
		UserID:     UserIDReadAction,
		ActionName: ActionNameReadAction,
	}
	return &request, nil
}
