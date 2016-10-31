package model

import (
	"fmt"
	pb "github.com/adamryman/ambition-model/ambition-service"
	"github.com/adamryman/ambition-model/database"
	"os"
)

func init() {
	var err error
	err = database.New()
	if err != nil {
		os.Exit(1)
		panic(err)
	}
}

func CreateAction(in *pb.Action) (*pb.ActionResponse, error) {
	var resp pb.ActionResponse
	fmt.Println("CREATE ACTION")
	action := database.Action(in)
	err := action.Create()
	fmt.Println("CREATE ACTION DONE")

	if err != nil {
		resp.Error = err.Error()
	}
	resp.Action = action.Action

	return &resp, nil
}

func ReadAction(in *pb.Action) (*pb.ActionResponse, error) {
	var resp pb.ActionResponse
	action := database.Action(in)
	var err error

	switch {
	case in.ActionId != 0:
		err = action.ReadByActionId()
		fmt.Println("actionId")
	case in.TrelloId != "":
		err = action.ReadByTrelloId()
		fmt.Println("TrelloId")
	default:
		fmt.Println("test")
	}
	if err != nil {
		resp.Error = err.Error()
		return &resp, nil
	}
	resp.Action = in
	return &resp, nil
}

func CreateOccurrence(in *pb.CreateOccurrenceRequest) (*pb.OccurrenceResponse, error) {
	var resp pb.OccurrenceResponse

	occurrence, err := database.CreateOccurrence(in)
	if err != nil {
		resp.Error = err.Error()
		return &resp, nil
	}
	resp.Occurrence = occurrence

	return &resp, nil
}
