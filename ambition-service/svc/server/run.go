package server

import (
	"net"
	"net/http"
	"net/http/pprof"

	// 3d Party
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// Go Kit
	"github.com/go-kit/kit/log"

	// This Service
	pb "github.com/adamryman/ambition-model/ambition-service"
	"github.com/adamryman/ambition-model/ambition-service/handlers"
	"github.com/adamryman/ambition-model/ambition-service/middlewares"
	"github.com/adamryman/ambition-model/ambition-service/svc"
)

// Config contains the required fields for running a server
type Config struct {
	HTTPAddr  string
	DebugAddr string
	GRPCAddr  string
}

// Run starts a new http server, gRPC server, and a debug server with the
// passed config and logger
func Run(cfg Config, logger log.Logger) {
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	// Business domain.
	var service pb.AmbitionServer
	{
		service = handlers.NewService()
		// Wrap Service with middlewares. See middlewares/service.go
		service = middlewares.WrapService(service)
	}

	// Endpoint domain.
	var (
		createactionEndpoint          = svc.MakeCreateActionEndpoint(service)
		createoccurrenceEndpoint      = svc.MakeCreateOccurrenceEndpoint(service)
		readactionEndpoint            = svc.MakeReadActionEndpoint(service)
		readactionsEndpoint           = svc.MakeReadActionsEndpoint(service)
		readoccurrencesbydateEndpoint = svc.MakeReadOccurrencesByDateEndpoint(service)
		readoccurrencesEndpoint       = svc.MakeReadOccurrencesEndpoint(service)
	)

	endpoints := svc.Endpoints{
		CreateActionEndpoint:          createactionEndpoint,
		CreateOccurrenceEndpoint:      createoccurrenceEndpoint,
		ReadActionEndpoint:            readactionEndpoint,
		ReadActionsEndpoint:           readactionsEndpoint,
		ReadOccurrencesByDateEndpoint: readoccurrencesbydateEndpoint,
		ReadOccurrencesEndpoint:       readoccurrencesEndpoint,
	}

	// Wrap selected Endpoints with middlewares. See middlewares/endpoints.go
	endpoints = middlewares.WrapEndpoints(endpoints)

	// Mechanical domain.
	errc := make(chan error)
	ctx := context.Background()

	// Interrupt handler.
	go handlers.InterruptHandler(errc)

	// Debug listener.
	go func() {
		logger := log.NewContext(logger).With("transport", "debug")

		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

		logger.Log("addr", cfg.DebugAddr)
		errc <- http.ListenAndServe(cfg.DebugAddr, m)
	}()

	// HTTP transport.
	go func() {
		logger := log.NewContext(logger).With("transport", "HTTP")
		h := svc.MakeHTTPHandler(ctx, endpoints, logger)
		logger.Log("addr", cfg.HTTPAddr)
		errc <- http.ListenAndServe(cfg.HTTPAddr, h)
	}()

	// gRPC transport.
	go func() {
		logger := log.NewContext(logger).With("transport", "gRPC")

		ln, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			errc <- err
			return
		}

		srv := svc.MakeGRPCServer(ctx, endpoints)
		s := grpc.NewServer()
		pb.RegisterAmbitionServer(s, srv)

		logger.Log("addr", cfg.GRPCAddr)
		errc <- s.Serve(ln)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
