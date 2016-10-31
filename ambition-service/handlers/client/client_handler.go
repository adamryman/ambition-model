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
func ReadAction(ActionIdReadAction int64, UserIdReadAction int64, ActionNameReadAction string, TrelloIdReadAction string) (*pb.Action, error) {

	request := pb.Action{
		ActionId:   ActionIdReadAction,
		UserId:     UserIdReadAction,
		ActionName: ActionNameReadAction,
		TrelloId:   TrelloIdReadAction,
	}
	return &request, nil
}

// CreateAction implements Service.
func CreateAction(ActionIdCreateAction int64, UserIdCreateAction int64, ActionNameCreateAction string, TrelloIdCreateAction string) (*pb.Action, error) {

	request := pb.Action{
		ActionId:   ActionIdCreateAction,
		UserId:     UserIdCreateAction,
		ActionName: ActionNameCreateAction,
		TrelloId:   TrelloIdCreateAction,
	}
	return &request, nil
}

// ReadOccurrences implements Service.
func ReadOccurrences(OccurrenceIdReadOccurrences int64, ActionIdReadOccurrences int64, DatetimeReadOccurrences string, DataReadOccurrences string) (*pb.Occurrence, error) {

	request := pb.Occurrence{
		OccurrenceId: OccurrenceIdReadOccurrences,
		ActionId:     ActionIdReadOccurrences,
		Datetime:     DatetimeReadOccurrences,
		Data:         DataReadOccurrences,
	}
	return &request, nil
}

// CreateOccurrence implements Service.
func CreateOccurrence(OccurrenceIdCreateOccurrence int64, ActionIdCreateOccurrence int64, DatetimeCreateOccurrence string, DataCreateOccurrence string) (*pb.Occurrence, error) {

	request := pb.Occurrence{
		OccurrenceId: OccurrenceIdCreateOccurrence,
		ActionId:     ActionIdCreateOccurrence,
		Datetime:     DatetimeCreateOccurrence,
		Data:         DataCreateOccurrence,
	}
	return &request, nil
}
