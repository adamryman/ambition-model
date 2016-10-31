package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	// 3d Party
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// Go Kit
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	// This Service
	pb "github.com/adamryman/ambition-model/ambition-service"
	svc "github.com/adamryman/ambition-model/ambition-service/generated"
	handler "github.com/adamryman/ambition-model/ambition-service/handlers/server"
)

func main() {
	var (
		debugAddr = flag.String("debug.addr", ":5060", "Debug and metrics listen address")
		httpAddr  = flag.String("http.addr", ":5050", "HTTP listen address")
		grpcAddr  = flag.String("grpc.addr", ":5040", "gRPC (HTTP) listen address")
	)
	flag.Parse()

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
	var service pb.AmbitionServiceServer
	{
		service = handler.NewService()
		// add service logging and metrics here
	}

	// Endpoint domain.

	var readactionsEndpoint endpoint.Endpoint
	{
		readactionsEndpoint = svc.MakeReadActionsEndpoint(service)
		// Add endpoint tracing, instrumentation and logging here
	}

	var readactionEndpoint endpoint.Endpoint
	{
		readactionEndpoint = svc.MakeReadActionEndpoint(service)
		// Add endpoint tracing, instrumentation and logging here
	}

	var createactionEndpoint endpoint.Endpoint
	{
		createactionEndpoint = svc.MakeCreateActionEndpoint(service)
		// Add endpoint tracing, instrumentation and logging here
	}

	var readoccurrencesEndpoint endpoint.Endpoint
	{
		readoccurrencesEndpoint = svc.MakeReadOccurrencesEndpoint(service)
		// Add endpoint tracing, instrumentation and logging here
	}

	var createoccurrenceEndpoint endpoint.Endpoint
	{
		createoccurrenceEndpoint = svc.MakeCreateOccurrenceEndpoint(service)
		// Add endpoint tracing, instrumentation and logging here
	}

	endpoints := svc.Endpoints{
		ReadActionsEndpoint:      readactionsEndpoint,
		ReadActionEndpoint:       readactionEndpoint,
		CreateActionEndpoint:     createactionEndpoint,
		ReadOccurrencesEndpoint:  readoccurrencesEndpoint,
		CreateOccurrenceEndpoint: createoccurrenceEndpoint,
	}

	// Mechanical domain.
	errc := make(chan error)
	ctx := context.Background()

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

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
		pb.RegisterAmbitionServiceServer(s, srv)

		logger.Log("addr", *grpcAddr)
		errc <- s.Serve(ln)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
