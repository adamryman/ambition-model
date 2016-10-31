// Package grpc provides a gRPC client for the AmbitionService service.
package grpc

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/adamryman/ambition-model/ambition-service"
	svc "github.com/adamryman/ambition-model/ambition-service/generated"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.AmbitionServiceServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var readactionsEndpoint endpoint.Endpoint
	{
		readactionsEndpoint = grpctransport.NewClient(
			conn,
			"ambition.AmbitionService",
			"ReadActions",
			svc.EncodeGRPCReadActionsRequest,
			svc.DecodeGRPCReadActionsResponse,
			pb.ActionsResponse{},
			clientOptions...,
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
			clientOptions...,
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
			clientOptions...,
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
			clientOptions...,
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
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		ReadActionsEndpoint:      readactionsEndpoint,
		ReadActionEndpoint:       readactionEndpoint,
		CreateActionEndpoint:     createactionEndpoint,
		ReadOccurrencesEndpoint:  readoccurrencesEndpoint,
		CreateOccurrenceEndpoint: createoccurrenceEndpoint,
	}, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.RequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
