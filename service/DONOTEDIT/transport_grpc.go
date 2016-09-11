package addsvc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	//stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	//"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/adamryman/ambition-truss/service/DONOTEDIT/pb"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AddServer.
func MakeGRPCServer(ctx context.Context, endpoints Endpoints /*, tracer stdopentracing.Tracer, logger log.Logger*/) pb.AmbitionServiceServer {
	//options := []grpctransport.ServerOption{
	//grpctransport.ServerErrorLogger(logger),
	//}
	return &grpcServer{
		// ambitionservice

		createaction: grpctransport.NewServer(
			ctx,
			endpoints.CreateActionEndpoint,
			DecodeGRPCCreateActionRequest,
			EncodeGRPCCreateActionResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "CreateAction", logger)))...,
		),
		createoccurrence: grpctransport.NewServer(
			ctx,
			endpoints.CreateOccurrenceEndpoint,
			DecodeGRPCCreateOccurrenceRequest,
			EncodeGRPCCreateOccurrenceResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "CreateOccurrence", logger)))...,
		),
	}
}

type grpcServer struct {
	createaction     grpctransport.Handler
	createoccurrence grpctransport.Handler
}

// Methods

func (s *grpcServer) CreateAction(ctx context.Context, req *pb.CreateActionRequest) (*pb.CreateActionResponse, error) {
	_, rep, err := s.createaction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateActionResponse), nil
}

func (s *grpcServer) CreateOccurrence(ctx context.Context, req *pb.CreateOccurrenceRequest) (*pb.CreateOccurrenceResponse, error) {
	_, rep, err := s.createoccurrence.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateOccurrenceResponse), nil
}

// Server Decode

// DecodeGRPCCreateActionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createaction request to a user-domain createaction request. Primarily useful in a server.
func DecodeGRPCCreateActionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateActionRequest)
	return req, nil
}

// DecodeGRPCCreateOccurrenceRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createoccurrence request to a user-domain createoccurrence request. Primarily useful in a server.
func DecodeGRPCCreateOccurrenceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateOccurrenceRequest)
	return req, nil
}

// Client Decode

// DecodeGRPCCreateActionResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createaction reply to a user-domain createaction response. Primarily useful in a client.
func DecodeGRPCCreateActionResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateActionResponse)
	return reply, nil
}

// DecodeGRPCCreateOccurrenceResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createoccurrence reply to a user-domain createoccurrence response. Primarily useful in a client.
func DecodeGRPCCreateOccurrenceResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateOccurrenceResponse)
	return reply, nil
}

// Server Encode

// EncodeGRPCCreateActionResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain createaction response to a gRPC createaction reply. Primarily useful in a server.
func EncodeGRPCCreateActionResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.CreateActionResponse)
	return resp, nil
}

// EncodeGRPCCreateOccurrenceResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain createoccurrence response to a gRPC createoccurrence reply. Primarily useful in a server.
func EncodeGRPCCreateOccurrenceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.CreateOccurrenceResponse)
	return resp, nil
}

// Client Encode

// EncodeGRPCCreateActionRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createaction request to a gRPC createaction request. Primarily useful in a client.
func EncodeGRPCCreateActionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(pb.CreateActionRequest)
	return &req, nil
}

// EncodeGRPCCreateOccurrenceRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createoccurrence request to a gRPC createoccurrence request. Primarily useful in a client.
func EncodeGRPCCreateOccurrenceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(pb.CreateOccurrenceRequest)
	return &req, nil
}
