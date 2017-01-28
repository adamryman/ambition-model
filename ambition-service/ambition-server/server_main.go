package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"

	// 3d Party
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// Go Kit
	"github.com/go-kit/kit/log"

	// This Service
	pb "github.com/adamryman/ambition-model/ambition-service"
	svc "github.com/adamryman/ambition-model/ambition-service/generated"
	handler "github.com/adamryman/ambition-model/ambition-service/handlers/server"
	middlewares "github.com/adamryman/ambition-model/ambition-service/middlewares"
)

func main() {
	var (
		debugAddr = flag.String("debug.addr", ":5060", "Debug and metrics listen address")
		httpAddr  = flag.String("http.addr", ":5050", "HTTP listen address")
		grpcAddr  = flag.String("grpc.addr", ":5040", "gRPC (HTTP) listen address")
	)
	flag.Parse()

	// Use environment variables, if set. Flags have priority over Env vars.
	if os.Getenv("DEBUG_ADDR") != "" && *debugAddr == ":5060" {
		*debugAddr = os.Getenv("DEBUG_ADDR")
	}
	if os.Getenv("HTTP_ADDR") != "" && *httpAddr == ":5050" {
		*httpAddr = os.Getenv("HTTP_ADDR")
	}
	if os.Getenv("GRPC_ADDR") != "" && *grpcAddr == ":5040" {
		*grpcAddr = os.Getenv("GRPC_ADDR")
	}
	if os.Getenv("PORT") != "" && *httpAddr == ":5050" {
		*httpAddr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	// Business domain.
	var service pb.AmbitionServer
	{
		service = handler.NewService()
		// Wrap Service with middlewares. See ../middlewares/service.go
		service = middlewares.WrapService(service)
	}

	// Endpoint domain.
	var (
		createactionEndpoint     = svc.MakeCreateActionEndpoint(service)
		createoccurrenceEndpoint = svc.MakeCreateOccurrenceEndpoint(service)
		readactionEndpoint       = svc.MakeReadActionEndpoint(service)
		readactionsEndpoint      = svc.MakeReadActionsEndpoint(service)
		readoccurrencesEndpoint  = svc.MakeReadOccurrencesEndpoint(service)
	)

	// Wrap Endpoints with middlewares. See ../middlewares/endpoints.go
	endpoints := svc.Endpoints{
		CreateActionEndpoint:     createactionEndpoint,
		CreateOccurrenceEndpoint: createoccurrenceEndpoint,
		ReadActionEndpoint:       readactionEndpoint,
		ReadActionsEndpoint:      readactionsEndpoint,
		ReadOccurrencesEndpoint:  readoccurrencesEndpoint,
	}

	// Wrap selected Endpoints with middlewares. See ../middlewares/endpoints.go
	endpoints = middlewares.WrapEndpoints(endpoints)

	// Mechanical domain.
	errc := make(chan error)
	ctx := context.Background()

	// Interrupt handler.
	go handler.InterruptHandler(errc)

	// Debug listener.
	go func() {
		logger := log.NewContext(logger).With("transport", "debug")

		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

		logger.Log("addr", *debugAddr)
		errc <- http.ListenAndServe(*debugAddr, m)
	}()

	// HTTP transport.
	go func() {
		logger := log.NewContext(logger).With("transport", "HTTP")
		h := svc.MakeHTTPHandler(ctx, endpoints, logger)
		logger.Log("addr", *httpAddr)
		errc <- http.ListenAndServe(*httpAddr, h)
	}()

	// gRPC transport.
	go func() {
		logger := log.NewContext(logger).With("transport", "gRPC")

		ln, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errc <- err
			return
		}

		srv := svc.MakeGRPCServer(ctx, endpoints)
		s := grpc.NewServer()
		pb.RegisterAmbitionServer(s, srv)

		logger.Log("addr", *grpcAddr)
		errc <- s.Serve(ln)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
