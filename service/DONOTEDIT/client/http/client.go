// Package http provides an HTTP client for the add service.

package http

import (
	"net/url"
	"strings"
	//"time"

	//jujuratelimit "github.com/juju/ratelimit"
	//stdopentracing "github.com/opentracing/opentracing-go"
	//"github.com/sony/gobreaker"

	//"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	//"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/ratelimit"
	//"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"

	// This Service
	svc "github.com/adamryman/ambition-truss/service/DONOTEDIT"
	handler "github.com/adamryman/ambition-truss/service/server"
	//pb "github.com/adamryman/ambition-truss/service/DONOTEDIT/pb"
)

var (
	_ = endpoint.Chain
	_ = httptransport.NewClient
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".

//func New(instance string, tracer stdopentracing.Tracer, logger log.Logger) (addsvc.Service, error) {
func New(instance string) (handler.Service, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	_ = u

	// We construct a single ratelimiter middleware, to limit the total outgoing
	// QPS from this client to all methods on the remote instance. We also
	// construct per-endpoint circuitbreaker middlewares to demonstrate how
	// that's done, although they could easily be combined into a single breaker
	// for the entire remote instance, too.

	//limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(100, 100))

	var CreateActionZeroEndpoint endpoint.Endpoint
	{
		CreateActionZeroEndpoint = httptransport.NewClient(
			"post",
			copyURL(u, "/action"),
			svc.EncodeHTTPCreateActionZeroRequest,
			svc.DecodeHTTPCreateActionResponse,
			//httptransport.ClientBefore(opentracing.FromHTTPRequest(tracer, "Sum", logger)),
		).Endpoint()
		/*
			sumEndpoint = opentracing.TraceClient(tracer, "Sum")(sumEndpoint)
			sumEndpoint = limiter(sumEndpoint)
			sumEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
				Name:    "Sum",
				Timeout: 30 * time.Second,
			}))(sumEndpoint)
		*/
	}

	var CreateOccurrenceZeroEndpoint endpoint.Endpoint
	{
		CreateOccurrenceZeroEndpoint = httptransport.NewClient(
			"post",
			copyURL(u, "/action/"),
			svc.EncodeHTTPCreateOccurrenceZeroRequest,
			svc.DecodeHTTPCreateOccurrenceResponse,
			//httptransport.ClientBefore(opentracing.FromHTTPRequest(tracer, "Sum", logger)),
		).Endpoint()
		/*
			sumEndpoint = opentracing.TraceClient(tracer, "Sum")(sumEndpoint)
			sumEndpoint = limiter(sumEndpoint)
			sumEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
				Name:    "Sum",
				Timeout: 30 * time.Second,
			}))(sumEndpoint)
		*/
	}

	return svc.Endpoints{

		CreateActionEndpoint:     CreateActionZeroEndpoint,
		CreateOccurrenceEndpoint: CreateOccurrenceZeroEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
