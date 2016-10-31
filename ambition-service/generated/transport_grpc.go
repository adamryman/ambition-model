package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/adamryman/ambition-model/ambition-service"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AmbitionServiceServer.
func MakeGRPCServer(ctx context.Context, endpoints Endpoints) pb.AmbitionServiceServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	return &grpcServer{
		// ambitionservice

		readactions: grpctransport.NewServer(
			ctx,
			endpoints.ReadActionsEndpoint,
			DecodeGRPCReadActionsRequest,
			EncodeGRPCReadActionsResponse,
			serverOptions...,
		),
		readaction: grpctransport.NewServer(
			ctx,
			endpoints.ReadActionEndpoint,
			DecodeGRPCReadActionRequest,
			EncodeGRPCReadActionResponse,
			serverOptions...,
		),
		createaction: grpctransport.NewServer(
			ctx,
			endpoints.CreateActionEndpoint,
			DecodeGRPCCreateActionRequest,
			EncodeGRPCCreateActionResponse,
			serverOptions...,
		),
		readoccurrences: grpctransport.NewServer(
			ctx,
			endpoints.ReadOccurrencesEndpoint,
			DecodeGRPCReadOccurrencesRequest,
			EncodeGRPCReadOccurrencesResponse,
			serverOptions...,
		),
		createoccurrence: grpctransport.NewServer(
			ctx,
			endpoints.CreateOccurrenceEndpoint,
			DecodeGRPCCreateOccurrenceRequest,
			EncodeGRPCCreateOccurrenceResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the AmbitionServiceServer interface
type grpcServer struct {
	readactions      grpctransport.Handler
	readaction       grpctransport.Handler
	createaction     grpctransport.Handler
	readoccurrences  grpctransport.Handler
	createoccurrence grpctransport.Handler
}

// Methods for grpcServer to implement AmbitionServiceServer interface

func (s *grpcServer) ReadActions(ctx context.Context, req *pb.ReadActionsRequest) (*pb.ActionsResponse, error) {
	_, rep, err := s.readactions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ActionsResponse), nil
}

func (s *grpcServer) ReadAction(ctx context.Context, req *pb.Action) (*pb.ActionResponse, error) {
	_, rep, err := s.readaction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ActionResponse), nil
}

func (s *grpcServer) CreateAction(ctx context.Context, req *pb.Action) (*pb.ActionResponse, error) {
	_, rep, err := s.createaction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ActionResponse), nil
}

func (s *grpcServer) ReadOccurrences(ctx context.Context, req *pb.Occurrence) (*pb.OccurrenceResponse, error) {
	_, rep, err := s.readoccurrences.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.OccurrenceResponse), nil
}

func (s *grpcServer) CreateOccurrence(ctx context.Context, req *pb.Occurrence) (*pb.OccurrenceResponse, error) {
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
	req := grpcReq.(*pb.ReadActionsRequest)
	return req, nil
}

// DecodeGRPCReadActionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC readaction request to a user-domain readaction request. Primarily useful in a server.
func DecodeGRPCReadActionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Action)
	return req, nil
}

// DecodeGRPCCreateActionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createaction request to a user-domain createaction request. Primarily useful in a server.
func DecodeGRPCCreateActionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Action)
	return req, nil
}

// DecodeGRPCReadOccurrencesRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC readoccurrences request to a user-domain readoccurrences request. Primarily useful in a server.
func DecodeGRPCReadOccurrencesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Occurrence)
	return req, nil
}

// DecodeGRPCCreateOccurrenceRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createoccurrence request to a user-domain createoccurrence request. Primarily useful in a server.
func DecodeGRPCCreateOccurrenceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Occurrence)
	return req, nil
}

// Client Decode

// DecodeGRPCReadActionsResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC readactions reply to a user-domain readactions response. Primarily useful in a client.
func DecodeGRPCReadActionsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ActionsResponse)
	return reply, nil
}

// DecodeGRPCReadActionResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC readaction reply to a user-domain readaction response. Primarily useful in a client.
func DecodeGRPCReadActionResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
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
	resp := response.(*pb.ActionsResponse)
	return resp, nil
}

// EncodeGRPCReadActionResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain readaction response to a gRPC readaction reply. Primarily useful in a server.
func EncodeGRPCReadActionResponse(_ context.Context, response interface{}) (interface{}, error) {
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
	req := request.(*pb.ReadActionsRequest)
	return req, nil
}

// EncodeGRPCReadActionRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain readaction request to a gRPC readaction request. Primarily useful in a client.
func EncodeGRPCReadActionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Action)
	return req, nil
}

// EncodeGRPCCreateActionRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createaction request to a gRPC createaction request. Primarily useful in a client.
func EncodeGRPCCreateActionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Action)
	return req, nil
}

// EncodeGRPCReadOccurrencesRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain readoccurrences request to a gRPC readoccurrences request. Primarily useful in a client.
func EncodeGRPCReadOccurrencesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Occurrence)
	return req, nil
}

// EncodeGRPCCreateOccurrenceRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createoccurrence request to a gRPC createoccurrence request. Primarily useful in a client.
func EncodeGRPCCreateOccurrenceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Occurrence)
	return req, nil
}

// Helpers

func metadataToContext(ctx context.Context, md *metadata.MD) context.Context {
	for k, v := range *md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
