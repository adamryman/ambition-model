package svc

// This file contains the Service definition, and a basic service
// implementation. It also includes service middlewares.

import (
//_ "errors"
//_ "time"

//_ "golang.org/x/net/context"

//_ "github.com/go-kit/kit/log"
//_ "github.com/go-kit/kit/metrics"

//pb "github.com/adamryman/ambition-truss/ambition/ambition-service"
)

// Middleware describes a service (as opposed to endpoint) middleware.
//type Middleware func(Service) Service

// ServiceLoggingMiddleware returns a service middleware that logs the
// parameters and result of each method invocation.
//func ServiceLoggingMiddleware(logger log.Logger) Middleware {
//return func(next Service) Service {
//return serviceLoggingMiddleware{
//logger: logger,
//next:   next,
//}
//}
//}

//type serviceLoggingMiddleware struct {
//logger log.Logger
//next   Service
//}

//func (mw serviceLoggingMiddleware) Sum(ctx context.Context, a, b int) (v int, err error) {
//defer func(begin time.Time) {
//mw.logger.Log(
//"method", "Sum",
//"a", a, "b", b, "result", v, "error", err,
//"took", time.Since(begin),
//)
//}(time.Now())
//return mw.next.Sum(ctx, a, b)
//}

//func (mw serviceLoggingMiddleware) Concat(ctx context.Context, a, b string) (v string, err error) {
//defer func(begin time.Time) {
//mw.logger.Log(
//"method", "Concat",
//"a", a, "b", b, "result", v, "error", err,
//"took", time.Since(begin),
//)
//}(time.Now())
//return mw.next.Concat(ctx, a, b)
//}

// ServiceInstrumentingMiddleware returns a service middleware that instruments
// the number of integers summed and characters concatenated over the lifetime of
// the service.
//func ServiceInstrumentingMiddleware(ints, chars metrics.Counter) Middleware {
//return func(next Service) Service {
//return serviceInstrumentingMiddleware{
//ints:  ints,
//chars: chars,
//next:  next,
//}
//}
//}

//type serviceInstrumentingMiddleware struct {
//ints  metrics.Counter
//chars metrics.Counter
//next  Service
//}

//func (mw serviceInstrumentingMiddleware) Sum(ctx context.Context, a, b int) (int, error) {
//v, err := mw.next.Sum(ctx, a, b)
//mw.ints.Add(uint64(v))
//return v, err
//}

//func (mw serviceInstrumentingMiddleware) Concat(ctx context.Context, a, b string) (string, error) {
//v, err := mw.next.Concat(ctx, a, b)
//mw.chars.Add(uint64(len(v)))
//return v, err
//}

// These annoying helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to represent errors.
// There is special casing to treat empty strings as nil errors.

//func str2err(s string) error {
//if s == "" {
//return nil
//}
//return errors.New(s)
//}

//func err2str(err error) string {
//if err == nil {
//return ""
//}
//return err.Error()
//}
