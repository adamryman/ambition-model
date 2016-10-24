// Package http provides an HTTP client for the AmbitionService service.

package http

import (
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	// This Service
	svc "github.com/adamryman/ambition-model/ambition-service/generated"
	handler "github.com/adamryman/ambition-model/ambition-service/handlers/server"
)

var (
	_ = endpoint.Chain
	_ = httptransport.NewClient
)

// New returns a service backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string) (handler.Service, error) {
	//options := []httptransport.ServerOption{
	//httptransport.ServerBefore(),
	//}

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	_ = u

	return svc.Endpoints{}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
