package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats. It also includes endpoint middlewares.

import (
	//_ "fmt"
	//_ "time"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	//_ "github.com/go-kit/kit/log"
	//_ "github.com/go-kit/kit/metrics"

	pb "github.com/adamryman/ambition-truss/ambition-service"
	handler "github.com/adamryman/ambition-truss/ambition-service/handlers/server"
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
}

// Endpoints

func (e Endpoints) CreateAction(ctx context.Context, in *pb.CreateActionRequest) (*pb.CreateActionResponse, error) {
	response, err := e.CreateActionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.CreateActionResponse), nil
}

func (e Endpoints) CreateOccurrence(ctx context.Context, in *pb.CreateOccurrenceRequest) (*pb.CreateOccurrenceResponse, error) {
	response, err := e.CreateOccurrenceEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.CreateOccurrenceResponse), nil
}

// Make Endpoints

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

// MIDDLE WARE

// EndpointInstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
//func EndpointInstrumentingMiddleware(duration metrics.TimeHistogram) endpoint.Middleware {
//return func(next endpoint.Endpoint) endpoint.Endpoint {
//return func(ctx context.Context, request interface{}) (response interface{}, err error) {

//defer func(begin time.Time) {
//f := metrics.Field{Key: "success", Value: fmt.Sprint(err == nil)}
//duration.With(f).Observe(time.Since(begin))
//}(time.Now())
//return next(ctx, request)

//}
//}
//}

// EndpointLoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
//func EndpointLoggingMiddleware(logger log.Logger) endpoint.Middleware {
//return func(next endpoint.Endpoint) endpoint.Endpoint {
//return func(ctx context.Context, request interface{}) (response interface{}, err error) {

//defer func(begin time.Time) {
//logger.Log("error", err, "took", time.Since(begin))
//}(time.Now())
//return next(ctx, request)

//}
//}
//}
