package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats. It also includes endpoint middlewares.

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/adamryman/ambition-model/ambition-service"
	handler "github.com/adamryman/ambition-model/ambition-service/handlers/server"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them
// into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	ReadActionsEndpoint      endpoint.Endpoint
	ReadActionEndpoint       endpoint.Endpoint
	CreateActionEndpoint     endpoint.Endpoint
	ReadOccurrencesEndpoint  endpoint.Endpoint
	CreateOccurrenceEndpoint endpoint.Endpoint
}

// Endpoints

func (e Endpoints) ReadActions(ctx context.Context, in *pb.ReadActionsRequest) (*pb.ActionsResponse, error) {
	response, err := e.ReadActionsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ActionsResponse), nil
}

func (e Endpoints) ReadAction(ctx context.Context, in *pb.ReadActionRequest) (*pb.ActionResponse, error) {
	response, err := e.ReadActionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ActionResponse), nil
}

func (e Endpoints) CreateAction(ctx context.Context, in *pb.CreateActionRequest) (*pb.ActionResponse, error) {
	response, err := e.CreateActionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ActionResponse), nil
}

func (e Endpoints) ReadOccurrences(ctx context.Context, in *pb.ReadOccurrencesRequest) (*pb.OccurrenceResponse, error) {
	response, err := e.ReadOccurrencesEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.OccurrenceResponse), nil
}

func (e Endpoints) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.OccurrenceResponse, error) {
	response, err := e.CreateOccurrenceEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.OccurrenceResponse), nil
}

// Make Endpoints

func MakeReadActionsEndpoint(s handler.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.ReadActionsRequest)
		v, err := s.ReadActions(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeReadActionEndpoint(s handler.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.ReadActionRequest)
		v, err := s.ReadAction(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeCreateActionEndpoint(s handler.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CreateActionRequest)
		v, err := s.CreateAction(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeReadOccurrencesEndpoint(s handler.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.ReadOccurrencesRequest)
		v, err := s.ReadOccurrences(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeCreateOccurrenceEndpoint(s handler.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CreateOccurrenceRequest)
		v, err := s.CreateOccurrence(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
