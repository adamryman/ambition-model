package clienthandler

import (
	pb "github.com/adamryman/ambition-model/ambition-service"
)

// ReadActions implements Service.
func ReadActions(UserIdReadActions int64) (*pb.ReadActionsRequest, error) {

	request := pb.ReadActionsRequest{
		UserId: UserIdReadActions,
	}
	return &request, nil
}

// ReadAction implements Service.
func ReadAction(ActionIdReadAction int64, ActionNameReadAction string) (*pb.ReadActionRequest, error) {

	request := pb.ReadActionRequest{
		ActionId:   ActionIdReadAction,
		ActionName: ActionNameReadAction,
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
func ReadOccurrences(UserIdReadOccurrences int64, ActionIdReadOccurrences int64) (*pb.ReadOccurrencesRequest, error) {

	request := pb.ReadOccurrencesRequest{
		UserId:   UserIdReadOccurrences,
		ActionId: ActionIdReadOccurrences,
	}
	return &request, nil
}

// CreateOccurrence implements Service.
func CreateOccurrence(ActionIdCreateOccurrence int64, DatetimeCreateOccurrence string, DataCreateOccurrence string) (*pb.CreateOccurrenceRequest, error) {

	request := pb.CreateOccurrenceRequest{
		ActionId: ActionIdCreateOccurrence,
		Datetime: DatetimeCreateOccurrence,
		Data:     DataCreateOccurrence,
	}
	return &request, nil
}
