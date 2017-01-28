package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats. It also includes endpoint middlewares.

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/adamryman/ambition-model/ambition-service"
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
	CreateActionEndpoint     endpoint.Endpoint
	CreateOccurrenceEndpoint endpoint.Endpoint
	ReadActionsEndpoint      endpoint.Endpoint
	ReadActionEndpoint       endpoint.Endpoint
}

// Endpoints

func (e Endpoints) CreateAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	response, err := e.CreateActionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Action), nil
}

func (e Endpoints) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.Occurrence, error) {
	response, err := e.CreateOccurrenceEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Occurrence), nil
}

func (e Endpoints) ReadActions(ctx context.Context, in *pb.User) (*pb.ActionsResponse, error) {
	response, err := e.ReadActionsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ActionsResponse), nil
}

func (e Endpoints) ReadAction(ctx context.Context, in *pb.Action) (*pb.Action, error) {
	response, err := e.ReadActionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Action), nil
}

// Make Endpoints

func MakeCreateActionEndpoint(s pb.AmbitionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Action)
		v, err := s.CreateAction(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeCreateOccurrenceEndpoint(s pb.AmbitionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CreateOccurrenceRequest)
		v, err := s.CreateOccurrence(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeReadActionsEndpoint(s pb.AmbitionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.User)
		v, err := s.ReadActions(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeReadActionEndpoint(s pb.AmbitionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Action)
		v, err := s.ReadAction(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAll wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// The middlewares will be applied in the order passed, with the first
// middleware being the outermost middleware.
func (e *Endpoints) WrapAll(middlewares ...endpoint.Middleware) {
	if len(middlewares) == 0 {
		return
	}
	m := endpoint.Chain(middlewares[0], middlewares[1:]...)

	e.CreateActionEndpoint = m(e.CreateActionEndpoint)
	e.CreateOccurrenceEndpoint = m(e.CreateOccurrenceEndpoint)
	e.ReadActionsEndpoint = m(e.ReadActionsEndpoint)
	e.ReadActionEndpoint = m(e.ReadActionEndpoint)
}
