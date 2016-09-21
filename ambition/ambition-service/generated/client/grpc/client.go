// Package grpc provides a gRPC client for the add service.
package grpc

import (
	//"time"

	//jujuratelimit "github.com/juju/ratelimit"
	//stdopentracing "github.com/opentracing/opentracing-go"
	//"github.com/sony/gobreaker"
	"google.golang.org/grpc"

	//"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	//"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/ratelimit"
	//"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/adamryman/ambition-truss/ambition/ambition-service"
	svc "github.com/adamryman/ambition-truss/ambition/ambition-service/generated"
	handler "github.com/adamryman/ambition-truss/ambition/ambition-service/handlers/server"
)

// New returns an AddService backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn /*, tracer stdopentracing.Tracer, logger log.Logger*/) handler.Service {
	// We construct a single ratelimiter middleware, to limit the total outgoing
	// QPS from this client to all methods on the remote instance. We also
	// construct per-endpoint circuitbreaker middlewares to demonstrate how
	// that's done, although they could easily be combined into a single breaker
	// for the entire remote instance, too.

	//limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(100, 100))

	var createactionEndpoint endpoint.Endpoint
	{
		createactionEndpoint = grpctransport.NewClient(
			conn,
			"ambition.AmbitionService",
			"CreateAction",
			svc.EncodeGRPCCreateActionRequest,
			svc.DecodeGRPCCreateActionResponse,
			pb.CreateActionResponse{},
			//grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "CreateAction", logger)),
		).Endpoint()
		//createactionEndpoint = opentracing.TraceClient(tracer, "CreateAction")(createactionEndpoint)
		//createactionEndpoint = limiter(createactionEndpoint)
		//createactionEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		//Name:    "CreateAction",
		//Timeout: 30 * time.Second,
		//}))(createactionEndpoint)
	}

	var createoccurrenceEndpoint endpoint.Endpoint
	{
		createoccurrenceEndpoint = grpctransport.NewClient(
			conn,
			"ambition.AmbitionService",
			"CreateOccurrence",
			svc.EncodeGRPCCreateOccurrenceRequest,
			svc.DecodeGRPCCreateOccurrenceResponse,
			pb.CreateOccurrenceResponse{},
			//grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "CreateOccurrence", logger)),
		).Endpoint()
		//createoccurrenceEndpoint = opentracing.TraceClient(tracer, "CreateOccurrence")(createoccurrenceEndpoint)
		//createoccurrenceEndpoint = limiter(createoccurrenceEndpoint)
		//createoccurrenceEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		//Name:    "CreateOccurrence",
		//Timeout: 30 * time.Second,
		//}))(createoccurrenceEndpoint)
	}

	return svc.Endpoints{

		CreateActionEndpoint:     createactionEndpoint,
		CreateOccurrenceEndpoint: createoccurrenceEndpoint,
	}
}
