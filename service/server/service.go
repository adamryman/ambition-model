package handler

import (
	_ "errors"
	_ "time"

	"golang.org/x/net/context"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/metrics"

	pb "github.com/adamryman/ambition-truss/service/DONOTEDIT/pb"
)

func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

func (s basicService) CreateAction(ctx context.Context, in pb.CreateActionRequest) (pb.CreateActionResponse, error) {
	_ = ctx
	_ = in
	response := pb.CreateActionResponse{}
	return response, nil
}

func (s basicService) CreateOccurrence(ctx context.Context, in pb.CreateOccurrenceRequest) (pb.CreateOccurrenceResponse, error) {
	_ = ctx
	_ = in
	response := pb.CreateOccurrenceResponse{}
	return response, nil
}

type Service interface {
	CreateAction(ctx context.Context, in pb.CreateActionRequest) (pb.CreateActionResponse, error)
	CreateOccurrence(ctx context.Context, in pb.CreateOccurrenceRequest) (pb.CreateOccurrenceResponse, error)
}
