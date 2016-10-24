// Package grpc provides a gRPC client for the AmbitionService service.
package grpc

import (
	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/adamryman/ambition-model/ambition-service"
	svc "github.com/adamryman/ambition-model/ambition-service/generated"
	handler "github.com/adamryman/ambition-model/ambition-service/handlers/server"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn) handler.Service {
	//options := []grpctransport.ServerOption{
	//grpctransport.ServerBefore(),
	//}
	var readactionsEndpoint endpoint.Endpoint
	{
		readactionsEndpoint = grpctransport.NewClient(
			conn,
			"ambition.AmbitionService",
			"ReadActions",
			svc.EncodeGRPCReadActionsRequest,
			svc.DecodeGRPCReadActionsResponse,
			pb.ActionsResponse{},
			//options...,
		).Endpoint()
	}

	var readactionEndpoint endpoint.Endpoint
	{
		readactionEndpoint = grpctransport.NewClient(
			conn,
			"ambition.AmbitionService",
			"ReadAction",
			svc.EncodeGRPCReadActionRequest,
			svc.DecodeGRPCReadActionResponse,
			pb.ActionResponse{},
			//options...,
		).Endpoint()
	}

	var createactionEndpoint endpoint.Endpoint
	{
		createactionEndpoint = grpctransport.NewClient(
			conn,
			"ambition.AmbitionService",
			"CreateAction",
			svc.EncodeGRPCCreateActionRequest,
			svc.DecodeGRPCCreateActionResponse,
			pb.ActionResponse{},
			//options...,
		).Endpoint()
	}

	var readoccurrencesEndpoint endpoint.Endpoint
	{
		readoccurrencesEndpoint = grpctransport.NewClient(
			conn,
			"ambition.AmbitionService",
			"ReadOccurrences",
			svc.EncodeGRPCReadOccurrencesRequest,
			svc.DecodeGRPCReadOccurrencesResponse,
			pb.OccurrenceResponse{},
			//options...,
		).Endpoint()
	}

	var createoccurrenceEndpoint endpoint.Endpoint
	{
		createoccurrenceEndpoint = grpctransport.NewClient(
			conn,
			"ambition.AmbitionService",
			"CreateOccurrence",
			svc.EncodeGRPCCreateOccurrenceRequest,
			svc.DecodeGRPCCreateOccurrenceResponse,
			pb.OccurrenceResponse{},
			//options...,
		).Endpoint()
	}

	return svc.Endpoints{
		ReadActionsEndpoint:      readactionsEndpoint,
		ReadActionEndpoint:       readactionEndpoint,
		CreateActionEndpoint:     createactionEndpoint,
		ReadOccurrencesEndpoint:  readoccurrencesEndpoint,
		CreateOccurrenceEndpoint: createoccurrenceEndpoint,
	}
}
