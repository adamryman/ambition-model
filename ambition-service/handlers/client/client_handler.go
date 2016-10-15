package clienthandler

import (
	pb "github.com/adamryman/ambition-truss/ambition-service"
)

// ReadActions implements Service.
func ReadActions(UserIdReadActions int64) (*pb.ActionsRequest, error) {

	request := pb.ActionsRequest{
		UserId: UserIdReadActions,
	}
	return &request, nil
}

// CreateAction implements Service.
func CreateAction(UserIdCreateAction int64, ActionNameCreateAction string) (*pb.CreateActionRequest, error) {

	request := pb.CreateActionRequest{
		UserId:     UserIdCreateAction,
		ActionName: ActionNameCreateAction,
	}
	return &request, nil
}

// ReadOccurrences implements Service.
func ReadOccurrences(UserIdReadOccurrences int64, ActionIdReadOccurrences int64) (*pb.OccurrencesRequest, error) {

	request := pb.OccurrencesRequest{
		UserId:   UserIdReadOccurrences,
		ActionId: ActionIdReadOccurrences,
	}
	return &request, nil
}

// CreateOccurrence implements Service.
func CreateOccurrence(ActionIdCreateOccurrence int64, EpocTimeCreateOccurrence int64) (*pb.CreateOccurrenceRequest, error) {

	request := pb.CreateOccurrenceRequest{
		ActionId: ActionIdCreateOccurrence,
		EpocTime: EpocTimeCreateOccurrence,
	}
	return &request, nil
}
