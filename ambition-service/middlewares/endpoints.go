package middlewares

import (
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	svc "github.com/adamryman/ambition-model/ambition-service/generated"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass in the middlewares you want applied to every endpoint.
	in.WrapAll(EndpointLoggingMiddleware(log.NewLogfmtLogger(os.Stdout)))

	// How to apply a middleware selectively.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	return in
}

// EndpointLoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func EndpointLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Log("error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}
