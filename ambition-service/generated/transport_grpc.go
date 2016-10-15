package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	//stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	//"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/adamryman/ambition-truss/ambition-service"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AddServer.
func MakeGRPCServer(ctx context.Context, endpoints Endpoints /*, tracer stdopentracing.Tracer, logger log.Logger*/) pb.AmbitionServiceServer {
	//options := []grpctransport.ServerOption{
	//grpctransport.ServerErrorLogger(logger),
	//}
	return &grpcServer{
		// ambitionservice

		readactions: grpctransport.NewServer(
			ctx,
			endpoints.ReadActionsEndpoint,
			DecodeGRPCReadActionsRequest,
			EncodeGRPCReadActionsResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "ReadActions", logger)))...,
		),
		createaction: grpctransport.NewServer(
			ctx,
			endpoints.CreateActionEndpoint,
			DecodeGRPCCreateActionRequest,
			EncodeGRPCCreateActionResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "CreateAction", logger)))...,
		),
		readoccurrences: grpctransport.NewServer(
			ctx,
			endpoints.ReadOccurrencesEndpoint,
			DecodeGRPCReadOccurrencesRequest,
			EncodeGRPCReadOccurrencesResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "ReadOccurrences", logger)))...,
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
	readactions      grpctransport.Handler
	createaction     grpctransport.Handler
	readoccurrences  grpctransport.Handler
	createoccurrence grpctransport.Handler
}

// Methods

func (s *grpcServer) ReadActions(ctx context.Context, req *pb.ActionsRequest) (*pb.ActionResponse, error) {
	_, rep, err := s.readactions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ActionResponse), nil
}

func (s *grpcServer) CreateAction(ctx context.Context, req *pb.CreateActionRequest) (*pb.ActionResponse, error) {
	_, rep, err := s.createaction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ActionResponse), nil
}

func (s *grpcServer) ReadOccurrences(ctx context.Context, req *pb.OccurrencesRequest) (*pb.OccurrenceResponse, error) {
	_, rep, err := s.readoccurrences.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.OccurrenceResponse), nil
}

func (s *grpcServer) CreateOccurrence(ctx context.Context, req *pb.CreateOccurrenceRequest) (*pb.OccurrenceResponse, error) {
	_, rep, err := s.createoccurrence.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.OccurrenceResponse), nil
}

// Server Decode

// DecodeGRPCReadActionsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC readactions request to a user-domain readactions request. Primarily useful in a server.
func DecodeGRPCReadActionsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ActionsRequest)
	return req, nil
}

// DecodeGRPCCreateActionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createaction request to a user-domain createaction request. Primarily useful in a server.
func DecodeGRPCCreateActionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateActionRequest)
	return req, nil
}

// DecodeGRPCReadOccurrencesRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC readoccurrences request to a user-domain readoccurrences request. Primarily useful in a server.
func DecodeGRPCReadOccurrencesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.OccurrencesRequest)
	return req, nil
}

// DecodeGRPCCreateOccurrenceRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createoccurrence request to a user-domain createoccurrence request. Primarily useful in a server.
func DecodeGRPCCreateOccurrenceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateOccurrenceRequest)
	return req, nil
}

// Client Decode

// DecodeGRPCReadActionsResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC readactions reply to a user-domain readactions response. Primarily useful in a client.
func DecodeGRPCReadActionsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ActionResponse)
	return reply, nil
}

// DecodeGRPCCreateActionResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createaction reply to a user-domain createaction response. Primarily useful in a client.
func DecodeGRPCCreateActionResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ActionResponse)
	return reply, nil
}

// DecodeGRPCReadOccurrencesResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC readoccurrences reply to a user-domain readoccurrences response. Primarily useful in a client.
func DecodeGRPCReadOccurrencesResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.OccurrenceResponse)
	return reply, nil
}

// DecodeGRPCCreateOccurrenceResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createoccurrence reply to a user-domain createoccurrence response. Primarily useful in a client.
func DecodeGRPCCreateOccurrenceResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.OccurrenceResponse)
	return reply, nil
}

// Server Encode

// EncodeGRPCReadActionsResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain readactions response to a gRPC readactions reply. Primarily useful in a server.
func EncodeGRPCReadActionsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ActionResponse)
	return resp, nil
}

// EncodeGRPCCreateActionResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain createaction response to a gRPC createaction reply. Primarily useful in a server.
func EncodeGRPCCreateActionResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ActionResponse)
	return resp, nil
}

// EncodeGRPCReadOccurrencesResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain readoccurrences response to a gRPC readoccurrences reply. Primarily useful in a server.
func EncodeGRPCReadOccurrencesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.OccurrenceResponse)
	return resp, nil
}

// EncodeGRPCCreateOccurrenceResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain createoccurrence response to a gRPC createoccurrence reply. Primarily useful in a server.
func EncodeGRPCCreateOccurrenceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.OccurrenceResponse)
	return resp, nil
}

// Client Encode

// EncodeGRPCReadActionsRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain readactions request to a gRPC readactions request. Primarily useful in a client.
func EncodeGRPCReadActionsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ActionsRequest)
	return req, nil
}

// EncodeGRPCCreateActionRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createaction request to a gRPC createaction request. Primarily useful in a client.
func EncodeGRPCCreateActionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateActionRequest)
	return req, nil
}

// EncodeGRPCReadOccurrencesRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain readoccurrences request to a gRPC readoccurrences request. Primarily useful in a client.
func EncodeGRPCReadOccurrencesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.OccurrencesRequest)
	return req, nil
}

// EncodeGRPCCreateOccurrenceRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createoccurrence request to a gRPC createoccurrence request. Primarily useful in a client.
func EncodeGRPCCreateOccurrenceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateOccurrenceRequest)
	return req, nil
}
