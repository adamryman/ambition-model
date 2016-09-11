package clienthandler

import (
	pb "github.com/adamryman/ambition-truss/service/DONOTEDIT/pb"
)

func CreateAction(UserIdCreateAction int64, ActionNameCreateAction string) (pb.CreateActionRequest, error) {
	request := pb.CreateActionRequest{
		UserId:     UserIdCreateAction,
		ActionName: ActionNameCreateAction}
	return request, nil
}

func CreateOccurrence(ActionIdCreateOccurrence int64, EpocTimeCreateOccurrence int64) (pb.CreateOccurrenceRequest, error) {
	request := pb.CreateOccurrenceRequest{
		ActionId: ActionIdCreateOccurrence,
		EpocTime: EpocTimeCreateOccurrence}
	return request, nil
}
