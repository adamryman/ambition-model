package model

import (
	//"fmt"
	pb "github.com/adamryman/ambition-truss/ambition-service"
	"github.com/adamryman/ambition-truss/database"
)

func init() {
	var err error
	err = database.New()
	if err != nil {
		panic(err)
	}
}

func CreateAction(in *pb.CreateActionRequest) (*pb.ActionResponse, error) {
	var resp *pb.ActionResponse

	action, err := database.CreateAction(in)
	if err != nil {
		resp.Error = err.Error()
		return resp, nil
	}
	resp.Action = action

	return resp, nil
}

func ReadAction(in *pb.ReadActionRequest) (*pb.ActionResponse, error) {
	var resp *pb.ActionResponse

	action, err := database.ReadAction(in)
	if err != nil {
		resp.Error = err.Error()
		return resp, nil
	}
	resp.Action = action
	return resp, nil

}

func CreateOccurrence(in *pb.CreateOccurrenceRequest) (*pb.OccurrenceResponse, error) {
	var resp *pb.OccurrenceResponse

	occurrence, err := database.CreateOccurrence(in)
	if err != nil {
		resp.Error = err.Error()
		return resp, nil
	}
	resp.Occurrence = occurrence

	return resp, nil
}
